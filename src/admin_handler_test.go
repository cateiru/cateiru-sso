package src_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestAdminUsersHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminUsersHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: 複数のユーザーを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUsersHandler(c)
		require.NoError(t, err)

		response := []models.User{}
		require.NoError(t, m.Json(&response))
		require.NotEqual(t, len(response), 0)
	})

	t.Run("成功: offsetを指定できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?offset=2", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUsersHandler(c)
		require.NoError(t, err)
	})
}

func TestAdminUserDetailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminUserDetailHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock(fmt.Sprintf("/?user_id=%s", u.ID), http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	email2 := RandomEmail(t)
	u2 := RegisterUser(t, ctx, email2)

	staff := models.Staff{
		UserID: u2.ID,
	}
	err := staff.Insert(ctx, h.DB, boil.Infer())
	require.NoError(t, err)

	RegisterBrand(t, ctx, "test", "", &u2)
	RegisterClient(t, ctx, &u2)

	t.Run("成功: ユーザーの詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?user_id=%s", u2.ID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserDetailHandler(c)
		require.NoError(t, err)

		response := src.UserDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.User.ID, u2.ID)
	})

	t.Run("失敗: user_idを指定しないとエラー", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserDetailHandler(c)
		require.EqualError(t, err, "code=400, message=user_id is required")
	})

	t.Run("失敗: user_idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?user_id=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserDetailHandler(c)
		require.EqualError(t, err, "code=404, message=user not found")
	})
}

func TestAdminUserBrandHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	brandId := RegisterBrand(t, ctx, "test", "")

	StaffAndSessionTest(t, h.AdminUserBrandHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		form.Insert("brand_id", brandId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: 追加できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		form.Insert("brand_id", brandId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserBrandHandler(c)
		require.NoError(t, err)

		ub, err := models.UserBrands(
			models.UserBrandWhere.UserID.EQ(u.ID),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.True(t, ub)
	})

	t.Run("失敗: ユーザーIDが不正な場合エラー", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", "invalid")
		form.Insert("brand_id", brandId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserBrandHandler(c)
		require.EqualError(t, err, "code=404, message=user not found")
	})

	t.Run("失敗: ブランドIDが不正な場合エラー", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		form.Insert("brand_id", "invalid")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserBrandHandler(c)
		require.EqualError(t, err, "code=404, message=brand not found")
	})

	t.Run("失敗: すでにユーザーはそのブランドに入っている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		brandId := RegisterBrand(t, ctx, "test", "", &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		form.Insert("brand_id", brandId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserBrandHandler(c)
		require.EqualError(t, err, "code=409, message=user brand already exists")
	})
}

func TestAdminUserBrandDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminUserBrandDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		brandId := RegisterBrand(t, ctx, "test", "", u)

		m, err := easy.NewMock(fmt.Sprintf("/?user_id=%s&brand_id=%s", u.ID, brandId), http.MethodDelete, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: 削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		brandId := RegisterBrand(t, ctx, "test", "", &u)

		m, err := easy.NewMock(fmt.Sprintf("/?user_id=%s&brand_id=%s", u.ID, brandId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserBrandDeleteHandler(c)
		require.NoError(t, err)

		ub, err := models.UserBrands(
			models.UserBrandWhere.UserID.EQ(u.ID),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, ub)
	})

	t.Run("失敗: ユーザーIDが不正な場合エラー", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		brandId := RegisterBrand(t, ctx, "test", "", &u)

		m, err := easy.NewMock(fmt.Sprintf("/?user_id=invalid&brand_id=%s", brandId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserBrandDeleteHandler(c)
		require.EqualError(t, err, "code=404, message=user not found")
	})
}

func TestAdminStaffHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminStaffHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		form.Insert("is_staff", "true")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: スタッフになれる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		form.Insert("is_staff", "true")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.NoError(t, err)

		existStaff, err := models.Staffs(
			models.StaffWhere.UserID.EQ(u2.ID),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.True(t, existStaff)
	})

	t.Run("成功: スタッフをはずせる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		ToStaff(t, ctx, &u)
		ToStaff(t, ctx, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		form.Insert("is_staff", "false")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.NoError(t, err)

		existStaff, err := models.Staffs(
			models.StaffWhere.UserID.EQ(u2.ID),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, existStaff)
	})

	t.Run("成功: すでにスタッフでもメモを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		ToStaff(t, ctx, &u)
		ToStaff(t, ctx, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		form.Insert("memo", "hogehoge")
		form.Insert("is_staff", "true")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.NoError(t, err)

		staff, err := models.Staffs(
			models.StaffWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, staff.Memo.String, "hogehoge", "メモが更新されている")
	})

	t.Run("成功: 自分自身のメモ更新はできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		form.Insert("memo", "hogehoge")
		form.Insert("is_staff", "true")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.NoError(t, err)

		staff, err := models.Staffs(
			models.StaffWhere.UserID.EQ(u.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, staff.Memo.String, "hogehoge", "メモが更新されている")
	})

	t.Run("失敗: ユーザーIDが不正な場合エラー", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", "invalid")
		form.Insert("is_staff", "true")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.EqualError(t, err, "code=404, message=user not found")
	})

	t.Run("失敗: 自分自身は変更できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		form.Insert("is_staff", "false")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.EqualError(t, err, "code=400, message=can't change yourself")
	})

	t.Run("失敗: スタッフを外そうとしたがスタッフではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		form.Insert("is_staff", "false")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.EqualError(t, err, "code=400, message=user is not staff")
	})
}

