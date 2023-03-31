package src

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserDetailResponse struct {
	User *models.User `json:"user"`

	Staff  *models.Staff   `json:"staff,omitempty"`
	Brands []*models.Brand `json:"brand,omitempty"`

	Clients []*models.Client `json:"client,omitempty"`
}

// すべてのユーザー一覧を取得する
// `?offset=0` 指定可能。
// 一度に返すユーザーの件数は50件
func (h *Handler) AdminUsersHandler(c echo.Context) error {
	ctx := c.Request().Context()

	offset := c.QueryParam("offset")
	offsetInt := 0
	if offset != "" {
		offsetIntA, err := strconv.Atoi(offset)
		if err != nil {
			return NewHTTPError(http.StatusBadRequest, "invalid offset")
		}
		offsetInt = offsetIntA
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	allUsers, err := models.Users(
		qm.OrderBy("id DESC"),
		qm.Limit(50),
		qm.Offset(offsetInt),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, allUsers)
}

// 指定したユーザーの詳細を取得する
func (h *Handler) AdminUserDetailHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.QueryParam("user_id")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err != nil {
		return err
	}

	// staff
	staff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// brand
	// SELECT brand.* FROM brands
	// INNER JOIN staffs ON staffs.brand_id = brands.id
	// WHERE staffs.user_id = ?;
	brands, err := models.Brands(
		qm.Select("brand.*"),
		qm.InnerJoin("user_brand ON brand.id = user_brand.brand_id"),
		qm.Where("user_brand.user_id = ?", user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	// userが作成したclient
	clients, err := models.Clients(
		models.ClientWhere.OwnerUserID.EQ(user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UserDetailResponse{
		User: user,

		Staff:  staff,
		Brands: brands,

		Clients: clients,
	})
}

// ユーザーにブランドを追加する
func (h *Handler) AdminUserBrandHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.FormValue("user_id")
	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "user_id is required")
	}
	brandId := c.FormValue("brand_id")
	if brandId == "" {
		return NewHTTPError(http.StatusBadRequest, "brand_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err != nil {
		return err
	}

	existBrand, err := models.Brands(
		models.BrandWhere.ID.EQ(brandId),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existBrand {
		return NewHTTPError(http.StatusNotFound, "brand not found")
	}

	userBrand := models.UserBrand{
		UserID:  user.ID,
		BrandID: brandId,
	}
	if err := userBrand.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// ブランドを削除する
func (h *Handler) AdminUserBrandDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.QueryParam("user_id")
	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "user_id is required")
	}
	brandId := c.QueryParam("brand_id")
	if brandId == "" {
		return NewHTTPError(http.StatusBadRequest, "brand_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err != nil {
		return err
	}

	_, err = models.UserBrands(
		models.UserBrandWhere.UserID.EQ(user.ID),
		models.UserBrandWhere.BrandID.EQ(brandId),
	).DeleteAll(ctx, h.DB)
	if err != nil {
		return err
	}

	return nil
}

// スタッフフラグの追加と削除を行う
func (h *Handler) AdminStuffHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.FormValue("user_id")
	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "user_id is required")
	}
	memo := c.FormValue("memo")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	staff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(userId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		// 新規追加
		newStaff := models.Staff{
			UserID: userId,
			Memo:   null.NewString(memo, memo != ""),
		}
		if err := newStaff.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	// 削除
	if _, err := staff.Delete(ctx, h.DB); err != nil {
		return err
	}

	return nil
}

// TODO: 通知は後々実装する
func (h *Handler) AdminBroadcastHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	return nil
}

// ブランドを返す
// `brand_id`を指定するとそのidのブランドを返す
func (h *Handler) AdminBrandHandler(c echo.Context) error {
	ctx := c.Request().Context()

	brandId := c.QueryParam("brand_id")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	if brandId == "" {
		// すべてのブランドを返す
		brands, err := models.Brands().All(ctx, h.DB)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, brands)
	}

	brand, err := models.Brands(
		models.BrandWhere.ID.EQ(brandId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "brand not found")
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, []models.Brand{*brand})
}

// ブランドを新規作成する
func (h *Handler) AdminBrandCreateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	name := c.FormValue("name")
	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}
	description := c.FormValue("description")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	brandId, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	brand := models.Brand{
		ID:   brandId,
		Name: name,

		Description: null.NewString(description, description != ""),
	}
	if err := brand.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, brand)
}

// ブランドを更新する
func (h *Handler) AdminBrandUpdateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	brandId := c.FormValue("brand_id")
	if brandId == "" {
		return NewHTTPError(http.StatusBadRequest, "brand_id is required")
	}
	name := c.FormValue("name")
	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}
	description := c.FormValue("description")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	brand, err := models.Brands(
		models.BrandWhere.ID.EQ(brandId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "brand not found")
	}
	if err != nil {
		return err
	}

	brand.Name = name
	brand.Description = null.NewString(description, description != "")

	if _, err := brand.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (h *Handler) AdminBrandDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()

	brandId := c.FormValue("brand_id")
	if brandId == "" {
		return NewHTTPError(http.StatusBadRequest, "brand_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStuff(ctx, u); err != nil {
		return err
	}

	_, err = models.Brands(
		models.BrandWhere.ID.EQ(brandId),
	).DeleteAll(ctx, h.DB)
	if err != nil {
		return err
	}

	return nil
}
