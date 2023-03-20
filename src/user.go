package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type UserMeResponse struct {
	UserInfo *models.User    `json:"user"`
	Setting  *models.Setting `json:"setting,omitempty"`
}

type UserBrandResponse struct {
	Brand string `json:"brand,omitempty"`
}

type UpdateEmailTemplate struct {
	User     *models.User
	NewEmail string
	Code     string
	Period   time.Time
}

func (h *Handler) UserMeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	setting, err := models.Settings(
		models.SettingWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return c.JSON(http.StatusOK, &UserMeResponse{
		UserInfo: user,
		Setting:  setting,
	})
}

// ユーザ情報更新
func (h *Handler) UserUpdateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userName := c.FormValue("user_name")
	familyName := c.FormValue("family_name")
	middleName := c.FormValue("middle_name")
	givenName := c.FormValue("given_name")
	gender := c.FormValue("gender")
	birthDateStr := c.FormValue("birth_date")
	localeID := c.FormValue("locale_id")

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// ユーザ名は削除不可
	if !lib.ValidateUsername(userName) {
		return NewHTTPError(http.StatusBadRequest, "invalid user_name")
	}
	user.UserName = userName

	user.FamilyName = null.NewString(familyName, len(familyName) != 0)
	user.MiddleName = null.NewString(middleName, len(middleName) != 0)
	user.GivenName = null.NewString(givenName, len(givenName) != 0)

	// genderはデフォルト値が9
	if !lib.ValidateGender(gender) {
		return NewHTTPError(http.StatusBadRequest, "invalid gender")
	}
	user.Gender = gender

	if birthDateStr != "" {
		birthDate, ok := lib.ValidateBirthDate(birthDateStr)
		if !ok {
			return NewHTTPError(http.StatusBadRequest, "invalid birth_date")
		}

		user.Birthdate = null.NewTime(*birthDate, true)
	} else {
		user.Birthdate = null.NewTime(time.Now(), false)
	}

	// localeIDのデフォルト値はja-JP
	if !lib.ValidateLocale(localeID) {
		return NewHTTPError(http.StatusBadRequest, "invalid locale_id")
	}
	user.LocaleID = localeID

	if _, err := user.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// ユーザの設定を更新する
func (h *Handler) UserUpdateSettingHandler(c echo.Context) error {
	ctx := c.Request().Context()

	noticeEmail := c.FormValue("notice_email")
	noticeWebpush := c.FormValue("notice_webpush")

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	setting := models.Setting{
		UserID: user.ID,

		NoticeEmail:   noticeEmail == "true",
		NoticeWebpush: noticeWebpush == "true",
	}
	if err := setting.Upsert(ctx, h.DB, boil.Infer(), boil.Infer()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, setting)
}

func (h *Handler) UserBrandHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	brand, err := models.Brands(
		models.BrandWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return c.JSON(http.StatusOK, &UserBrandResponse{
			Brand: "",
		})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UserBrandResponse{
		Brand: brand.Brand,
	})
}