// TODO: ブロードキャスト通知は後で実装する
func TestAdminBroadcastHandler(t *testing.T) {
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBroadcastHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})
}

func TestAdminBrandHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBrandHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	RegisterBrand(t, ctx, "test", "")

	t.Run("成功: ブランドをすべて取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandHandler(c)
		require.NoError(t, err)

		response := []*models.Brand{}
		require.NoError(t, m.Json(&response))
		require.NotEqual(t, len(response), 0)
	})

	t.Run("成功: ブランドIDを指定して取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		brandId := RegisterBrand(t, ctx, "test", "")

		m, err := easy.NewMock(fmt.Sprintf("/?brand_id=%s", brandId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandHandler(c)
		require.NoError(t, err)

		response := []*models.Brand{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 1)
		require.Equal(t, response[0].ID, brandId)
	})

	t.Run("失敗: ブランドIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?brand_id=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandHandler(c)
		require.EqualError(t, err, "code=404, message=brand not found")
	})
}

func TestAdminBrandCreateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBrandCreateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		form := easy.NewMultipart()
		form.Insert("name", "test")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("ブランドを新規作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandCreateHandler(c)
		require.NoError(t, err)

		brand, err := models.Brands(
			models.BrandWhere.Name.EQ("test"),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, brand.Name, "test")
	})
}

func TestAdminBrandUpdateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBrandUpdateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		brandId := RegisterBrand(t, ctx, "test", "")

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("brand_id", brandId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("ブランドを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		brandId := RegisterBrand(t, ctx, "test", "")

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("brand_id", brandId)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandUpdateHandler(c)
		require.NoError(t, err)

		brand, err := models.Brands(
			models.BrandWhere.ID.EQ(brandId),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, brand.Name, "aaaaa")
	})

	t.Run("ブランドIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("brand_id", "invalid")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandUpdateHandler(c)
		require.EqualError(t, err, "code=404, message=brand not found")
	})
}

func TestAdminBrandDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBrandDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		brandId := RegisterBrand(t, ctx, "test", "")

		m, err := easy.NewMock(fmt.Sprintf("/?brand_id=%s", brandId), http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("ブランドを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		brandId := RegisterBrand(t, ctx, "test", "")

		m, err := easy.NewMock(fmt.Sprintf("/?brand_id=%s", brandId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminBrandDeleteHandler(c)
		require.NoError(t, err)

		existBrand, err := models.Brands(
			models.BrandWhere.ID.EQ(brandId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, existBrand)
	})
}

func TestAdminOrgHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("すべてのorgを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		for i := 0; i < 3; i++ {
			RegisterOrg(t, ctx)
		}

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgHandler(c)
		require.NoError(t, err)

		response := []models.Organization{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 3)
	})
}

func TestAdminOrgDetailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx)
		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgの詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx)

		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDetailHandler(c)
		require.NoError(t, err)

		response := src.OrgAdminDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.Org.ID, orgId)
		require.Len(t, response.Users, 1)

		require.Len(t, response.Clients, 1)
		require.Equal(t, response.Clients[0].ClientID, clientId)
	})

	t.Run("成功: orgユーザーがいない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDetailHandler(c)
		require.NoError(t, err)

		response := src.OrgAdminDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.Org.ID, orgId)
		require.Len(t, response.Users, 0)
	})

	t.Run("成功: clientはない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx)

		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDetailHandler(c)
		require.NoError(t, err)

		response := src.OrgAdminDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Clients, 0)
	})

	t.Run("失敗: org_idがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDetailHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", "invalid"), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDetailHandler(c)
		require.EqualError(t, err, "code=404, message=org not found")
	})
}

func TestAdminOrgCreateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgCreateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgを作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgCreateHandler(c)
		require.NoError(t, err)

		response := models.Organization{}
		require.NoError(t, m.Json(&response))

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(response.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, org.Name, "aaaaa")
		require.Equal(t, org.Link.String, "https://example.com")
		require.False(t, org.Image.Valid)
	})

	t.Run("成功: 画像を指定して作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgCreateHandler(c)
		require.NoError(t, err)

		response := models.Organization{}
		require.NoError(t, m.Json(&response))

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(response.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, org.Name, "aaaaa")
		require.Equal(t, org.Link.String, "https://example.com")

		// 画像のリンクが入っている
		require.True(t, org.Image.Valid)
	})

	t.Run("失敗: nameがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("link", "https://example.com")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgCreateHandler(c)
		require.EqualError(t, err, "code=400, message=name is required")
	})

	t.Run("失敗: linkのURLが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("link", "hogehoge")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgCreateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid link")
	})
}

func TestAdminOrgUpdateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgUpdateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com/aaaa")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgUpdateHandler(c)
		require.NoError(t, err)

		response := models.Organization{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ID, orgId)

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, org.Name, "aaaaa")
		require.Equal(t, org.Link.String, "https://example.com/aaaa")
		require.False(t, org.Image.Valid)
	})

	t.Run("成功: 画像を指定して更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com/aaaa")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgUpdateHandler(c)
		require.NoError(t, err)

		response := models.Organization{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ID, orgId)

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, org.Name, "aaaaa")
		require.Equal(t, org.Link.String, "https://example.com/aaaa")
		require.True(t, org.Image.Valid)
	})

	t.Run("失敗: org_idが無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com/aaaa")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		form.Insert("name", "aaaaa")
		form.Insert("link", "https://example.com/aaaa")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgUpdateHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: nameがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("link", "https://example.com/aaaa")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=name is required")
	})

	t.Run("失敗: linkのURLが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("name", "aaaaa")
		form.Insert("link", "aaaa")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid link")
	})
}

func TestAdminOrgDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodDelete, "")
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteHandler(c)
		require.NoError(t, err)

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, org)
	})

	t.Run("失敗: org_idが無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_id=aaaa", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})
}

func TestAdminOrgDeleteImageHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgDeleteImageHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx, u)

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		org.Image = null.NewString("https://example.com/aaaa", true)

		_, err = org.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodDelete, "")
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgの画像を削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		org.Image = null.NewString("https://example.com/aaaa", true)

		_, err = org.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteImageHandler(c)
		require.NoError(t, err)

		afterOrg, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(orgId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.False(t, afterOrg.Image.Valid)
	})

	t.Run("失敗: 画像を設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteImageHandler(c)
		require.EqualError(t, err, "code=400, message=image is not set")
	})

	t.Run("失敗: org_idが無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteImageHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_id=aaaaa", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgDeleteImageHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})
}

func TestAdminOrgMemberJoinHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgMemberJoinHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
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
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgにメンバーを追加できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

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

		err = h.AdminOrgMemberJoinHandler(c)
		require.NoError(t, err)

		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, orgUser.Role, "member")
	})

	t.Run("成功: 成功: emailでも招待できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

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

		err = h.AdminOrgMemberJoinHandler(c)
		require.NoError(t, err)

		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, orgUser.Role, "member")
	})

	t.Run("成功: roleを指定しないとguest", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

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

		err = h.AdminOrgMemberJoinHandler(c)
		require.NoError(t, err)

		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.UserID.EQ(u2.ID),
		).One(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, orgUser.Role, "guest")
	})

	t.Run("失敗: org_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberJoinHandler(c)
		require.EqualError(t, err, "code=400, message=org_id is required")
	})

	t.Run("失敗: org_idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", "invalid")
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberJoinHandler(c)
		require.EqualError(t, err, "code=404, message=organization not found")
	})

	t.Run("失敗: user_name_or_emailが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberJoinHandler(c)
		require.EqualError(t, err, "code=400, message=user_name_or_email is required")
	})

	t.Run("失敗: user_name_or_emailのユーザーが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", "invalid")
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberJoinHandler(c)
		require.EqualError(t, err, "code=404, message=user not found, unique=10")
	})

	t.Run("失敗: roleが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "aaaaa")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberJoinHandler(c)
		require.EqualError(t, err, "code=400, message=invalid role")
	})

	t.Run("失敗: ユーザーはすでにメンバーになっている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		InviteUserInOrg(t, ctx, orgId, &u2, "member")

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("org_id", orgId)
		form.Insert("user_name_or_email", u2.UserName)
		form.Insert("role", "member")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberJoinHandler(c)
		require.EqualError(t, err, "code=409, message=user already exists")
	})
}

func TestAdminOrgMemberRemoveHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminOrgMemberRemoveHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgUserId := InviteUserInOrg(t, ctx, orgId, &u2, "member")

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUserId), http.MethodPost, "")
		require.NoError(t, err)

		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: orgからユーザーを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		orgId := RegisterOrg(t, ctx)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgUserId := InviteUserInOrg(t, ctx, orgId, &u2, "member")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_user_id=%d", orgUserId), http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberRemoveHandler(c)
		require.NoError(t, err)

		existOrgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.ID.EQ(orgUserId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, existOrgUser)
	})

	t.Run("失敗: org_user_idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberRemoveHandler(c)
		require.EqualError(t, err, "code=400, message=org_user_id is required")
	})

	t.Run("失敗: org_user_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_user_id=invalid", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminOrgMemberRemoveHandler(c)
		require.EqualError(t, err, "code=400, message=invalid org_user_id")
	})
}

func TestAdminClientsHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminClientsHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: クライアント一覧を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminClientsHandler(c)
		require.NoError(t, err)

		response := []src.StaffClientResponse{}
		require.NoError(t, m.Json(&response))

		require.NotEqual(t, len(response), 0)
	})
}

func TestAdminClientDetailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminClientsHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		client := RegisterClient(t, ctx, u, "openid", "profile")

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", client.ClientID), http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: クライアントの詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		client := RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", client.ClientID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminClientDetailHandler(c)
		require.NoError(t, err)

		response := src.StaffClientDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.Client.ClientID, client.ClientID)
		require.Len(t, response.Scopes, 2) // openid, profile
	})

	t.Run("失敗: client_idがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminClientDetailHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: client_idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?client_id=invite", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminClientDetailHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

}

func TestAdminPreviewTemplateHTMLHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminPreviewTemplateHTMLHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/?name=test", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		c.SetParamNames("name")
		c.SetParamValues("test")
		return c
	})

	params := []string{
		"register",
		"register_resend",
		"update_email",
		"update_password",
		"invite_org",
		"test",
	}

	t.Run("アクセス可能", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		for _, param := range params {
			m, err := easy.NewMock("/:name", http.MethodGet, "")
			require.NoError(t, err)
			m.Cookie(cookie)

			c := m.Echo()
			c.SetParamNames("name")
			c.SetParamValues(param)

			err = h.AdminPreviewTemplateHTMLHandler(c)
			require.NoError(t, err, param)
		}
	})
}

func TestAdminRegisterSessionHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerRegisterSession := func(email string) {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:         session,
			Email:      email,
			VerifyCode: "123456",
			RetryCount: 0,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	StaffAndSessionTest(t, h.AdminRegisterSessionHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		registerEmail := RandomEmail(t)
		registerRegisterSession(registerEmail)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: セッション一覧を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		registerEmail := RandomEmail(t)
		registerRegisterSession(registerEmail)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminRegisterSessionHandler(c)
		require.NoError(t, err)

		response := []src.RegisterSessionResponse{}
		require.NoError(t, m.Json(&response))

		require.NotEqual(t, len(response), 0)
	})
}

func TestAdminDeleteRegisterSessionHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerRegisterSession := func(email string) {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:         session,
			Email:      email,
			VerifyCode: "123456",
			RetryCount: 0,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	StaffAndSessionTest(t, h.AdminDeleteRegisterSessionHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		registerEmail := RandomEmail(t)
		registerRegisterSession(registerEmail)

		m, err := easy.NewMock(fmt.Sprintf("/?email=%s", registerEmail), http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: セッションを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		registerEmail := RandomEmail(t)
		registerRegisterSession(registerEmail)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?email=%s", registerEmail), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminDeleteRegisterSessionHandler(c)
		require.NoError(t, err)

		exist, err := models.RegisterSessions(
			models.RegisterSessionWhere.Email.EQ(registerEmail),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, exist)
	})

	t.Run("失敗: メールアドレスが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		registerEmail := RandomEmail(t)
		registerRegisterSession(registerEmail)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminDeleteRegisterSessionHandler(c)
		require.EqualError(t, err, "code=400, message=email is required")
	})

	t.Run("失敗: そのメールアドレスのセッションが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		registerEmail := RandomEmail(t)
		registerRegisterSession(registerEmail)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?email=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()
		err = h.AdminDeleteRegisterSessionHandler(c)
		require.EqualError(t, err, "code=404, message=register session not found")
	})
}

func TestAdminUserNameHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminUserNameHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	}, func(c echo.Context) echo.Context {
		return c
	})

	t.Run("成功: 一覧を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		ToStaff(t, ctx, &u)

		userName, err := lib.RandomStr(10)
		require.NoError(t, err)

		// 最低1つが返ってくることを確認するため1つだけ追加しておく
		un := models.UserName{
			UserName: userName,
			UserID:   u.ID,
			Period:   time.Now().Add(h.C.UserNamePeriod),
		}
		err = un.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminUserNameHandler(c)
		require.NoError(t, err)

		response := []models.UserName{}
		require.NoError(t, m.Json(&response))

		require.NotEqual(t, len(response), 0)
	})
}
