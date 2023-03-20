package src

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func (h *Handler) UserMeHandler(c echo.Context) error {
	return nil
}

// ユーザ情報更新
func (h *Handler) UserUpdateHandler(c echo.Context) error {
	return nil
}

// ユーザの設定を更新する
func (h *Handler) UserUpdateSettingHandler(c echo.Context) error {
	return nil
}

func (h *Handler) UserBrandHandler(c echo.Context) error {
	return nil
}

// Email更新リクエスト
func (h *Handler) UserUpdateEmailHandler(c echo.Context) error {
	return nil
}

// Email更新の確認して実際に更新するハンドラ
func (h *Handler) UserUpdateEmailRegisterHandler(c echo.Context) error {
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
