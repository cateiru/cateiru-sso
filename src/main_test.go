package src_test

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

var DB *sql.DB
var C *src.Config

// これをしないとテストが失敗するため追加している
// ref. https://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests
var _ = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")
var _ = flag.String("test.config", "", "Overrides the default config")

func TestMain(m *testing.M) {
	src.InitLogging("test")

	C = src.TestConfig

	ctx := context.Background()
	db, err := sql.Open("mysql", C.DatabaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	err = resetDBTable(ctx, db)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	code := m.Run()
	os.Exit(code)
}

// テスト用にテーブルをクリアする
func resetDBTable(ctx context.Context, db *sql.DB) error {
	rows, err := queries.Raw("show tables").QueryContext(ctx, db)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		table := ""
		if err := rows.Scan(&table); err != nil {
			return err
		}

		// SQLインジェクションの影響は無いためSprintfを使用している
		if _, err := queries.Raw(fmt.Sprintf("TRUNCATE TABLE %s", table)).Exec(db); err != nil {
			return err
		}
	}

	return nil
}

// ランダムなEmailを作成する
func RandomEmail(t *testing.T) string {
	r, err := lib.RandomStr(10)
	require.NoError(t, err)
	return fmt.Sprintf("%s@exmaple.com", r)
}

// ユーザを新規作成する
func RegisterUser(t *testing.T, ctx context.Context, email string) models.User {
	id := ulid.Make()

	u := models.User{
		ID:    id.String(),
		Email: email,
	}

	err := u.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)

	dbU, err := models.Users(
		models.UserWhere.ID.EQ(id.String()),
	).One(ctx, DB)
	require.NoError(t, err)

	return *dbU
}

// テスト用
// ユーザのセッションを作成する
// セッションは1つ目のユーザのみで、2つ目以降のユーザはリフレッシュトークンのみ
func RegisterSession(t *testing.T, ctx context.Context, users ...*models.User) []*http.Cookie {
	if len(users) < 1 {
		t.Fatal("At least one user must be specified")
	}

	ua := &src.UserData{
		OS:       "Windows",
		Browser:  "Google Chrome",
		Device:   "",
		IsMobile: false,
	}
	ip := "203.0.113.2"

	session := src.NewSession(C, DB)
	createSession, err := session.NewRegisterSession(ctx, users[0], ua, ip)
	require.NoError(t, err)

	cookies := createSession.InsertCookie(C)

	// 他のユーザのリフレッシュトークンを設定する
	for _, u := range users[1:] {
		refreshToken, err := lib.RandomStr(63)
		require.NoError(t, err)
		id := ulid.Make()
		idBin, err := id.MarshalBinary()
		require.NoError(t, err)

		r := models.Refresh{
			ID:        refreshToken,
			UserID:    u.ID,
			HistoryID: idBin,

			Period: time.Now().Add(C.RefreshDBPeriod),
		}
		err = r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		history := models.LoginHistory{
			UserID: u.ID,

			RefreshID: idBin,

			Device:   null.NewString(ua.Device, true),
			Os:       null.NewString(ua.OS, true),
			Browser:  null.NewString(ua.Browser, true),
			IsMobile: null.NewBool(ua.IsMobile, true),

			IP: net.ParseIP(ip),
		}
		err = history.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID)
		cookies = append(cookies, &http.Cookie{
			Name:     refreshCookieName,
			Secure:   C.RefreshCookie.Secure,
			HttpOnly: C.RefreshCookie.HttpOnly,
			Path:     C.RefreshCookie.Path,
			MaxAge:   C.RefreshCookie.MaxAge,
			Expires:  time.Now().Add(time.Duration(C.RefreshCookie.MaxAge) * time.Second),
			SameSite: C.RefreshCookie.SameSite,

			Value: refreshToken,
		})
	}

	return cookies
}

// テスト用のダーミハンドラーを作成する
//
// モックしているやつ
// - ReCaptcha
// - Sender
func NewTestHandler(t *testing.T) *src.Handler {
	webauthn, err := lib.NewWebAuthn(C.WebAuthnConfig)
	require.NoError(t, err)

	s := src.NewSession(C, DB)

	return &src.Handler{
		DB:        DB,
		C:         C,
		ReCaptcha: &ReCaptchaMock{},
		Sender:    &SenderMock{},
		WebAuthn: &WebAuthnMock{
			M: webauthn,
		},
		Session:  s,
		Password: C.Password,
	}
}

// ユーザにパスワードを追加する
func RegisterPassword(t *testing.T, ctx context.Context, u *models.User, password ...string) {
	p := "password"
	if len(password) >= 1 {
		p = password[0]
	}
	hashed, salt, err := C.Password.HashPassword(p)
	require.NoError(t, err)

	passwordModel := models.Password{
		UserID: u.ID,
		Hash:   hashed,
		Salt:   salt,
	}
	err = passwordModel.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)
}
