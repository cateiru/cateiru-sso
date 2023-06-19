package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type UserMeResponse struct {
	UserInfo *models.User    `json:"user"`
	Setting  *models.Setting `json:"setting,omitempty"`
	IsStaff  bool            `json:"is_staff"`
}

type UserUserNameResponse struct {
	UserName string `json:"user_name"`
	Ok       bool   `json:"ok"`
	Message  string `json:"message"`
}

type UserBrandResponse struct {
	BrandNames []string `json:"brand,omitempty"`
}

type UserOtpResponse struct {
	Enable   bool      `json:"enable"`
	Modified null.Time `json:"modified,omitempty"`
}

type UpdateEmailTemplate struct {
	User     *models.User
	NewEmail string
	Code     string
	Period   time.Time
}

type UserUpdateEmailResponse struct {
	Session string `json:"session"`
}

type UserUpdateEmailRegisterResponse struct {
	Email string `json:"email"`
}

type UserAvatarResponse struct {
	Avatar string `json:"avatar"`
}

// 他の人に公開可能な情報
type PublicUserResponse struct {
	ID       string      `json:"id"`
	UserName string      `json:"user_name"`
	Avatar   null.String `json:"avatar,omitempty"`
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

	isStaff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UserMeResponse{
		UserInfo: user,
		Setting:  setting,
		IsStaff:  isStaff,
	})
}

// ユーザ情報更新
func (h *Handler) UserUpdateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userName := c.FormValue("user_name")      // 指定しないと更新しない
	familyName := c.FormValue("family_name")  // 指定しないと削除
	middleName := c.FormValue("middle_name")  // 指定しないと削除
	givenName := c.FormValue("given_name")    // 指定しないと削除
	gender := c.FormValue("gender")           // 指定しないと更新しない
	birthDateStr := c.FormValue("birth_date") // 指定しないと削除
	localeID := c.FormValue("locale_id")      // 指定しないと更新しない

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// ユーザ名は削除不可
	if userName != "" {
		if !lib.ValidateUsername(userName) {
			return NewHTTPError(http.StatusBadRequest, "invalid user_name")
		}
		existUser, err := models.Users(
			models.UserWhere.UserName.EQ(userName),
		).Exists(ctx, h.DB)
		if err != nil {
			return err
		}
		if existUser {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrAlreadyExistUser, "user already exists")
		}

		user.UserName = userName
	}

	user.FamilyName = null.NewString(familyName, len(familyName) != 0)
	user.MiddleName = null.NewString(middleName, len(middleName) != 0)
	user.GivenName = null.NewString(givenName, len(givenName) != 0)

	// genderはデフォルト値が9
	if gender != "" {
		if !lib.ValidateGender(gender) {
			return NewHTTPError(http.StatusBadRequest, "invalid gender")
		}
		user.Gender = gender
	}

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
	if localeID != "" {
		if !lib.ValidateLocale(localeID) {
			return NewHTTPError(http.StatusBadRequest, "invalid locale_id")
		}
		user.LocaleID = localeID
	}

	if _, err := user.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// リアルタイムでユーザーIDが使用されているかチェックするハンドラ
