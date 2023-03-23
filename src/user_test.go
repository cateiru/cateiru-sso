package src_test

import (
	"context"
	"net/http"
	"testing"
	"time"

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

		require.False(t, response.Ok)
		require.Equal(t, response.UserName, u.UserName)
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
	t.Run("成功: ブランドを取得できる", func(t *testing.T) {
		t.Run("ブランドが指定されている", func(t *testing.T) {})

		t.Run("ブランドは設定されていない", func(t *testing.T) {})
	})
}

func TestUserUpdateEmailHandler(t *testing.T) {
	t.Run("成功: メールアドレスに更新メールが送られる", func(t *testing.T) {})

	t.Run("失敗: メールアドレスが空", func(t *testing.T) {})

	t.Run("失敗: メールアドレスはすでに別のユーザが使用している", func(t *testing.T) {})

	t.Run("失敗: メールアドレスが不正", func(t *testing.T) {})
}

func TestUserUpdateEmailRegisterHandler(t *testing.T) {

	t.Run("成功: メールアドレスを更新できる", func(t *testing.T) {})

	t.Run("失敗: セッションが無い", func(t *testing.T) {})

	t.Run("失敗: セッションが不正", func(t *testing.T) {})

	t.Run("失敗: セッションの有効期限切れ", func(t *testing.T) {})
}

func TestUserAvatarHandler(t *testing.T) {
	t.Run("成功: アバターを新規作成できる", func(t *testing.T) {})

	t.Run("成功: アバターを更新できる", func(t *testing.T) {})

	t.Run("失敗: 画像が指定されていない", func(t *testing.T) {})
}

func TestUserDeleteAvatarHandler(t *testing.T) {
	t.Run("成功: アバターが削除されている", func(t *testing.T) {})
}

// TODO: クライアントの実装してから
func TestUserLogoutClient(t *testing.T) {
	t.Run("成功: 指定したクライアントからログアウトできる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが空", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: そもそもクライアントIDのクライアントにログインしていない", func(t *testing.T) {})
}

func TestRegisterUser(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		u, err := src.RegisterUser(ctx, DB, email)
		require.NoError(t, err)

		require.Equal(t, u.Email, email)
		require.Len(t, u.UserName, 8)
	})

	t.Run("すでにEmailが存在している場合はエラー", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		RegisterUser(t, ctx, email)

		_, err := src.RegisterUser(ctx, DB, email)
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

	t.Run("失敗", func(t *testing.T) {
		_, err := src.FindUserByUserNameOrEmail(ctx, DB, "aaaaaa")
		require.Error(t, err)
	})
}
