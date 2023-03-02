package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type LoginUser struct {
	Avatar            null.String `json:"avatar"`
	UserName          string      `json:"user_name"`
	AvailablePasskey  bool        `json:"available_passkey"`
	AvailablePassword bool        `json:"available_password"`
	AutoUsePasskey    bool        `json:"auto_use_passkey"`
}

// ユーザの情報を返す
// BOT使われると困るのでreCAPTCHA使いながら
func (h *Handler) LoginUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userNameOrEmail := c.FormValue("username_or_email")
	recaptcha := c.FormValue("recaptcha")
	ip := c.RealIP()

	if userNameOrEmail == "" {
		return NewHTTPError(http.StatusBadRequest, "username_or_email is empty")
	}
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}

	// reCAPTCHA
	if h.C.UseReCaptcha {
		order, err := h.ReCaptcha.ValidateOrder(recaptcha, ip)
		if err != nil {
			return err
		}
		// 検証に失敗した or スコアが閾値以下の場合はエラーにする
		if !order.Success || order.Score < h.C.ReCaptchaAllowScore {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrReCaptcha, "reCAPTCHA validation failed")
		}
	}

	user, err := findUserByUserNameOrEmail(ctx, h.DB, userNameOrEmail)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNotFoundUser, "user not found")
	}
	if err != nil {
		return err
	}

	availablePasskey := false
	autoUsePasskey := false
	availablePassword := false

	// Passkeyの判定
	ua, err := ParseUA(c.Request())
	if err != nil {
		return err
	}
	passkeyLoginDevices, err := models.PasskeyLoginDevices(
		models.PasskeyLoginDeviceWhere.UserID.EQ(user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	if len(passkeyLoginDevices) >= 1 {
		availablePasskey = true
	}

	if ua.Browser != "" && ua.OS != "" {
		for _, devices := range passkeyLoginDevices {
			// Passkeyを登録したOSと同じOSであれば自動ログイン
			// iCloudなどOSで共有可能な場合があるので
			if devices.IsRegisterDevice {
				if devices.Os.String == ua.OS {
					autoUsePasskey = true
					break
				}
			}
			// 過去にログインしたことあるブラウザ
			if devices.Os.String == ua.OS &&
				devices.Browser.String == ua.Browser &&
				devices.Device.String == ua.Device &&
				(!devices.IsMobile.Valid || devices.IsMobile.Bool == ua.IsMobile) { // IsMobileがある場合はそれも使用して判定する
				autoUsePasskey = true
				break
			}
		}
	}

	// パスワードの設定
	passwordExists, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	availablePassword = passwordExists

	// PasskeyかPasswordかならずどちらかは存在するはず
	if !availablePasskey && !availablePassword {
		return NewHTTPError(http.StatusInternalServerError, "no certificate")
	}

	loginUser := &LoginUser{
		Avatar:            user.Avatar,
		UserName:          user.UserName,
		AvailablePasskey:  availablePasskey,
		AutoUsePasskey:    autoUsePasskey,
		AvailablePassword: availablePassword,
	}
	return c.JSON(http.StatusOK, loginUser)
}

func findUserByUserNameOrEmail(ctx context.Context, db *sql.DB, userNameOrEmail string) (*models.User, error) {
	if lib.ValidateEmail(userNameOrEmail) {
		return models.Users(
			models.UserWhere.Email.EQ(userNameOrEmail),
		).One(ctx, db)
	}
	return models.Users(
		models.UserWhere.UserName.EQ(userNameOrEmail),
	).One(ctx, db)
}