func (h *Handler) UserUserNameHandler(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	userName := c.FormValue("user_name")
	if userName == "" {
		return c.JSON(http.StatusOK, &UserUserNameResponse{
			UserName: "",
			Ok:       false,
			Message:  "ユーザー名が指定されていません",
		})
	}

	if !lib.ValidateUsername(userName) {
		return c.JSON(http.StatusOK, &UserUserNameResponse{
			UserName: userName,
			Ok:       false,
			Message:  "ユーザー名は3文字以上15文字以下で半角英数字と'_'のみ使用できます",
		})
	}

	existUser, err := models.Users(
		models.UserWhere.UserName.EQ(userName),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	message := "ユーザー名は使用可能です"
	if existUser {
		message = "ユーザー名は既に使用されています"
	}

	return c.JSON(http.StatusOK, &UserUserNameResponse{
		UserName: userName,
		Ok:       !existUser,
		Message:  message,
	})
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

	// SELECT brand.* FROM brand
	// INNER JOIN user_brand ON
	//	 brand.id = user_brand.brand_id
	// WHERE user_brand.user_id = ?;
	brands, err := models.Brands(
		qm.Select("brand.*"),
		qm.InnerJoin("user_brand ON brand.id = user_brand.brand_id"),
		qm.Where("user_brand.user_id = ?", user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	brandNames := []string{}
	for _, b := range brands {
		brandNames = append(brandNames, b.Name)
	}

	return c.JSON(http.StatusOK, &UserBrandResponse{
		BrandNames: brandNames,
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

	existEmail, err := models.Users(
		models.UserWhere.Email.EQ(newEmail),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if existEmail {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrAlreadyExistUser, "email already used")
	}

	sessionId, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	code, err := lib.RandomNumber(6)
	if err != nil {
		return err
	}

	session := models.EmailVerifySession{
		ID:         sessionId,
		UserID:     user.ID,
		NewEmail:   newEmail,
		VerifyCode: code,
		Period:     time.Now().Add(h.C.UpdateEmailSessionPeriod),
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

	return c.JSON(http.StatusOK, &UserUpdateEmailResponse{
		Session: sessionId,
	})
}

// Email更新の確認して実際に更新するハンドラ
func (h *Handler) UserUpdateEmailRegisterHandler(c echo.Context) error {
	ctx := c.Request().Context()

	token := c.FormValue("update_token")
	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "update_token is empty")
	}
	code := c.FormValue("code")
	if code == "" {
		return NewHTTPError(http.StatusBadRequest, "code is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	session, err := models.EmailVerifySessions(
		models.EmailVerifySessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "invalid session")
	}
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
		return NewHTTPUniqueError(http.StatusForbidden, ErrAuthenticationFailed, "invalid code")
	}

	// Emailを更新する
	user.Email = session.NewEmail
	if _, err := user.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	// セッションは削除
	if _, err := session.Delete(ctx, h.DB); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UserUpdateEmailRegisterResponse{
		Email: user.Email,
	})
}

// アバターの更新
func (h *Handler) UserAvatarHandler(c echo.Context) error {
	ctx := c.Request().Context()

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, err)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	contentType := fileHeader.Header.Get("Content-Type")
	if !lib.ValidateContentType(contentType) {
		return NewHTTPError(http.StatusBadRequest, "invalid Content-Type")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	path := filepath.Join("avatar", user.ID)
	if err := h.Storage.Write(ctx, path, file, contentType); err != nil {
		return err
	}

	// ローカル環境では /[bucket-name]/avatar/[image] となるので
	p, err := url.JoinPath(h.C.CDNHost.Path, path)
	if err != nil {
		return err
	}

	// 画像URLはCDNをかますのでCDNのホストにする
	url := &url.URL{
		Scheme: h.C.CDNHost.Scheme,
		Host:   h.C.CDNHost.Host,
		Path:   p,
	}
	// user更新（設定していない場合）
	if !user.Avatar.Valid {
		user.Avatar = null.NewString(url.String(), true)
		if _, err := user.Update(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
	}

	if err := h.CDN.Purge(url.String()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UserAvatarResponse{
		Avatar: url.String(),
	})
}

// アバターの削除
func (h *Handler) UserDeleteAvatarHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	if !user.Avatar.Valid {
		return NewHTTPError(http.StatusBadRequest, "avatar is not set")
	}

	path := filepath.Join("avatar", user.ID)
	if err := h.Storage.Delete(ctx, path); err != nil {
		return err
	}

	// user更新
	user.Avatar = null.NewString("", false)
	if _, err := user.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	// ローカル環境では /[bucket-name]/avatar/[image] となるので
	p, err := url.JoinPath(h.C.CDNHost.Path, path)
	if err != nil {
		return err
	}

	// CDNをパージ
	url := &url.URL{
		Scheme: h.C.CDNHost.Scheme,
		Host:   h.C.CDNHost.Host,
		Path:   p,
	}
	if err := h.CDN.Purge(url.String()); err != nil {
		return err
	}

	return nil
}

// SSOクライアントからログアウトする
// TODO: SSOクライアントのログインを実装してからやる
func (h *Handler) UserLogoutClientHandler(c echo.Context) error {
	return nil
}

// ユーザを新規に作成する
// 最初は、ユーザ名などの情報はデフォルト値に設定する（ユーザ登録フローの簡略化のため）
func RegisterUser(ctx context.Context, db boil.ContextExecutor, email string, ids ...string) (*models.User, error) {
	// もう一度Emailが登録されていないか確認する
	exist, err := models.Users(models.UserWhere.Email.EQ(email)).Exists(ctx, db)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, NewHTTPUniqueError(http.StatusBadRequest, ErrImpossibleRegisterAccount, "impossible register account")
	}

	var id string
	if (len(ids)) <= 0 {
		id = ulid.Make().String()
	} else {
		id = ids[0]
	}

	u := models.User{
		ID:    id,
		Email: email,
	}
	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	L.Info("register user",
		zap.String("email", email),
	)

	return models.Users(
		models.UserWhere.ID.EQ(id),
	).One(ctx, db)
}

// ユーザ名かEmailを使用してユーザを引く
func FindUserByUserNameOrEmail(ctx context.Context, db *sql.DB, userNameOrEmail string) (*models.User, error) {
	if lib.ValidateEmail(userNameOrEmail) {
		return models.Users(
			models.UserWhere.Email.EQ(userNameOrEmail),
		).One(ctx, db)
	}
	u, err := models.Users(
		models.UserWhere.UserName.EQ(userNameOrEmail),
	).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewHTTPUniqueError(http.StatusNotFound, ErrNotFoundUser, "user not found")
	}
	return u, err
}
