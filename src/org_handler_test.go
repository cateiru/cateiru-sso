package src_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestOrgGetHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgGetHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetHandler(c)
		require.NoError(t, err)

		response := []src.OrgResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 1)

		require.Equal(t, response[0].ID, orgId)
		require.Equal(t, response[0].Name, "test")

		require.Equal(t, response[0].Role, "owner")
	})

	t.Run("複数のorgに所属している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)
		RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetHandler(c)
		require.NoError(t, err)

		response := []src.OrgResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)
	})

	t.Run("成功: どこにも所属していない場合は空の配列が返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetHandler(c)
		require.NoError(t, err)

		response := []src.OrgResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 0)
	})
}

func TestOrgGetSimpleListHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgGetSimpleListHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: org一覧を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetSimpleListHandler(c)
		require.NoError(t, err)

		response := []src.OrgSimpleResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 1)
	})

	t.Run("成功: org_idを指定すると所属しているかを確認する", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetSimpleListHandler(c)
		require.NoError(t, err)

		response := []src.OrgSimpleResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 1)
	})

	t.Run("成功: ownerとmemberロールのみ取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		org1 := RegisterOrg(t, ctx)
		org2 := RegisterOrg(t, ctx)
		org3 := RegisterOrg(t, ctx)

		InviteUserInOrg(t, ctx, org1, &u, "owner")
		InviteUserInOrg(t, ctx, org2, &u, "member")
		InviteUserInOrg(t, ctx, org3, &u, "guest")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetSimpleListHandler(c)
		require.NoError(t, err)

		response := []src.OrgSimpleResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)
		findOwnerOrg := false
		findMemberOrg := false
		findGuestOrg := false

		for _, org := range response {
			if org.ID == org1 {
				findOwnerOrg = true
			}
			if org.ID == org2 {
				findMemberOrg = true
			}
			if org.ID == org3 {
				findGuestOrg = true
			}
		}

		require.True(t, findOwnerOrg)
		require.True(t, findMemberOrg)
		require.False(t, findGuestOrg)
	})

	t.Run("失敗: org_idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_id=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetSimpleListHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: orgにユーザーが所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetSimpleListHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどguest", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetSimpleListHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestOrgGetDetailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgGetDetailHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: Owner権限で詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "owner")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetDetailHandler(c)
		require.NoError(t, err)

		response := src.OrgDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ID, orgId)
		require.Equal(t, response.Name, "test")
		require.Equal(t, response.Image.String, "")
		require.Equal(t, response.Link.String, "")

		require.Equal(t, response.Role, "owner")
	})

	t.Run("成功: Member権限で詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetDetailHandler(c)
		require.NoError(t, err)

		response := src.OrgDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ID, orgId)
		require.Equal(t, response.Name, "test")
		require.Equal(t, response.Image.String, "")
		require.Equal(t, response.Link.String, "")

		require.Equal(t, response.Role, "member")
	})

	t.Run("成功: Guest権限で詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetDetailHandler(c)
		require.NoError(t, err)

		response := src.OrgDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ID, orgId)
		require.Equal(t, response.Name, "test")
		require.Equal(t, response.Image.String, "")
		require.Equal(t, response.Link.String, "")

		require.Equal(t, response.Role, "guest")
	})

	t.Run("失敗: org_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetDetailHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", "invalid"), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetDetailHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetDetailHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})
}

func TestOrgGetMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgGetMemberHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: orgのメンバー一覧が返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		// u2も参加させる
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetMemberHandler(c)
		require.NoError(t, err)

		response := []src.OrgUserResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)

		// order byしているため順序が保証されているはず
		require.Equal(t, response[0].User.ID, u.ID)
		require.Equal(t, response[0].User.UserName, u.UserName)
		require.Equal(t, response[0].Role, "owner")

		require.Equal(t, response[1].User.ID, u2.ID)
		require.Equal(t, response[1].User.UserName, u2.UserName)
		require.Equal(t, response[1].Role, "member")
	})

	t.Run("失敗: org_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetMemberHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		// u2も参加させる
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_id=aaaaaa", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetMemberHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		// u2も参加させる
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgGetMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})
}

func TestOrgPostMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgPostMemberHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: orgのメンバーを招待できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.NoError(t, err)

		response := &src.OrgUserResponse{}
		require.NoError(t, m.Json(response))

		require.Equal(t, response.User.ID, u2.ID)
		require.Equal(t, response.User.UserName, u2.UserName)
		require.Equal(t, response.Role, "member")

		dbOrgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, dbOrgUser.ID, response.ID)
		require.Equal(t, dbOrgUser.Role, "member")
	})

	t.Run("成功: emailでも招待できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.Email)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.NoError(t, err)

		response := &src.OrgUserResponse{}
		require.NoError(t, m.Json(response))

		require.Equal(t, response.User.ID, u2.ID)
		require.Equal(t, response.User.UserName, u2.UserName)
		require.Equal(t, response.Role, "member")

		dbOrgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, dbOrgUser.ID, response.ID)
		require.Equal(t, dbOrgUser.Role, "member")
	})

	t.Run("成功: roleが空の場合はguestになる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.NoError(t, err)

		response := &src.OrgUserResponse{}
		require.NoError(t, m.Json(response))

		require.Equal(t, response.User.ID, u2.ID)
		require.Equal(t, response.User.UserName, u2.UserName)
		require.Equal(t, response.Role, "guest")

		dbOrgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, dbOrgUser.ID, response.ID)
		require.Equal(t, dbOrgUser.Role, "guest")
	})

	t.Run("失敗: org_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name_or_email", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", "invalid")
		form.Insert("user_name_or_email", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})

	t.Run("失敗: user_name_or_emailが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=400, message=user_name_or_email is required")
	})

	t.Run("失敗: user_name_or_emailが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", "invalid")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=404, message=user not found, unique=10")
	})

	t.Run("失敗: roleが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "invalid")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=400, message=invalid role")
	})

	t.Run("失敗: ユーザーはすでにメンバーになっている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgPostMemberHandler(c)
		require.EqualError(t, err, "code=409, message=user already exists")
	})
}

func TestOrgUpdateMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgUpdateMemberHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("org_user_id", fmt.Sprint(orgUser.ID))
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: roleを変更できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_user_id", fmt.Sprint(orgUser.ID))
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.NoError(t, err)

		response := &src.OrgUserResponse{}
		require.NoError(t, m.Json(response))

		dbOrgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, dbOrgUser.Role, "member")
	})

	t.Run("失敗: org_user_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.EqualError(t, err, "code=400, message=org_user_id is required")
	})

	t.Run("失敗: org_user_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_user_id", "invalid")
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.EqualError(t, err, "code=400, message=invalid org_user_id")

	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_user_id", fmt.Sprint(orgUser.ID))
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_user_id", fmt.Sprint(orgUser.ID))
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})

	t.Run("失敗: roleが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_user_id", fmt.Sprint(orgUser.ID))
		form.Insert("role", "aaaa")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.EqualError(t, err, "code=400, message=invalid role")
	})

	t.Run("失敗: 自分自身のroleは変更できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)

		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_user_id", fmt.Sprint(orgUser.ID))
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgUpdateMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you can't change your role")
	})
}

func TestOrgDeleteMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgDeleteMemberHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUser.ID), http.MethodPost, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: orgからユーザーを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUser.ID), http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgDeleteMemberHandler(c)
		require.NoError(t, err)

		existOrgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.ID.EQ(orgUser.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existOrgUser)
	})

	t.Run("失敗: org_user_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgDeleteMemberHandler(c)
		require.EqualError(t, err, "code=400, message=org_user_id is required")
	})

	t.Run("失敗: org_user_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_user_id=invalid", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgDeleteMemberHandler(c)
		require.EqualError(t, err, "code=400, message=invalid org_user_id")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUser.ID), http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgDeleteMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		InviteUserInOrg(t, ctx, orgId, &u2, "guest")
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUser.ID), http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgDeleteMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})

	t.Run("失敗: 自分自身は削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)

		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUser.ID), http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgDeleteMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you can't delete yourself")
	})
}

func TestOrgInvitedMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// OrgSessionを追加する
	registerInviteOrgSession := func(email string, orgId string) {
		token, err := lib.RandomStr(31)
		require.NoError(t, err)

		inviteOrgSession := models.InviteOrgSession{
			Token:  token,
			Email:  email,
			Period: time.Now().Add(h.C.InviteOrgSessionPeriod),

			OrgID: orgId,
		}
		err = inviteOrgSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)
	}

	SessionTest(t, h.OrgInvitedMemberHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 招待中の一覧を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		sendEmail := RandomEmail(t)
		registerInviteOrgSession(sendEmail, orgId)

		cookie := RegisterSession(t, ctx, &u)
		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInvitedMemberHandler(c)
		require.NoError(t, err)

		response := []src.OrgInviteMemberResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 1)
		require.Equal(t, sendEmail, response[0].Email)
	})

	t.Run("成功: 招待中のユーザーがいない場合はからの配列が返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)
		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInvitedMemberHandler(c)
		require.NoError(t, err)

		response := []src.OrgInviteMemberResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 0)
	})

	t.Run("失敗: org_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInvitedMemberHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)
		m, err := easy.NewMock("/?org_id=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInvitedMemberHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		sendEmail := RandomEmail(t)
		registerInviteOrgSession(sendEmail, orgId)

		cookie := RegisterSession(t, ctx, &u)
		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInvitedMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		sendEmail := RandomEmail(t)
		registerInviteOrgSession(sendEmail, orgId)

		cookie := RegisterSession(t, ctx, &u)
		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInvitedMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})
}

func TestOrgInviteNewMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OrgInviteNewMemberHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		sendEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 対象のEmailに対して招待できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.NoError(t, err)

		inviteSessionExist, err := models.InviteOrgSessions(
			models.InviteOrgSessionWhere.Email.EQ(sendEmail),
			models.InviteOrgSessionWhere.OrgID.EQ(orgId),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, inviteSessionExist)
	})

	t.Run("失敗: org_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("org_id", "invalid")
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: emailが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=400, message=email is required")
	})

	t.Run("失敗: emailの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("email", "invalid")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=400, message=invalid email")
	})

	t.Run("失敗: そのemailのユーザーはすでに存在している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)
		RegisterUser(t, ctx, sendEmail)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=400, message=user already exists")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("email", sendEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteNewMemberHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})
}

func TestOrgInviteMemberDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerInviteOrgSession := func(email string, orgId string) *models.InviteOrgSession {
		token, err := lib.RandomStr(31)
		require.NoError(t, err)

		inviteOrgSession := models.InviteOrgSession{
			Token:  token,
			Email:  email,
			Period: time.Now().Add(h.C.InviteOrgSessionPeriod),

			OrgID: orgId,
		}
		err = inviteOrgSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		// emailは複数存在できるが衝突しないという浅い考え
		dbInviteOrgSession, err := models.InviteOrgSessions(
			models.InviteOrgSessionWhere.Email.EQ(email),
		).One(ctx, h.DB)
		require.NoError(t, err)

		return dbInviteOrgSession
	}

	SessionTest(t, h.OrgInviteMemberDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		sendEmail := RandomEmail(t)
		orgSession := registerInviteOrgSession(sendEmail, orgId)

		m, err := easy.NewMock(fmt.Sprintf("/?invite_id=%d", orgSession.ID), http.MethodDelete, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 招待をキャンセルできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)
		orgSession := registerInviteOrgSession(sendEmail, orgId)

		m, err := easy.NewMock(fmt.Sprintf("/?invite_id=%d", orgSession.ID), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteMemberDeleteHandler(c)
		require.NoError(t, err)

		existOrgSession, err := orgSession.Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, existOrgSession)
	})

	t.Run("失敗:invite_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteMemberDeleteHandler(c)
		require.EqualError(t, err, "code=400, message=invite_id is required")
	})

	t.Run("失敗:invite_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?invite_id=invalid", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteMemberDeleteHandler(c)
		require.EqualError(t, err, "code=400, message=invite_id is invalid")
	})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)
		orgSession := registerInviteOrgSession(sendEmail, orgId)

		m, err := easy.NewMock(fmt.Sprintf("/?invite_id=%d", orgSession.ID), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteMemberDeleteHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this organization, unique=16")
	})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		cookie := RegisterSession(t, ctx, &u)

		sendEmail := RandomEmail(t)
		orgSession := registerInviteOrgSession(sendEmail, orgId)

		m, err := easy.NewMock(fmt.Sprintf("/?invite_id=%d", orgSession.ID), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.OrgInviteMemberDeleteHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner")
	})
}
