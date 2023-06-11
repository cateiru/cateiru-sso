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
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/types"
)

var DB *sql.DB
var C *src.Config

// これをしないとテストが失敗するため追加している
// ref. https://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests
var _ = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")
var _ = flag.String("test.config", "", "Overrides the default config")

func TestMain(m *testing.M) {
	src.InitLogging("test")

	os.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")

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
		fmt.Print("-----------------------------\n")
		fmt.Print("データベースに接続できません。\n\n")
		fmt.Println("* データベースを立ち上げる場合は`./scripts/docker-compose-db.sh up -d`を実行してください。")
		fmt.Println("* すでにデータベースを立ち上げている場合は`src/config.go`の設定を見直してください。")
		fmt.Print("-----------------------------\n\n")
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

		if _, err := queries.Raw("SET FOREIGN_KEY_CHECKS = 0").ExecContext(ctx, db); err != nil {
			return err
		}

		// SQLインジェクションの影響は無いためSprintfを使用している
		if _, err := queries.Raw(fmt.Sprintf("TRUNCATE TABLE %s", table)).ExecContext(ctx, db); err != nil {
			return err
		}
	}

	if _, err := queries.Raw("SET FOREIGN_KEY_CHECKS = 1").ExecContext(ctx, db); err != nil {
		return err
	}

	return nil
}

// ----- 以下、テスト用便利メソッド -------

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

// ユーザーをスタッフにする
func ToStaff(t *testing.T, ctx context.Context, u *models.User) {
	staff := models.Staff{
		UserID: u.ID,
	}
	err := staff.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)
}

