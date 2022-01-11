package tools

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/pquerna/otp/totp"
)

type DummyUser struct {
	UserID string
	Mail   string
	Otp    *secure.OnetimePassword
	Roles  []string
}

func NewDummyUser() *DummyUser {
	userID := utils.CreateID(30)
	mail := fmt.Sprintf("%s@example.com", utils.CreateID(5))

	return &DummyUser{
		UserID: userID,
		Mail:   mail,
		Otp:    nil,
		Roles:  []string{"user"},
	}
}

// Roleを追加する
func (c *DummyUser) AddRole(role string) *DummyUser {
	c.Roles = append(c.Roles, role)
	return c
}

// OTPをセットする
func (c *DummyUser) NewOTP() (*DummyUser, error) {

	otp, err := secure.NewOnetimePassword("test")
	if err != nil {
		return nil, err
	}
	c.Otp = otp
	return c, nil
}

// OTPのパスコードを生成する
func (c *DummyUser) GenOTPCode() (string, error) {
	return totp.GenerateCode(c.Otp.GetSecret(), time.Now().UTC())
}

// ユーザを追加する
// (テスト用)
func (c *DummyUser) AddUserInfo(ctx context.Context, db *database.Database) (*models.User, error) {
	userInfo := &models.User{
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
		UserName:  "TestUserName",
		Theme:     "Dark",
		AvatarUrl: "",

		Mail: c.Mail,

		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := userInfo.Add(ctx, db); err != nil {
		return nil, err
	}

	role := &models.Role{
		Role: c.Roles,

		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := role.Add(ctx, db); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// ユーザの認証情報を追加する
// (テスト用)
func (c *DummyUser) AddUserCert(ctx context.Context, db *database.Database) (*models.Certification, error) {
	otpSecret := ""
	otpBackups := []string{}

	// OTPが設定されている場合はセットする
	if c.Otp != nil {
		otpSecret = c.Otp.GetSecret()
		for i := 0; 10 > i; i++ {
			otpBackups = append(otpBackups, utils.CreateID(10))
		}
	}

	password := "password"

	hashedPassword, err := secure.PWHash(password)
	if err != nil {
		return nil, err
	}

	certification := &models.Certification{
		AccountCreateDate: time.Now(),

		OnetimePasswordSecret:  otpSecret,
		OnetimePasswordBackups: otpBackups,

		UserMailPW: models.UserMailPW{
			Mail:     c.Mail,
			Password: hashedPassword.Key,
			Salt:     hashedPassword.Salt,
		},
		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := certification.Add(ctx, db); err != nil {
		return nil, err
	}

	return certification, nil
}

// session-tokenとrefresh-tokenをセットする
// テスト用
func (c *DummyUser) AddLoginToken(ctx context.Context, db *database.Database, now time.Time) (string, string, error) {
	sessionToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

	session := &models.SessionInfo{
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: now,
			PeriodHour: 6,
		},
		UserId: models.UserId{
			UserId: c.UserID,
		},
	}
	refresh := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: now,
			PeriodDay:  7,
		},
		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := session.Add(ctx, db); err != nil {
		return "", "", err
	}
	if err := refresh.Add(ctx, db); err != nil {
		return "", "", err
	}

	return sessionToken, refreshToken, nil
}

// cookieをセットする
// テスト用
func SetCookie(jar *cookiejar.Jar, key string, value string, exp *net.CookieExp, url *url.URL) {
	cookie := &http.Cookie{
		Name:  key,
		Value: value,

		Secure:   false,
		Path:     "/",
		Domain:   "",
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
	}

	if !exp.IsSession {
		cookie.Expires = time.Now().Add(exp.GetTime())
		cookie.MaxAge = exp.GetNum()
	}

	jar.SetCookies(url, []*http.Cookie{cookie})
}

// responseをstringに変換する
func ConvertResp(resp *http.Response) string {
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return buf.String()
}

// responseをbytesに変換する
func ConvertByteResp(resp *http.Response) []byte {
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return buf.Bytes()
}
