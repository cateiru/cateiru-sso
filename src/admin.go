package src

import (
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type BrandUserBrandResponse struct {
	ID        uint   `json:"id"`
	BrandID   string `json:"brand_id"`
	BrandName string `json:"brand_name"`

	CreatedAt time.Time `json:"created_at"`
}

type StaffClientResponse struct {
	ClientID string      `json:"client_id"`
	Name     string      `json:"name"`
	Image    null.String `json:"image"`
}

type UserDetailResponse struct {
	User *models.User `json:"user"`

	Staff      *models.Staff            `json:"staff,omitempty"`
	UserBrands []BrandUserBrandResponse `json:"user_brands"`

	Clients []StaffClientResponse `json:"clients"`
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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
	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "user_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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
	userBrands, err := models.UserBrands(
		models.UserBrandWhere.UserID.EQ(user.ID),
		qm.Load(models.UserBrandRels.Brand),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	brandUserBrands := make([]BrandUserBrandResponse, len(userBrands))
	for i, userBrand := range userBrands {
		brandUserBrands[i] = BrandUserBrandResponse{
			ID:        userBrand.ID,
			BrandID:   userBrand.R.Brand.ID,
			BrandName: userBrand.R.Brand.Name,

			CreatedAt: userBrand.CreatedAt,
		}
	}

	// userが作成したclient
	clients, err := models.Clients(
		models.ClientWhere.OwnerUserID.EQ(user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	staffClients := make([]StaffClientResponse, len(clients))
	for i, client := range clients {
		staffClients[i] = StaffClientResponse{
			ClientID: client.ClientID,
			Name:     client.Name,
			Image:    client.Image,
		}
	}

	return c.JSON(http.StatusOK, UserDetailResponse{
		User: user,

		Staff:      staff,
		UserBrands: brandUserBrands,

		Clients: staffClients,
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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
func (h *Handler) AdminStaffHandler(c echo.Context) error {
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	// userが存在するかチェックする
	existUser, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existUser {
		return NewHTTPError(http.StatusNotFound, "user not found")
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
		return nil
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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

	return c.JSON(http.StatusOK, []*models.Brand{brand})
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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

// ブランドを削除する
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
	if err := h.Session.RequireStaff(ctx, u); err != nil {
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

// org取得
func (h *Handler) AdminOrgHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	orgs, err := models.Organizations(
		qm.OrderBy(models.OrganizationColumns.UpdatedAt),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, orgs)
}

// org作成（管理者用）
// args:
// - name: 組織名
// - link?: 組織の外部リンク（HPとか）
// - image?: 組織の画像
func (h *Handler) AdminOrgCreateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	name := c.FormValue("name")
	link := c.FormValue("link")

	imageHeader, err := c.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	parsedLink := null.NewString("", false)
	if link != "" {
		u, ok := lib.ValidateURL(link)
		if !ok {
			return NewHTTPError(http.StatusBadRequest, "invalid link")
		}
		parsedLink = null.NewString(u.String(), true)
	}

	orgId := ulid.Make()

	org := models.Organization{
		ID:   orgId.String(),
		Name: name,
		Link: parsedLink,
	}

	// 画像をアップロードする（ある場合）
	if imageHeader != nil {
		file, err := imageHeader.Open()
		if err != nil {
			return err
		}
		contentType := imageHeader.Header.Get("Content-Type")
		if !lib.ValidateContentType(contentType) {
			return NewHTTPError(http.StatusBadRequest, "invalid Content-Type")
		}
		path := filepath.Join("org", orgId.String())
		if err := h.Storage.Write(ctx, path, file, contentType); err != nil {
			return err
		}

		// ローカル環境では /[bucket-name]/org/[image] となるので
		p, err := url.JoinPath(h.C.CDNHost.Path, path)
		if err != nil {
			return err
		}

		url := &url.URL{
			Scheme: h.C.CDNHost.Scheme,
			Host:   h.C.CDNHost.Host,
			Path:   p,
		}
		if err := h.CDN.Purge(url.String()); err != nil {
			return err
		}

		org.Image = null.NewString(url.String(), true)
	}

	if err := org.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	dbOrg, err := models.Organizations(
		models.OrganizationWhere.ID.EQ(orgId.String()),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dbOrg)
}

// org更新（管理者用）
// args:
// - name: 組織名
// - link?: 組織の外部リンク（HPとか）
// - image?: 組織の画像。指定しないと変更しない
func (h *Handler) AdminOrgUpdateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	orgId := c.FormValue("org_id")
	name := c.FormValue("name")
	link := c.FormValue("link")

	imageHeader, err := c.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}
	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	parsedLink := null.NewString("", false)
	if link != "" {
		u, ok := lib.ValidateURL(link)
		if !ok {
			return NewHTTPError(http.StatusBadRequest, "invalid link")
		}
		parsedLink = null.NewString(u.String(), true)
	}

	org, err := models.Organizations(
		models.OrganizationWhere.ID.EQ(orgId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}
	if err != nil {
		return err
	}

	org.Name = name
	org.Link = parsedLink

	// 画像をアップロードする（ある場合）
	if imageHeader != nil {
		file, err := imageHeader.Open()
		if err != nil {
			return err
		}
		contentType := imageHeader.Header.Get("Content-Type")
		if !lib.ValidateContentType(contentType) {
			return NewHTTPError(http.StatusBadRequest, "invalid Content-Type")
		}
		path := filepath.Join("org", orgId)
		if err := h.Storage.Write(ctx, path, file, contentType); err != nil {
			return err
		}

		// ローカル環境では /[bucket-name]/org/[image] となるので
		p, err := url.JoinPath(h.C.CDNHost.Path, path)
		if err != nil {
			return err
		}

		url := &url.URL{
			Scheme: h.C.CDNHost.Scheme,
			Host:   h.C.CDNHost.Host,
			Path:   p,
		}
		if err := h.CDN.Purge(url.String()); err != nil {
			return err
		}

		org.Image = null.NewString(url.String(), true)
	}

	if _, err := org.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	dbOrg, err := models.Organizations(
		models.OrganizationWhere.ID.EQ(orgId),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dbOrg)
}

// org削除（管理者用）
func (h *Handler) AdminOrgDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	orgId := c.QueryParam("org_id")
	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).One(ctx, tx)
		if errors.Is(err, sql.ErrNoRows) {
			return NewHTTPError(http.StatusNotFound, "organization not found")
		}
		if err != nil {
			return err
		}

		if _, err := org.Delete(ctx, tx); err != nil {
			return err
		}

		_, err = models.OrganizationUsers(
			models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		).DeleteAll(ctx, tx)
		if err != nil {
			return err
		}

		_, err = models.InviteOrgSessions(
			models.InviteOrgSessionWhere.OrgID.EQ(orgId),
		).DeleteAll(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// org画像の削除
func (h *Handler) AdminOrgDeleteImageHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	orgId := c.QueryParam("org_id")
	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}

	org, err := models.Organizations(
		models.OrganizationWhere.ID.EQ(orgId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}
	if err != nil {
		return err
	}

	if !org.Image.Valid {
		return NewHTTPError(http.StatusBadRequest, "image is not set")
	}

	path := filepath.Join("org", orgId)
	if err := h.Storage.Delete(ctx, path); err != nil {
		return err
	}

	// ローカル環境では /[bucket-name]/org/[image] となるので
	p, err := url.JoinPath(h.C.CDNHost.Path, path)
	if err != nil {
		return err
	}

	url := &url.URL{
		Scheme: h.C.CDNHost.Scheme,
		Host:   h.C.CDNHost.Host,
		Path:   p,
	}
	if err := h.CDN.Purge(url.String()); err != nil {
		return err
	}

	org.Image = null.NewString("", false)

	if _, err := org.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}
