package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type OrgResponse struct {
	ID string `json:"id"`

	Name  string      `json:"name"`
	Image null.String `json:"image,omitempty"`
	Link  null.String `json:"link,omitempty"`
}

type OrgUserResponse struct {
	ID uint `json:"id"`

	User PublicUserResponse `json:"user"`
	Role string             `json:"role"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func getOrgUser(ctx context.Context, db *sql.DB, userId string, orgId string) (*OrgUserResponse, error) {
	orgUser, err := models.OrganizationUsers(
		models.OrganizationUserWhere.UserID.EQ(userId),
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		qm.Load(models.OrganizationUserRels.User),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	return &OrgUserResponse{
		ID: orgUser.ID,

		User: PublicUserResponse{
			ID:       orgUser.R.User.ID,
			UserName: orgUser.R.User.UserName,
			Avatar:   orgUser.R.User.Avatar,
		},
		Role: orgUser.Role,

		CreatedAt: orgUser.CreatedAt,
		UpdatedAt: orgUser.UpdatedAt,
	}, nil
}

// 所属しているorgを返す
// すべてのロールについて返す
func (h *Handler) OrgGetHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// 	SELECT organization.*
	// FROM organization
	// INNER JOIN organization_user
	//     ON organization.id = organization_user.organization_id
	// WHERE organization_user.user_id = ?
	// AND organization_user.role IN ('owner', 'member', 'guest')
	// ORDER BY organization.name ASC;
	orgs, err := models.Organizations(
		qm.InnerJoin("organization_user ON organization_user.organization_id = organization.id"),
		qm.Where("organization_user.user_id = ?", u.ID),
		qm.WhereIn("organization_user.role IN ?", []string{"owner", "member", "guest"}),
		qm.OrderBy("organization.name ASC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	response := make([]*OrgResponse, len(orgs))
	for i, org := range orgs {
		response[i] = &OrgResponse{
			ID:    org.ID,
			Name:  org.Name,
			Image: org.Image,
			Link:  org.Link,
		}
	}

	return c.JSON(http.StatusOK, response)
}

// orgに所属しているメンバーを返す
// ownerロールのみ
func (h *Handler) OrgGetMemberHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orgId := c.Param("org_id")
	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
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

	// ユーザーがownerかどうかを見る
	orgUser, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		models.OrganizationUserWhere.UserID.EQ(u.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}
	if err != nil {
		return err
	}
	if orgUser.Role != "owner" {
		return NewHTTPError(http.StatusForbidden, "you are not owner")
	}

	orgUsers, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		qm.Load(models.OrganizationUserRels.User),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	response := make([]*OrgUserResponse, len(orgUsers))
	for i, orgUser := range orgUsers {
		response[i] = &OrgUserResponse{
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

	return c.JSON(http.StatusOK, response)
}

// orgにメンバーを追加する
// ownerロールのみ
// すでにアカウントが存在している場合
func (h *Handler) OrgPostMemberHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orgId := c.FormValue("org_id")
	if orgId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_id is required")
	}
	userId := c.FormValue("user_id")
	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "user_id is required")
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

	inviteUserExist, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !inviteUserExist {
		return NewHTTPError(http.StatusNotFound, "user not found")
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

	// ユーザーがownerかどうかを見る
	orgUser, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		models.OrganizationUserWhere.UserID.EQ(u.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}
	if err != nil {
		return err
	}
	if orgUser.Role != "owner" {
		return NewHTTPError(http.StatusForbidden, "you are not owner")
	}

	// すでにメンバーになっているかどうかを見る
	orgUserExist, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgId),
		models.OrganizationUserWhere.UserID.EQ(userId),
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
		UserID:         userId,
		Role:           role,
	}
	if err := newOrgUser.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	dbOrgUser, err := getOrgUser(ctx, h.DB, userId, orgId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dbOrgUser)
}

// メンバーの権限を変える
// ownerロールのみ
func (h *Handler) OrgUpdateMemberHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orgUserId := c.FormValue("org_user_id")
	if orgUserId == "" {
		return NewHTTPError(http.StatusBadRequest, "org_user_id is required")
	}
	orgUserIdInt, err := strconv.Atoi(orgUserId)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "invalid org_user_id")
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

	orgUser, err := models.OrganizationUsers(
		models.OrganizationUserWhere.ID.EQ(uint(orgUserIdInt)),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization user not found")
	}
	if err != nil {
		return err
	}

	// ユーザーがownerかどうかを見る
	orgUserOwner, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgUser.OrganizationID),
		models.OrganizationUserWhere.UserID.EQ(u.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}
	if err != nil {
		return err
	}
	if orgUserOwner.Role != "owner" {
		return NewHTTPError(http.StatusForbidden, "you are not owner")
	}

	// 権限を変更する
	orgUser.Role = role

	if _, err := orgUser.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	dbOrgUser, err := getOrgUser(ctx, h.DB, orgUser.UserID, orgUser.OrganizationID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dbOrgUser)
}

// メンバーを削除する
// ownerロールのみ
func (h *Handler) OrgDeleteMemberHandler(c echo.Context) error {
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

	orgUser, err := models.OrganizationUsers(
		models.OrganizationUserWhere.ID.EQ(uint(orgUserIdInt)),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization user not found")
	}
	if err != nil {
		return err
	}

	// ユーザーがownerかどうかを見る
	orgUserOwner, err := models.OrganizationUsers(
		models.OrganizationUserWhere.OrganizationID.EQ(orgUser.OrganizationID),
		models.OrganizationUserWhere.UserID.EQ(u.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "organization not found")
	}
	if err != nil {
		return err
	}
	if orgUserOwner.Role != "owner" {
		return NewHTTPError(http.StatusForbidden, "you are not owner")
	}

	// メンバーを削除する
	if _, err := orgUser.Delete(ctx, h.DB); err != nil {
		return err
	}

	return nil
}

// アカウントを持っていない人に対してメールアドレスに招待メールを送信する
func (h *Handler) OrgInviteMemberHandler(c echo.Context) error {
	return nil
}
