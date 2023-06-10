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
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestAdminUsersHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminUsersHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 複数のユーザーを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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
	})

	t.Run("成功: 追加できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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
}

func TestAdminUserBrandDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminUserBrandDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		brandId := RegisterBrand(t, ctx, "test", "", u)

		m, err := easy.NewMock(fmt.Sprintf("/?user_id=%s&brand_id=%s", u.ID, brandId), http.MethodDelete, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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
		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: スタッフになれる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)
		staff2 := models.Staff{
			UserID: u2.ID,
		}
		err = staff2.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
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

	t.Run("失敗: ユーザーIDが不正な場合エラー", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_id", "invalid")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.AdminStaffHandler(c)
		require.EqualError(t, err, "code=404, message=user not found")
	})
}

func TestAdminBroadcastHandler(t *testing.T) {
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBroadcastHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})
}

func TestAdminBrandHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	StaffAndSessionTest(t, h.AdminBroadcastHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	RegisterBrand(t, ctx, "test", "")

	t.Run("成功: ブランドをすべて取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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
	})

	t.Run("ブランドを新規作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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
	})

	t.Run("ブランドを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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
	})

	t.Run("ブランドを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

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

func TestAdminOrgCreateHandler(t *testing.T) {

}

func TestAdminOrgUpdateHandler(t *testing.T) {

}

func TestAdminOrgDeleteHandler(t *testing.T) {

}

func TestAdminOrgDeleteImageHandler(t *testing.T) {

}