// UserAgentを設定する
func SetUserData(t *testing.T, m *easy.MockHandler, userData *src.UserData) {
	// iPhone safari
	if userData.Browser == "Safari" && userData.OS == "iOS" && userData.Device == "iPhone" && userData.IsMobile {
		m.R.Header.Set("User-Agent", `Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1`)
		return
	}
	// mac safari
	if userData.Browser == "Safari" && userData.OS == "macOS" && userData.Device == "" && !userData.IsMobile {
		m.R.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8`)
		return
	}

	// UA-CH
	m.R.Header.Set("Sec-Ch-Ua", fmt.Sprintf(`"%s";v="110"`, userData.Browser))
	m.R.Header.Set("Sec-Ch-Ua-Platform", fmt.Sprintf(`"%s"`, userData.OS))
	mobile := "?0"
	if userData.IsMobile {
		mobile = "?1"
	}
	m.R.Header.Set("Sec-Ch-Ua-Mobile", mobile)
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

	storage := lib.NewCloudStorage(C.StorageBucketName)

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
		Storage: &StorageMock{
			S: storage,
		},
		CDN: &CDNMock{},
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

// ユーザーにPasskeyを登録する
func RegisterPasskey(t *testing.T, ctx context.Context, u *models.User, userData ...src.UserData) {
	id, err := lib.RandomBytes(64)
	require.NoError(t, err)

	ua := &src.UserData{
		Device:   "",
		OS:       "Windows",
		Browser:  "Google Chrome",
		IsMobile: false,
	}
	if len(userData) > 0 {
		ua = &userData[0]
	}
	ip := "203.0.113.2"

	credential := webauthn.Credential{
		ID: id,
		Flags: webauthn.CredentialFlags{
			BackupState: true,
		},
	}

	// 認証を追加
	rowCredential := types.JSON{}
	err = rowCredential.Marshal(credential)
	require.NoError(t, err)

	passkey := models.Webauthn{
		UserID:     u.ID,
		Credential: rowCredential,

		Device:   null.NewString(ua.Device, true),
		Os:       null.NewString(ua.OS, true),
		Browser:  null.NewString(ua.Browser, true),
		IsMobile: null.NewBool(ua.IsMobile, true),

		IP: net.ParseIP(ip),
	}
	err = passkey.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)
}

// OTPを設定する
// 戻り値はOTPのシークレットとバックアップコード
func RegisterOTP(t *testing.T, ctx context.Context, u *models.User) (string, []string) {
	otp, err := lib.NewOTP(C.OTPIssuer, u.UserName)
	require.NoError(t, err)

	secret := otp.GetSecret()

	otpDB := models.Otp{
		UserID: u.ID,
		Secret: secret,
	}
	err = otpDB.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)

	backups := make([]string, 10)
	for i := 0; 10 > i; i++ {
		code, err := lib.RandomStr(15)
		require.NoError(t, err)
		backups[i] = code

		backupDB := models.OtpBackup{
			UserID: u.ID,
			Code:   code,
		}
		err = backupDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	return secret, backups
}

// SSOクライアントを作成する
// 戻り値は、(clientID, clientSecret)
func RegisterClient(t *testing.T, ctx context.Context, u *models.User, scopes ...string) (string, string) {
	clientID := ulid.Make()

	secret, err := lib.RandomStr(63)
	require.NoError(t, err)

	client := models.Client{
		ClientID: clientID.String(),

		Name: "test",

		OwnerUserID:  u.ID,
		ClientSecret: secret,
	}
	err = client.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)

	for _, scope := range scopes {
		clientScope := models.ClientScope{
			ClientID: clientID.String(),
			Scope:    scope,
		}
		err = clientScope.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	return clientID.String(), secret
}

// クライアントのAllow Ruleを作成する
func RegisterAllowRules(t *testing.T, ctx context.Context, clientId string, isUserId bool, value string) {
	if isUserId {
		rule := models.ClientAllowRule{
			ClientID: clientId,

			UserID: null.StringFrom(value),
		}
		err := rule.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	} else {
		rule := models.ClientAllowRule{
			ClientID: clientId,

			EmailDomain: null.StringFrom(value),
		}
		err := rule.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}
}

func RegisterBrand(t *testing.T, ctx context.Context, name string, description string, u ...*models.User) string {
	brandId, err := lib.RandomStr(31)
	require.NoError(t, err)

	brand := models.Brand{
		ID:          brandId,
		Name:        name,
		Description: null.NewString(description, description != ""),
	}
	err = brand.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)

	for _, user := range u {
		brandUser := models.UserBrand{
			BrandID: brand.ID,
			UserID:  user.ID,
		}
		err = brandUser.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	return brandId
}

// そのHandlerが認証が必要かどうかをテストする
func SessionTest(t *testing.T, h func(c echo.Context) error, newMock func(ctx context.Context, u *models.User) *easy.MockHandler) {
	ctx := context.Background()

	t.Run("正しく認証できている", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		m := newMock(ctx, &user)
		sessionCookies := RegisterSession(t, ctx, &user)
		m.Cookie(sessionCookies)
		c := m.Echo()

		err := h(c)
		require.NoError(t, err)
	})

	t.Run("Cookieが空の場合認証できない", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		m := newMock(ctx, &user)
		c := m.Echo()

		err := h(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
	})
}

// そのHandlerがスタッフ限定かどうかをテストする
func StaffAndSessionTest(t *testing.T, h func(c echo.Context) error, newMock func(ctx context.Context, u *models.User) *easy.MockHandler) {
	ctx := context.Background()

	defaultEmail := RandomEmail(t)
	defaultUser := RegisterUser(t, ctx, defaultEmail)

	secondEmail := RandomEmail(t)
	secondUser := RegisterUser(t, ctx, secondEmail)
	sessionCookies := RegisterSession(t, ctx, &secondUser)

	t.Run("正しく認証できていてスタッフである", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &user)

		staffSessionCookies := RegisterSession(t, ctx, &user)

		m := newMock(ctx, &user)
		m.Cookie(staffSessionCookies)
		c := m.Echo()

		err := h(c)
		require.NoError(t, err)
	})

	t.Run("Cookieが空の場合認証できない", func(t *testing.T) {
		m := newMock(ctx, &defaultUser)
		c := m.Echo()

		err := h(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
	})

	t.Run("スタッフではない", func(t *testing.T) {
		m := newMock(ctx, &secondUser)
		m.Cookie(sessionCookies)
		c := m.Echo()

		err := h(c)
		require.EqualError(t, err, "code=403, message=require staff")
	})
}

// Organizationを作成する
func RegisterOrg(t *testing.T, ctx context.Context, ownerUsers ...*models.User) string {
	id := ulid.Make()

	org := models.Organization{
		ID:   id.String(),
		Name: "test",
	}
	err := org.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)

	for _, u := range ownerUsers {
		orgUser := models.OrganizationUser{
			OrganizationID: org.ID,
			UserID:         u.ID,

			Role: "owner",
		}
		err = orgUser.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	return id.String()
}

// Organizationにユーザーを招待する
// roleは`owner`, `member`, `guest`
func InviteUserInOrg(t *testing.T, ctx context.Context, orgId string, u *models.User, role string) {
	orgUser := models.OrganizationUser{
		OrganizationID: orgId,
		UserID:         u.ID,

		Role: role,
	}
	err := orgUser.Insert(ctx, DB, boil.Infer())
	require.NoError(t, err)
}