// Email更新リクエスト
func (h *Handler) UserUpdateEmailHandler(c echo.Context) error {
	ctx := c.Request().Context()

	newEmail := c.FormValue("new_email")
	if !lib.ValidateEmail(newEmail) {
		return NewHTTPError(http.StatusBadRequest, "empty new email")
	}
	recaptcha := c.FormValue("recaptcha")
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	ip := c.RealIP()

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

	id, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	code, err := lib.RandomNumber(6)
	if err != nil {
		return err
	}

	session := models.EmailVerifySession{
		ID:       id,
		UserID:   user.ID,
		NewEmail: newEmail,
		Period:   time.Now().Add(h.C.UpdateEmailSessionPeriod),
	}
	if err := session.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	m := &lib.MailBody{
		EmailAddress: newEmail,
		Subject:      "メールアドレスの確認して更新します",
		Data: UpdateEmailTemplate{
			User:     user,
			NewEmail: newEmail,
			Code:     code,
			Period:   session.Period,
		},
		PlainTextFileName: "update_email.gtpl",
		HTMLTextFileName:  "update_email.html",
	}

	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}

	msg, id, err := h.Sender.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("NewEmail", newEmail),
			zap.String("OldEmail", user.Email),
			zap.String("UserID", user.ID),
			zap.String("UserName", user.UserName),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", ip),
			zap.String("Device", ua.Device),
			zap.String("Browser", ua.Browser),
			zap.String("OS", ua.OS),
			zap.Bool("IsMobile", ua.IsMobile),
		)
		return err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("NewEmail", newEmail),
		zap.String("OldEmail", user.Email),
		zap.String("UserID", user.ID),
		zap.String("UserName", user.UserName),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", ip),
		zap.String("Device", ua.Device),
		zap.String("Browser", ua.Browser),
		zap.String("OS", ua.OS),
		zap.Bool("IsMobile", ua.IsMobile),
	)

	return nil
}

// Email更新の確認して実際に更新するハンドラ
func (h *Handler) UserUpdateEmailRegisterHandler(c echo.Context) error {
	ctx := c.Request().Context()

	token := c.FormValue("update_token")
	if token == "" {
		return NewHTTPError(http.StatusOK, "update_token is empty")
	}
	code := c.FormValue("code")
	if code == "" {
		return NewHTTPError(http.StatusOK, "code is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	session, err := models.EmailVerifySessions(
		models.EmailVerifySessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}
	if time.Now().After(session.Period) {
		// セッションは削除
		if _, err := session.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}
	if session.RetryCount >= h.C.UpdateEmailRetryCount {
		// セッションは削除
		if _, err := session.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExceededRetry, "exceeded retry")
	}
	if session.UserID != user.ID {
		return NewHTTPError(http.StatusBadRequest, "invalid user")
	}

	session.RetryCount++

	if code != session.VerifyCode {
		// リトライ回数更新する
		if _, err := session.Update(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		return NewHTTPError(http.StatusForbidden, "invalid code")
	}

	// Emailを更新する
	user.Email = session.NewEmail
	if _, err := user.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}
	return nil
}

// アバターの更新
func (h *Handler) UserAvatarHandler(c echo.Context) error {
	return nil
}

// アバターの削除
func (h *Handler) UserDeleteAvatarHandler(c echo.Context) error {
	return nil
}

// SSOクライアントからログアウトする
// TODO: SSOクライアントのログインを実装してからやる
func (h *Handler) UserLogoutClient(c echo.Context) error {
	return nil
}

// ユーザを新規に作成する
// 最初は、ユーザ名などの情報はデフォルト値に設定する（ユーザ登録フローの簡略化のため）
func RegisterUser(ctx context.Context, db *sql.DB, email string) (*models.User, error) {
	// もう一度Emailが登録されていないか確認する
	exist, err := models.Users(models.UserWhere.Email.EQ(email)).Exists(ctx, db)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, NewHTTPUniqueError(http.StatusBadRequest, ErrImpossibleRegisterAccount, "impossible register account")
	}

	id := ulid.Make()

	u := models.User{
		ID:    id.String(),
		Email: email,
	}
	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	L.Info("register user",
		zap.String("email", email),
	)

	return models.Users(
		models.UserWhere.ID.EQ(id.String()),
	).One(ctx, db)
}

// ユーザ名かEmailを使用してユーザを引く
func FindUserByUserNameOrEmail(ctx context.Context, db *sql.DB, userNameOrEmail string) (*models.User, error) {
	if lib.ValidateEmail(userNameOrEmail) {
		return models.Users(
			models.UserWhere.Email.EQ(userNameOrEmail),
		).One(ctx, db)
	}
	return models.Users(
		models.UserWhere.UserName.EQ(userNameOrEmail),
	).One(ctx, db)
}
