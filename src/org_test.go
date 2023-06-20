package src_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

// TODO: セッションのテスト
func TestOrgGetHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

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

func TestOrgGetMemberHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

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
		require.EqualError(t, err, "code=403, message=you are not member")
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
		require.EqualError(t, err, "code=403, message=you are not owner")
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
		require.EqualError(t, err, "code=404, message=organization not found")
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
	t.Run("成功: orgからユーザーを削除できる", func(t *testing.T) {})

	t.Run("失敗: org_user_idが空", func(t *testing.T) {})

	t.Run("失敗: org_user_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})

	t.Run("失敗: 自分自身は削除できない", func(t *testing.T) {})
}

func TestOrgInvitedMemberHandler(t *testing.T) {
	t.Run("成功: 招待中の一覧を取得できる", func(t *testing.T) {})

	t.Run("失敗: org_idが空", func(t *testing.T) {})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}

func TestOrgInviteNewMemberHandler(t *testing.T) {
	t.Run("成功: 対象のEmailに対して招待できる", func(t *testing.T) {})

	t.Run("失敗: org_idが空", func(t *testing.T) {})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: emailが空", func(t *testing.T) {})

	t.Run("失敗: emailの値が不正", func(t *testing.T) {})

	t.Run("失敗: そのemailのユーザーはすでに存在している", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}

func TestOrgInviteMemberDeleteHandler(t *testing.T) {
	t.Run("成功: 招待をキャンセルできる", func(t *testing.T) {})

	t.Run("失敗:invite_idが空", func(t *testing.T) {})

	t.Run("失敗:invite_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}
