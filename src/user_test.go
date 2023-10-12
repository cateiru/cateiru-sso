package src_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestUserMeHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserMeHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: ユーザ情報を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		// 設定追加
		setting := models.Setting{
			UserID:        u.ID,
			NoticeEmail:   false,
			NoticeWebpush: true,
		}
		require.NoError(t, setting.Insert(ctx, DB, boil.Infer()))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserMeHandler(c)
		require.NoError(t, err)

		response := src.UserMeResponse{}
		require.NoError(t, m.Json(&response))

		require.NotNil(t, response.UserInfo)
		require.Equal(t, response.UserInfo.ID, u.ID)
		require.Equal(t, response.UserInfo.UserName, u.UserName)
		require.Equal(t, response.UserInfo.Email, email)

		require.NotNil(t, response.Setting)
		require.Equal(t, response.Setting.UserID, u.ID)
		require.Equal(t, response.Setting.NoticeEmail, false)

		require.False(t, response.IsStaff)

		require.False(t, response.JoinedOrganization)
	})

	t.Run("成功: 設定がない場合は空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserMeHandler(c)
		require.NoError(t, err)

		response := src.UserMeResponse{}
		require.NoError(t, m.Json(&response))

		require.NotNil(t, response.UserInfo)
		require.Equal(t, response.UserInfo.ID, u.ID)
		require.Equal(t, response.UserInfo.UserName, u.UserName)
		require.Equal(t, response.UserInfo.Email, email)

		require.Nil(t, response.Setting)

		require.False(t, response.IsStaff)

		require.False(t, response.JoinedOrganization)
	})

	t.Run("スタッフの場合", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		staff := models.Staff{
			UserID: u.ID,
		}
		err := staff.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserMeHandler(c)
		require.NoError(t, err)

		response := src.UserMeResponse{}
		require.NoError(t, m.Json(&response))

		require.NotNil(t, response.UserInfo)
		require.Equal(t, response.UserInfo.ID, u.ID)
		require.Equal(t, response.UserInfo.UserName, u.UserName)
		require.Equal(t, response.UserInfo.Email, email)

		require.Nil(t, response.Setting)

		require.True(t, response.IsStaff)

		require.False(t, response.JoinedOrganization)
	})

	t.Run("組織に加入している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOrg(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		// 設定追加
		setting := models.Setting{
			UserID:        u.ID,
			NoticeEmail:   false,
			NoticeWebpush: true,
		}
		require.NoError(t, setting.Insert(ctx, DB, boil.Infer()))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserMeHandler(c)
		require.NoError(t, err)

		response := src.UserMeResponse{}
		require.NoError(t, m.Json(&response))

		require.NotNil(t, response.UserInfo)
		require.Equal(t, response.UserInfo.ID, u.ID)
		require.Equal(t, response.UserInfo.UserName, u.UserName)
		require.Equal(t, response.UserInfo.Email, email)

		require.NotNil(t, response.Setting)
		require.Equal(t, response.Setting.UserID, u.ID)
		require.Equal(t, response.Setting.NoticeEmail, false)

		require.False(t, response.IsStaff)

		require.True(t, response.JoinedOrganization)
	})
}

func TestUserUpdateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserUpdateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		form := easy.NewMultipart()
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: すべてのパラメータが変更できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		newUserName, err := lib.RandomStr(15)
		require.NoError(t, err)
		familyName, err := lib.RandomStr(5)
		require.NoError(t, err)
		middleName, err := lib.RandomStr(5)
		require.NoError(t, err)
		givenName, err := lib.RandomStr(5)
		require.NoError(t, err)
		gender := "1"                                           // 男性
		birthDate := time.Now().Add(-10 * 365 * 24 * time.Hour) // 現在時刻 - 10年
		locale := "en-US"

		form := easy.NewMultipart()
		form.Insert("user_name", newUserName)
		form.Insert("family_name", familyName)
		form.Insert("middle_name", middleName)
		form.Insert("given_name", givenName)
		form.Insert("gender", gender)
		form.Insert("birth_date", birthDate.Format(time.DateOnly))
		form.Insert("locale_id", locale)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.NoError(t, err)

		response := models.User{}
		require.NoError(t, m.Json(&response))
		require.Equal(t, response.UserName, newUserName)
		require.Equal(t, response.FamilyName.String, familyName)
		require.Equal(t, response.MiddleName.String, middleName)
		require.Equal(t, response.GivenName.String, givenName)
		require.Equal(t, response.Gender, gender)
		require.Equal(t, response.Birthdate.Time.Format(time.DateOnly), birthDate.Format(time.DateOnly))
		require.Equal(t, response.LocaleID, locale)

		dbUser, err := models.Users(
			models.UserWhere.ID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, dbUser.UserName, newUserName)
		require.Equal(t, dbUser.FamilyName.String, familyName)
		require.Equal(t, dbUser.MiddleName.String, middleName)
		require.Equal(t, dbUser.GivenName.String, givenName)
		require.Equal(t, dbUser.Gender, gender)
		require.Equal(t, dbUser.Birthdate.Time.Format(time.DateOnly), birthDate.Format(time.DateOnly))
		require.Equal(t, dbUser.LocaleID, locale)
	})

	t.Run("成功: family_name, middle_name, given_nameを指定しないと削除される", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		u.FamilyName = null.NewString("family", true)
		u.MiddleName = null.NewString("middle", true)
		u.GivenName = null.NewString("given", true)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		birthDate := time.Now().Add(-10 * 365 * 24 * time.Hour) // 現在時刻 - 10年

		form := easy.NewMultipart()
		form.Insert("family_name", "")
		form.Insert("middle_name", "")
		form.Insert("given_name", "")
		form.Insert("birth_date", birthDate.Format(time.DateOnly))
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.NoError(t, err)

		response := models.User{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.FamilyName.Valid)
		require.False(t, response.MiddleName.Valid)
		require.False(t, response.GivenName.Valid)

		dbUser, err := models.Users(
			models.UserWhere.ID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, dbUser.FamilyName.Valid)
		require.False(t, dbUser.MiddleName.Valid)
		require.False(t, dbUser.GivenName.Valid)
	})

	t.Run("成功: 誕生日を指定しないと削除される", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		birthDate := time.Now().Add(-10 * 365 * 24 * time.Hour) // 現在時刻 - 10年

		u.Birthdate = null.NewTime(birthDate, true)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("birth_date", "")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.NoError(t, err)

		response := models.User{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.Birthdate.Valid)

		dbUser, err := models.Users(
			models.UserWhere.ID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, dbUser.Birthdate.Valid)
	})

	t.Run("成功: 何も更新しなくても特に更新はしない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.NoError(t, err)
	})

	t.Run("成功: 自分のユーザー名は更新可能", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", u.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.NoError(t, err)
	})

	t.Run("成功: ユーザ名がuser_nameにすでに存在しているが持っている場合", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		userName, err := lib.RandomStr(5)
		require.NoError(t, err)

		dbUserName := models.UserName{
			UserName: userName,
			UserID:   u.ID,
			Period:   time.Now().Add(24 * time.Hour),
		}
		require.NoError(t, dbUserName.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", userName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.NoError(t, err)
	})

	t.Run("失敗: ユーザ名が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		newUserName, err := lib.RandomStr(10)
		require.NoError(t, err)

		newUserName += "+++" // 不正な文字を追加する

		form := easy.NewMultipart()
		form.Insert("user_name", newUserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid user_name")
	})

	t.Run("失敗: ユーザ名がすでに存在している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=user already exists, unique=12")
	})

	t.Run("失敗: ユーザ名がuser_nameにすでに存在している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		userName, err := lib.RandomStr(5)
		require.NoError(t, err)

		dbUserName := models.UserName{
			UserName: userName,
			UserID:   u2.ID,
			Period:   time.Now().Add(24 * time.Hour),
		}
		require.NoError(t, dbUserName.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", userName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=user already exists, unique=12")
	})

	t.Run("失敗: genderの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("gender", "8")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid gender")
	})

	t.Run("失敗: 誕生日が未来の値", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		birthDate := time.Now().Add(1 * 365 * 24 * time.Hour) // 現在時刻 + 1年

		form := easy.NewMultipart()
		form.Insert("birth_date", birthDate.Format(time.DateOnly))
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid birth_date")
	})

	t.Run("失敗: 誕生日の日付が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("birth_date", "2010-10-2Z10:00:00")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid birth_date")
	})

	t.Run("失敗: locale_idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("locale_id", "xx-YY-ZZ")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=invalid locale_id")
	})
}

func TestUserUserNameHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserUserNameHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		userName, err := lib.RandomStr(15)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("user_name", userName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("ユーザ名が存在する", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", u2.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUserNameHandler(c)
		require.NoError(t, err)

		response := src.UserUserNameResponse{}
		require.NoError(t, m.Json(&response))

		require.False(t, response.Ok)
		require.Equal(t, response.UserName, u2.UserName)
		require.Equal(t, response.Message, "ユーザー名は既に使用されています")
	})

	t.Run("自分のユーザー名はOK", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", u.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUserNameHandler(c)
		require.NoError(t, err)

		response := src.UserUserNameResponse{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.Ok)
		require.Equal(t, response.UserName, u.UserName)
		require.Equal(t, response.Message, "ユーザー名は使用可能です")
	})

	t.Run("ユーザ名が存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		userName, err := lib.RandomStr(15)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("user_name", userName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUserNameHandler(c)
		require.NoError(t, err)

		response := src.UserUserNameResponse{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.Ok)
		require.Equal(t, response.UserName, userName)
		require.Equal(t, response.Message, "ユーザー名は使用可能です")
	})

	t.Run("ユーザー名は存在しないが、不正な値", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", "a")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUserNameHandler(c)
		require.NoError(t, err)

		response := src.UserUserNameResponse{}
		require.NoError(t, m.Json(&response))

		require.False(t, response.Ok)
		require.Equal(t, response.UserName, "a")
		require.Equal(t, response.Message, "ユーザー名は3文字以上15文字以下で半角英数字と'_'のみ使用できます")
	})

	t.Run("user_nameに含まれている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		userName, err := lib.RandomStr(5)
		require.NoError(t, err)

		dbUserName := models.UserName{
			UserName: userName,
			UserID:   u2.ID,
			Period:   time.Now().Add(24 * time.Hour),
		}
		require.NoError(t, dbUserName.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", userName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUserNameHandler(c)
		require.NoError(t, err)

		response := src.UserUserNameResponse{}
		require.NoError(t, m.Json(&response))

		require.False(t, response.Ok)
		require.Equal(t, response.UserName, userName)
		require.Equal(t, response.Message, "このユーザー名は使用できません")
	})

	t.Run("user_nameに含まれているが自分で作ったものなのでOK", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		userName, err := lib.RandomStr(5)
		require.NoError(t, err)

		dbUserName := models.UserName{
			UserName: userName,
			UserID:   u.ID,
			Period:   time.Now().Add(24 * time.Hour),
		}
		require.NoError(t, dbUserName.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("user_name", userName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUserNameHandler(c)
		require.NoError(t, err)

		response := src.UserUserNameResponse{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.Ok)
		require.Equal(t, response.UserName, userName)
		require.Equal(t, response.Message, "ユーザー名は使用可能です")
	})
}

func TestUserUpdateSettingHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserUpdateSettingHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		form := easy.NewMultipart()
		form.Insert("notice_email", "true")
		form.Insert("notice_webpush", "false")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: 設定を新規作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("notice_email", "true")
		form.Insert("notice_webpush", "false")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateSettingHandler(c)
		require.NoError(t, err)

		response := models.Setting{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.NoticeEmail)
		require.False(t, response.NoticeWebpush)

		setting, err := models.Settings(
			models.SettingWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.True(t, setting.NoticeEmail)
		require.False(t, setting.NoticeWebpush)
	})

	t.Run("成功: 設定を更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		settingModel := models.Setting{
			UserID:        u.ID,
			NoticeEmail:   true,
			NoticeWebpush: true,
		}
		_, err := settingModel.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("notice_email", "true")
		form.Insert("notice_webpush", "false")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateSettingHandler(c)
		require.NoError(t, err)

		response := models.Setting{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.NoticeEmail)
		require.False(t, response.NoticeWebpush)

		setting, err := models.Settings(
			models.SettingWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.True(t, setting.NoticeEmail)
		require.False(t, setting.NoticeWebpush)
	})
}

func TestUserBrandHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserBrandHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: ブランドを取得できる", func(t *testing.T) {
		t.Run("ブランドが指定されていない", func(t *testing.T) {
			email := RandomEmail(t)
			u := RegisterUser(t, ctx, email)

			cookies := RegisterSession(t, ctx, &u)

			m, err := easy.NewMock("/", http.MethodGet, "")
			require.NoError(t, err)
			m.Cookie(cookies)

			c := m.Echo()

			err = h.UserBrandHandler(c)
			require.NoError(t, err)

			response := src.UserBrandResponse{}
			require.NoError(t, m.Json(&response))

			require.Len(t, response.BrandNames, 0)
		})

		t.Run("ブランドは設定されている", func(t *testing.T) {
			email := RandomEmail(t)
			u := RegisterUser(t, ctx, email)

			brandId, err := lib.RandomStr(31)
			require.NoError(t, err)
			brand := models.Brand{
				ID:   brandId,
				Name: "pro",
			}
			err = brand.Insert(ctx, DB, boil.Infer())
			require.NoError(t, err)
			userBrand := models.UserBrand{
				UserID:  u.ID,
				BrandID: brandId,
			}
			err = userBrand.Insert(ctx, DB, boil.Infer())
			require.NoError(t, err)

			cookies := RegisterSession(t, ctx, &u)

			m, err := easy.NewMock("/", http.MethodGet, "")
			require.NoError(t, err)
			m.Cookie(cookies)

			c := m.Echo()

			err = h.UserBrandHandler(c)
			require.NoError(t, err)

			response := src.UserBrandResponse{}
			require.NoError(t, m.Json(&response))

			require.Len(t, response.BrandNames, 1)
			require.Equal(t, response.BrandNames[0], "pro")
		})
	})
}

func TestUserUpdateEmailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserUpdateEmailHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		newEmail := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("new_email", newEmail)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: メールアドレスに更新メールが送られる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("new_email", newEmail)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailHandler(c)
		require.NoError(t, err)

		response := src.UserUpdateEmailResponse{}
		require.NoError(t, m.Json(&response))

		session, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(response.Session),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, session.UserID, u.ID)
		require.Equal(t, session.NewEmail, newEmail)
	})

	t.Run("失敗: メールアドレスが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailHandler(c)
		require.EqualError(t, err, "code=400, message=empty new email")
	})

	t.Run("失敗: メールアドレスはすでに別のユーザが使用している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		RegisterUser(t, ctx, newEmail)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("new_email", newEmail)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailHandler(c)
		require.EqualError(t, err, "code=400, message=email already used, unique=12")
	})

	t.Run("失敗: メールアドレスが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("new_email", "hogehoge")
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailHandler(c)
		require.EqualError(t, err, "code=400, message=empty new email")
	})

	t.Run("失敗: reCAPTCHAが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("new_email", newEmail)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA token is empty")
	})

	t.Run("失敗: reCAPTCHAが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("new_email", newEmail)
		form.Insert("recaptcha", "fail")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})
}

func TestUserUpdateEmailRegisterHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerEmailSession := func(u *models.User, newEmail string) (string, string) {

		sessionId, err := lib.RandomStr(31)
		require.NoError(t, err)
		code, err := lib.RandomNumber(6)
		require.NoError(t, err)

		session := models.EmailVerifySession{
			ID:         sessionId,
			UserID:     u.ID,
			NewEmail:   newEmail,
			VerifyCode: code,
			Period:     time.Now().Add(h.C.UpdateEmailSessionPeriod),
		}
		err = session.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		return sessionId, code
	}

	SessionTest(t, h.UserUpdateEmailRegisterHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		newEmail := RandomEmail(t)
		session, code := registerEmailSession(u, newEmail)

		form := easy.NewMultipart()
		form.Insert("update_token", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: メールアドレスを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		session, code := registerEmailSession(&u, newEmail)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("update_token", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailRegisterHandler(c)
		require.NoError(t, err)

		response := src.UserUpdateEmailRegisterResponse{}
		require.NoError(t, m.Json(&response))
		require.Equal(t, response.Email, newEmail)

		newUser, err := models.Users(
			models.UserWhere.ID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, newUser.Email, newEmail)

		// セッションは削除されている
		existSession, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(session),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existSession)
	})

	t.Run("失敗: codeが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		session, _ := registerEmailSession(&u, newEmail)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("update_token", session)
		form.Insert("code", "12345")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailRegisterHandler(c)
		require.EqualError(t, err, "code=403, message=invalid code, unique=13")

		// リトライカウント++されている
		se, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, se.RetryCount, uint8(1))
	})

	t.Run("失敗: リトライ上限", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		session, code := registerEmailSession(&u, newEmail)

		// リトライ上限にする
		se, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		se.RetryCount = C.UpdateEmailRetryCount
		_, err = se.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("update_token", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailRegisterHandler(c)
		require.EqualError(t, err, "code=403, message=exceeded retry, unique=4")

		// セッションは削除されている
		existSession, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(session),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existSession)
	})

	t.Run("失敗: セッションが無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		_, code := registerEmailSession(&u, newEmail)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailRegisterHandler(c)
		require.EqualError(t, err, "code=400, message=update_token is empty")
	})

	t.Run("失敗: セッションが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		_, code := registerEmailSession(&u, newEmail)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("update_token", "hogehoge")
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailRegisterHandler(c)
		require.EqualError(t, err, "code=400, message=invalid session")
	})

	t.Run("失敗: セッションの有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		newEmail := RandomEmail(t)
		session, code := registerEmailSession(&u, newEmail)

		// セッションの有効期限をきらす
		se, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		se.Period = time.Now().Add(-10 * 24 * time.Hour)
		_, err = se.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("update_token", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserUpdateEmailRegisterHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")

		// セッションは削除されている
		existSession, err := models.EmailVerifySessions(
			models.EmailVerifySessionWhere.ID.EQ(session),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existSession)
	})
}

func TestUserAvatarHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.UserAvatarHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form := easy.NewMultipart()
		form.InsertFile("image", image)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: アバターを新規作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form := easy.NewMultipart()
		form.InsertFile("image", image)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserAvatarHandler(c)
		require.NoError(t, err)

		dbUser, err := models.Users(
			models.UserWhere.ID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		path := filepath.Join("avatar", u.ID)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}
		require.Equal(t, dbUser.Avatar.String, url.String())

		time.Sleep(1 * time.Second)

		storage := lib.NewCloudStorage(C.StorageBucketName)
		_, contentType, err := storage.Read(ctx, path)
		require.NoError(t, err)
		require.Equal(t, contentType, "image/png")
	})

	t.Run("成功: アバターを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		u.Avatar = null.NewString("https://example.com/avatar", true)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form := easy.NewMultipart()
		form.InsertFile("image", image)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserAvatarHandler(c)
		require.NoError(t, err)

		time.Sleep(1 * time.Second)

		path := filepath.Join("avatar", u.ID)
		storage := lib.NewCloudStorage(C.StorageBucketName)
		_, contentType, err := storage.Read(ctx, path)
		require.NoError(t, err)
		require.Equal(t, contentType, "image/png")
	})

	t.Run("失敗: 画像が指定されていない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserAvatarHandler(c)
		require.EqualError(t, err, "code=400, message=http: no such file")
	})

	t.Run("失敗: 画像ファイルじゃない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		image, err := os.Open("./user_test.go")
		require.NoError(t, err)
		defer image.Close()
		form := easy.NewMultipart()
		form.InsertFile("image", image)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserAvatarHandler(c)
		require.EqualError(t, err, "code=400, message=invalid Content-Type")
	})
}

func TestUserDeleteAvatarHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	s := lib.NewCloudStorage(C.StorageBucketName)

	SessionTest(t, h.UserDeleteAvatarHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		path := filepath.Join("avatar", u.ID)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		u.Avatar = null.NewString(url.String(), true)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: アバターが削除されている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		path := filepath.Join("avatar", u.ID)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		u.Avatar = null.NewString(url.String(), true)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserDeleteAvatarHandler(c)
		require.NoError(t, err)

		time.Sleep(1 * time.Second)

		_, _, err = s.Read(ctx, path)
		require.ErrorIs(t, err, storage.ErrObjectNotExist)
	})

	t.Run("失敗: アバターは設定されていない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.UserDeleteAvatarHandler(c)
		require.EqualError(t, err, "code=400, message=avatar is not set")
	})
}

// TODO: クライアントの実装してから
func TestUserLogoutClientHandler(t *testing.T) {
	t.Run("成功: 指定したクライアントからログアウトできる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが空", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: そもそもクライアントIDのクライアントにログインしていない", func(t *testing.T) {})
}

func TestRegisterUser(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		u, err := src.RegisterUser(ctx, DB, C, email)
		require.NoError(t, err)

		require.Equal(t, u.Email, email)
		require.Len(t, u.UserName, 8)

		existUser, err := models.UserExists(ctx, DB, u.ID)
		require.NoError(t, err)
		require.True(t, existUser)

		existStaff, err := models.StaffExists(ctx, DB, u.ID)
		require.NoError(t, err)
		require.False(t, existStaff)
	})

	t.Run("メールアドレスのドメインが一致するとスタッフになる", func(t *testing.T) {
		ctx := context.Background()

		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@example.test", r) // ドメインを `.test` にする

		u, err := src.RegisterUser(ctx, DB, C, email)
		require.NoError(t, err)

		require.Equal(t, u.Email, email)
		require.Len(t, u.UserName, 8)

		existUser, err := models.UserExists(ctx, DB, u.ID)
		require.NoError(t, err)
		require.True(t, existUser)

		existStaff, err := models.StaffExists(ctx, DB, u.ID)
		require.NoError(t, err)
		require.True(t, existStaff)
	})

	t.Run("すでにEmailが存在している場合はエラー", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		RegisterUser(t, ctx, email)

		_, err := src.RegisterUser(ctx, DB, C, email)
		require.EqualError(t, err, "code=400, message=impossible register account, unique=3")
	})
}

func TestFindUserByUserNameOrEmail(t *testing.T) {
	ctx := context.Background()

	t.Run("成功: ユーザー名", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		user, err := src.FindUserByUserNameOrEmail(ctx, DB, u.UserName)
		require.NoError(t, err)

		require.Equal(t, user.ID, u.ID)
	})

	t.Run("成功: Email", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		user, err := src.FindUserByUserNameOrEmail(ctx, DB, u.Email)
		require.NoError(t, err)

		require.Equal(t, user.ID, u.ID)
	})

	t.Run("失敗: ユーザーが存在しない", func(t *testing.T) {
		_, err := src.FindUserByUserNameOrEmail(ctx, DB, "aaaaaa")
		require.EqualError(t, err, "code=404, message=user not found, unique=10")
	})
}
