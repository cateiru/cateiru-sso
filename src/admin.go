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

type OrgAdminDetailResponse struct {
	Org     *models.Organization  `json:"org"`
	Users   []OrgUserResponse     `json:"users"`
	Clients []StaffClientResponse `json:"clients"`
}

type StaffClientDetailResponse struct {
	Client *models.Client `json:"client"`

	RedirectUrls []string `json:"redirect_urls"`
	ReferrerUrls []string `json:"referrer_urls"`

	Scopes []string `json:"scopes"`

	AllowRules []ClientAllowUserRuleResponse `json:"allow_rules"`
}

type RegisterSessionResponse struct {
	Email         string      `json:"email"`
	EmailVerified bool        `json:"email_verified"`
	SendCount     uint8       `json:"send_count"`
	RetryCount    uint8       `json:"retry_count"`
	OrgId         null.String `json:"org_id"`

	Period    time.Time `json:"period"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	usersJson := make([]*models.User, len(allUsers))
	copy(usersJson, allUsers)

	return c.JSON(http.StatusOK, usersJson)
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
		models.ClientWhere.OrgID.IsNull(),
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

	existUserBrand, err := models.UserBrands(
		models.UserBrandWhere.UserID.EQ(user.ID),
		models.UserBrandWhere.BrandID.EQ(brandId),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if existUserBrand {
		return NewHTTPError(http.StatusConflict, "user brand already exists")
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
	isStaff := c.FormValue("is_staff")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	// userが存在するかチェックする
	toUser, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err != nil {
		return err
	}

	staff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(userId),
	).One(ctx, h.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if isStaff == "true" {
		// 本当はUpsert使いたいけど、なんかINSERTしかされないので存在しているかを見る
		if staff != nil {
			// すでに存在している場合は更新
			staff.Memo = null.NewString(memo, memo != "")
			if _, err := staff.Update(ctx, h.DB, boil.Infer()); err != nil {
				return err
			}
		} else {
			// 存在していない場合は新規作成
			newStaff := models.Staff{
				UserID: userId,
				Memo:   null.NewString(memo, memo != ""),
			}
			if err := newStaff.Insert(ctx, h.DB, boil.Infer()); err != nil {
				return err
			}
		}
	} else {
		// 自分自身は削除できない
		if toUser.ID == u.ID {
			return NewHTTPError(http.StatusBadRequest, "can't change yourself")
		}

		// スタッフを外す
		if staff == nil {
			return NewHTTPError(http.StatusBadRequest, "user is not staff")
		}

		if _, err := staff.Delete(ctx, h.DB); err != nil {
			return err
		}
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
		brandJson := make([]*models.Brand, len(brands))
		copy(brandJson, brands)
		return c.JSON(http.StatusOK, brandJson)
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

	brandId := ulid.Make()

	brand := models.Brand{
		ID:   brandId.String(),
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

	orgsJson := make([]*models.Organization, len(orgs))
	copy(orgsJson, orgs)

	return c.JSON(http.StatusOK, orgsJson)
}

func (h *Handler) AdminOrgDetailHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orgId := c.QueryParam("org_id")
	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	org, err := models.Organizations(
		models.OrganizationWhere.ID.EQ(orgId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "org not found")
	}
	if err != nil {
		return err
	}

	orgUsers, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		qm.Load(models.OrganizationUserRels.User),
		qm.OrderBy("role ASC, created_at ASC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	users := make([]OrgUserResponse, len(orgUsers))
	for i, orgUser := range orgUsers {
		users[i] = OrgUserResponse{
			ID: orgUser.ID,

			User: PublicUserResponse{
				ID:       orgUser.R.User.ID,
				UserName: orgUser.R.User.UserName,
				Avatar:   orgUser.R.User.Avatar,
			},
			Role: orgUser.Role,

			CreatedAt: orgUser.CreatedAt,
			UpdatedAt: orgUser.UpdatedAt,
		}
	}

	clients, err := models.Clients(
		models.ClientWhere.OrgID.EQ(null.NewString(orgId, true)),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	orgClients := make([]StaffClientResponse, len(clients))
	for i, client := range clients {
		orgClients[i] = StaffClientResponse{
			ClientID: client.ClientID,
			Name:     client.Name,
			Image:    client.Image,
		}
	}

	response := &OrgAdminDetailResponse{
		Org:     org,
		Users:   users,
		Clients: orgClients,
	}

	return c.JSON(http.StatusOK, response)
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

func (h *Handler) AdminOrgMemberJoinHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orgId := c.FormValue("org_id")
	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}
	userNameOrEmail := c.FormValue("user_name_or_email")
	if userNameOrEmail == "" {
		return NewHTTPError(http.StatusBadRequest, "user_name_or_email is required")
	}
	role := c.FormValue("role")
	if role == "" {
		// 指定しないとguest（一番下の権限）にする
		role = "guest"
	}
	if !lib.ValidateRole(role) {
		return NewHTTPError(http.StatusBadRequest, "invalid role")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	user, err := FindUserByUserNameOrEmail(ctx, h.DB, userNameOrEmail)
	if err != nil {
		return err
	}

	organizationExist, err := models.Organizations(
		models.OrganizationWhere.ID.EQ(orgId),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !organizationExist {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}

	// すでにメンバーになっているかどうかを見る
	orgUserExist, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		models.OrganizationUserWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if orgUserExist {
		return NewHTTPError(http.StatusConflict, "user already exists")
	}

	// メンバーに追加する
	newOrgUser := &models.OrganizationUser{
		OrganizationID: orgId,
		UserID:         user.ID,
		Role:           role,
	}
	if err := newOrgUser.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (h *Handler) AdminOrgMemberRemoveHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orgUserId := c.QueryParam("org_user_id")
	if orgUserId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_user_id is required")
	}
	orgUserIdInt, err := strconv.Atoi(orgUserId)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "invalid org_user_id")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	orgUser, err := models.OrganizationUsers(
		models.OrganizationUserWhere.ID.EQ(uint(orgUserIdInt)),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization user not found")
	}
	if err != nil {
		return err
	}

	// メンバーを削除する
	if _, err := orgUser.Delete(ctx, h.DB); err != nil {
		return err
	}

	return nil
}

// クライアント一覧
func (h *Handler) AdminClientsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	clients, err := models.Clients(
		qm.OrderBy(models.ClientColumns.CreatedAt),
		qm.Limit(50), // 一旦offsetなしで50件表示させる
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	clientsResponse := make([]*StaffClientResponse, len(clients))
	for i, client := range clients {
		clientsResponse[i] = &StaffClientResponse{
			ClientID: client.ClientID,
			Name:     client.Name,
			Image:    client.Image,
		}
	}

	return c.JSON(http.StatusOK, clientsResponse)
}

// クライアントの詳細
func (h *Handler) AdminClientDetailHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.QueryParam("client_id")
	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "client not found")
	}
	if err != nil {
		return err
	}

	redirectUrlRecords, err := models.ClientRedirects(
		models.ClientRedirectWhere.ClientID.EQ(client.ClientID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	redirectUrls := make([]string, len(redirectUrlRecords))
	for i, redirect := range redirectUrlRecords {
		redirectUrls[i] = redirect.URL
	}

	referrerUrlRecords, err := models.ClientReferrers(
		models.ClientReferrerWhere.ClientID.EQ(client.ClientID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	referrerUrls := make([]string, len(referrerUrlRecords))
	for i, referrer := range referrerUrlRecords {
		// referrerはホストのみを見るので
		referrerUrls[i] = referrer.Host
	}

	scopesRecords, err := models.ClientScopes(
		models.ClientScopeWhere.ClientID.EQ(client.ClientID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	scopes := make([]string, len(scopesRecords))
	for i, scope := range scopesRecords {
		scopes[i] = scope.Scope
	}

	allowRules, err := getClientAllowRules(ctx, h.DB, client.ClientID)
	if err != nil {
		return err
	}

	response := &StaffClientDetailResponse{
		Client:       client,
		RedirectUrls: redirectUrls,
		ReferrerUrls: referrerUrls,
		Scopes:       scopes,
		AllowRules:   allowRules,
	}

	return c.JSON(http.StatusOK, response)
}

// メールのテンプレートをプレビューする
// HTMLのみ対応
func (h *Handler) AdminPreviewTemplateHTMLHandler(c echo.Context) error {
	ctx := c.Request().Context()

	name := c.Param("name")
	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	ip := c.RealIP()
	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}

	e := NewEmail(h.Sender, h.C, "example@example.test", ua, ip, u)
	e.HasPreviewMode = true

	template := ""

	switch name {
	case "register":
		template, err = e.RegisterEmailVerify("123456")
	case "register_resend":
		template, err = e.ResendRegisterEmailVerify("123456")
	case "update_email":
		template, err = e.UpdateEmail("old@example.test", "123456")
	case "update_password":
		template, err = e.UpdatePassword("token", u.UserName)
	case "invite_org":
		template, err = e.InviteOrg("token", "OrgName", u.UserName)
	case "test":
		template, err = e.Test()
	}

	if err != nil {
		return err
	}

	if template == "" {
		return NewHTTPError(http.StatusBadRequest, "invalid name")
	}

	return c.HTML(http.StatusOK, template)
}

// 登録セッションの一覧を返す
// PKはセッショントークンとして使用しているので削除しない
// メールアドレスをPKとしてあつかう（uniqueなので）
func (h *Handler) AdminRegisterSessionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	registerSessions, err := models.RegisterSessions(
		qm.Where("period > NOW()"),
		qm.OrderBy("created_at DESC"),
		qm.Limit(50), // 一旦決め打ち
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	registerSessionsResponse := make([]RegisterSessionResponse, len(registerSessions))

	for i, registerSession := range registerSessions {
		registerSessionsResponse[i] = RegisterSessionResponse{
			Email:         registerSession.Email,
			EmailVerified: registerSession.EmailVerified,
			SendCount:     registerSession.SendCount,
			RetryCount:    registerSession.RetryCount,
			Period:        registerSession.Period,
			OrgId:         registerSession.OrgID,

			CreatedAt: registerSession.CreatedAt,
			UpdatedAt: registerSession.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, registerSessionsResponse)
}

// 登録セッションを削除する
// 通常、登録セッションは数十分で削除されるが、管理者が直接削除できるようにする
func (h *Handler) AdminDeleteRegisterSessionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.QueryParam("email")
	if email == "" {
		return NewHTTPError(http.StatusBadRequest, "email is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.Email.EQ(email),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "register session not found")
	}
	if err != nil {
		return err
	}

	if _, err := registerSession.Delete(ctx, h.DB); err != nil {
		return err
	}

	return nil
}

func (h *Handler) AdminUserNameHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}
	if err := h.Session.RequireStaff(ctx, u); err != nil {
		return err
	}

	userNames, err := models.UserNames(
		models.UserNameWhere.Period.GT(time.Now()),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userNames)
}
