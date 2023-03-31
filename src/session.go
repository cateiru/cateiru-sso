package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SessionInterface interface {
	SimpleLogin(ctx context.Context, c echo.Context, useSessionOnly ...bool) (*models.User, error)
	Login(ctx context.Context, cookies []*http.Cookie, useSessionOnly ...bool) (*models.User, []*http.Cookie, error)
	Logout(ctx context.Context, cookies []*http.Cookie, user *models.User) ([]*http.Cookie, error)
	NewRegisterSession(ctx context.Context, user *models.User, ua *UserData, ip string) (*RegisterSession, error)
	SwitchAccount(ctx context.Context, cookies []*http.Cookie, userID string) ([]*http.Cookie, error)
	LoggedInAccounts(ctx context.Context, cookies []*http.Cookie) ([]*models.User, error)
}

type RegisterSession struct {
	SessionToken string
	RefreshToken string
	UserID       string
}

type Session struct {
	SessionCookie    CookieConfig
	RefreshCookie    CookieConfig
	LoginUserCookie  CookieConfig
	LoginStateCookie CookieConfig
	SessionDBPeriod  time.Duration
	RefreshDBPeriod  time.Duration

	DB *sql.DB
}

func NewSession(c *Config, db *sql.DB) *Session {
	return &Session{
		SessionCookie:    c.SessionCookie,
		RefreshCookie:    c.RefreshCookie,
		LoginUserCookie:  c.LoginUserCookie,
		LoginStateCookie: c.LoginStateCookie,
		SessionDBPeriod:  c.SessionDBPeriod,
		RefreshDBPeriod:  c.RefreshDBPeriod,
		DB:               db,
	}
}

func (s *Session) SimpleLogin(ctx context.Context, c echo.Context, useSessionOnly ...bool) (*models.User, error) {
	sessionOnlyFlag := false
	if len(useSessionOnly) >= 1 {
		sessionOnlyFlag = useSessionOnly[0]
	}

	user, setCookies, err := s.Login(ctx, c.Cookies(), sessionOnlyFlag)
	if sessionOnlyFlag {
		for _, cookie := range setCookies {
			c.SetCookie(cookie)
		}
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ログインする
// 通常、セッショントークンを使用してログインを試みます。
// セッショントークンが存在しないor有効期限が切れている場合、リフレッシュトークンを使用してログインを試みます。
// リフレッシュトークンを使用してログインした場合、リフレッシュトークンの値は更新されます。
// エラー時にもcookieはが存在する可能性があるためset-cookiesする必要があります
//
// ```mermaid
// graph TD
// A([start]) --> B{セッショントークンがあるか}
// B -->|Yes| C[DBからセッショントークンを取得する]
// B -->|No| REFRESH([リフレッシュトークンからログイン])
// REFRESH --> H{ログインユーザCookieが存在するか}
// H -->|Yes| I[リフレッシュトークンを取得する]
// I --> M{リフレッシュトークンが存在するか}
// M --> |Yes| N[DBからリフレッシュを取得する]
// N --> O{リフレッシュがDBから取得できた}
// O -->|Yes| P{DBから取得したリフレッシュのユーザIDとリフレッシュCookieのユーザIDが同じ}
// P -->|Yes| Q{リフレッシュの有効期限が切れていない}
// Q -->|Yes| R[DBからユーザを取得する]
// R --> S{ユーザがDBから取得できた}
// S -->|Yes| T[前のリフレッシュトークンをDBから削除]
// T --> U[リフレッシュに紐付けられているセッションを削除]
// U --> V[リフレッシュトークンを新規作成]
// V --> W[セッショントークンを新規作成]
// W --> END
// S -->|No| ERR([ログイン失敗処理])
// Q -->|No| ERR
// P -->|No| ERR
// O -->|No| ERR
// M --> |No| ERR
// H -->|No| J{リフレッシュトークンが存在するか}
// J -->|Yes| K[他のアカウントでログインできる可能性エラーで返す]
// K --> END
// J -->|No| ERR
// C --> D{セッショントークンがDBから取得できた}
// D -->|Yes| E{セッショントークンの有効期限が切れていない}
// D -->|No| REFRESH
// E -->|Yes| F[ユーザをDBから取得する]
// E -->|No| REFRESH
// F --> G{ユーザがDBから取得できた}
// G -->|Yes| Z[ログイン完了]
// G -->|No| ERR
// ERR --> ERRA[セッショントークンがCookieにあればDBとCookieを削除]
// ERRA --> ERRB[リフレッシュトークンがCookieにあればDBとCookieを削除]
// ERRB --> ERRC[LoginUser cookieがあればcookieを削除]
// ERRC --> ERRD[LoginState CookieがあればCookieを削除]
// ERRD --> END
// Z --> END([end])
// ```
//
// https://mermaid.ink/img/pako:eNqdVm1P2lAU_ivkftIEDRBQ4MMSoSjqRAVfMgsfGqhKJtRASeaQZG0XFcXIlvmWkJgsDpzOl2XLls25_ZhLUf_Fbm9bWqDCXENIuZznec45Pc_tzYIoE6OBGyykqOVF0xQRTprQNdBFplkqxUa6TT09T0yeLOSvoSBA_jsUKlD4AoUNKPyC_KV0zxUgx0N-C3JbORnukVCrz-j0qslLEh70B-Tz7Sj4t-LOnvh7H3KHiCiiZwkwq6agbzDoC_m7SCicQmEXCp9Upg-NmWAd4RzyV5A_RiuRbplKIcDF-LP6CCiUMf6bl2Gex2lUjHh-IJZO5Ey0kvxaScOd0jCsZhiLj2U7lfCA_hiGywkEtJ62khmKBzB63EicK6hkKq4CuW3IHSnC41rhE9nmUPQ5MuK8qLd1mIDcSWuI2u3GwIJYRLN0oChPaMqThplf1Ep5cfNHrXR0f1iU0BvrkEcEZci9htwp-laYJjWmoK53irJxy4K4ZaGsFtapUyFNZYoU89u4ug6DUmfMb94fHivSU1h62mjMuLPbr8Xq9QHk3mBLqdWiUWmyFypKzzmNOWc6j25t7-quvFO9KdU2igp2BmNnyfYONgDOYqAvQOgbJBnaFwxKZtZsKB5_ru3ui-vl2-KaatpJfbh-JBqWxluXFK80rPnrYSP_a8ER7fmOktXrPen58u8hfwb5stIONBH6vUUeEH5L3Lm8E25qryqQR174KElxlbs_75CC0qnRxk6NtBblxRFEh524_YQSWgW-9kT_bC2ieZOWl32a0iCpd1pzgjrH-YypBnHhQ48w4pAmPtc4ZReF6s-1iD6qocfoRn4OweAA2b7R6gZ2ht99qDNXUk4nynKT-SQ-ldjT8SX2aG6Pyu0lnzIL8eR0mk6Zour7rE4SfQjvVfGEjA-xFEubvC0EDyZANE7vnPqzi6STsUg3MIMEnUpQ8Rg6aGSlmDBgF-kEHQZudBuj56nMEhsG4WQOhVIZlgmtJKPAzaYytBlklmMoHSJOoSNKArjnqaU0Wl2mksCdBS-A22619VocNnT12-1Wq8vhMIMV4O5z9tosLrvd7rA4XS6L3ZYzg5cMgxgsvU6H1erodzosfTanw4bi6VicZVJj8lEIn4iwwhyOx4q5vzI6g_c?type=png)](https://mermaid.live/edit#pako:eNqdVm1P2lAU_ivkftIEDRBQ4MMSoSjqRAVfMgsfGqhKJtRASeaQZG0XFcXIlvmWkJgsDpzOl2XLls25_ZhLUf_Fbm9bWqDCXENIuZznec45Pc_tzYIoE6OBGyykqOVF0xQRTprQNdBFplkqxUa6TT09T0yeLOSvoSBA_jsUKlD4AoUNKPyC_KV0zxUgx0N-C3JbORnukVCrz-j0qslLEh70B-Tz7Sj4t-LOnvh7H3KHiCiiZwkwq6agbzDoC_m7SCicQmEXCp9Upg-NmWAd4RzyV5A_RiuRbplKIcDF-LP6CCiUMf6bl2Gex2lUjHh-IJZO5Ey0kvxaScOd0jCsZhiLj2U7lfCA_hiGywkEtJ62khmKBzB63EicK6hkKq4CuW3IHSnC41rhE9nmUPQ5MuK8qLd1mIDcSWuI2u3GwIJYRLN0oChPaMqThplf1Ep5cfNHrXR0f1iU0BvrkEcEZci9htwp-laYJjWmoK53irJxy4K4ZaGsFtapUyFNZYoU89u4ug6DUmfMb94fHivSU1h62mjMuLPbr8Xq9QHk3mBLqdWiUWmyFypKzzmNOWc6j25t7-quvFO9KdU2igp2BmNnyfYONgDOYqAvQOgbJBnaFwxKZtZsKB5_ru3ui-vl2-KaatpJfbh-JBqWxluXFK80rPnrYSP_a8ER7fmOktXrPen58u8hfwb5stIONBH6vUUeEH5L3Lm8E25qryqQR174KElxlbs_75CC0qnRxk6NtBblxRFEh524_YQSWgW-9kT_bC2ieZOWl32a0iCpd1pzgjrH-YypBnHhQ48w4pAmPtc4ZReF6s-1iD6qocfoRn4OweAA2b7R6gZ2ht99qDNXUk4nynKT-SQ-ldjT8SX2aG6Pyu0lnzIL8eR0mk6Zour7rE4SfQjvVfGEjA-xFEubvC0EDyZANE7vnPqzi6STsUg3MIMEnUpQ8Rg6aGSlmDBgF-kEHQZudBuj56nMEhsG4WQOhVIZlgmtJKPAzaYytBlklmMoHSJOoSNKArjnqaU0Wl2mksCdBS-A22619VocNnT12-1Wq8vhMIMV4O5z9tosLrvd7rA4XS6L3ZYzg5cMgxgsvU6H1erodzosfTanw4bi6VicZVJj8lEIn4iwwhyOx4q5vzI6g_c
func (s *Session) Login(ctx context.Context, cookies []*http.Cookie, useSessionOnly ...bool) (*models.User, []*http.Cookie, error) {
	sessionOnlyFlag := false
	if len(useSessionOnly) >= 1 {
		sessionOnlyFlag = useSessionOnly[0]
	}

	var sessionCookie *http.Cookie = nil
	for _, cookie := range cookies {
		if cookie.Name == s.SessionCookie.Name {
			sessionCookie = cookie
			break
		}
	}
	if sessionCookie == nil {
		// リフレッシュトークンでログインを試みる
		return s.loginWithRefresh(ctx, cookies, sessionOnlyFlag)

	}

	session, err := models.Sessions(
		models.SessionWhere.ID.EQ(sessionCookie.Value),
	).One(ctx, s.DB)
	if errors.Is(err, sql.ErrNoRows) {
		if sessionOnlyFlag {
			return nil, []*http.Cookie{}, NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "login failed")
		} else {
			// リフレッシュトークンでログインを試みる
			return s.loginWithRefresh(ctx, cookies, sessionOnlyFlag)
		}
	}
	if err != nil {
		return nil, []*http.Cookie{}, err
	}
	// 有効期限が切れている
	if time.Now().After(session.Period) {
		// リフレッシュトークンでログインを試みる
		return s.loginWithRefresh(ctx, cookies, sessionOnlyFlag)
	}

	// ユーザDB叩く
	u, err := models.Users(
		models.UserWhere.ID.EQ(session.UserID),
	).One(ctx, s.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return s.loginFailed(ctx, cookies, "")
	}
	if err != nil {
		return nil, []*http.Cookie{}, err
	}

	return u, []*http.Cookie{}, nil
}

// リフレッシュトークンを使用してログインを試みる
func (s *Session) loginWithRefresh(ctx context.Context, cookies []*http.Cookie, sessionOnly bool) (*models.User, []*http.Cookie, error) {
	// sessionOnlyがtrueのときは、リフレッシュトークンを使用したログインは行わないでエラーにする
	// これは、アカウント変更やログアウトなどでCookieが更新されると良くないことが起きるためである
	if sessionOnly {
		return s.loginFailed(ctx, cookies, "")
	}

	var loginUserId *http.Cookie = nil
	refreshTokensCount := 0
	for _, cookie := range cookies {
		if cookie.Name == s.LoginUserCookie.Name {
			loginUserId = cookie
		}
		if strings.HasPrefix(cookie.Name, s.RefreshCookie.Name) {
			refreshTokensCount++
		}
	}
	if loginUserId == nil {
		// 他のアカウントでログインできる可能性がある場合
		if refreshTokensCount > 0 {
			return nil, []*http.Cookie{}, NewHTTPUniqueError(
				http.StatusForbidden,
				ErrBeAbleToLoginWithAnotherAccount,
				"you may be able to log in with another account",
			)
		}
		// ログイン失敗
		return s.loginFailed(ctx, cookies, "")
	}
	// リフレッシュトークンを取得する
	refreshTokenName := fmt.Sprintf("%s-%s", s.RefreshCookie.Name, loginUserId.Value)
	var refreshCookie *http.Cookie = nil
	for _, cookie := range cookies {
		if cookie.Name == refreshTokenName {
			refreshCookie = cookie
			break
		}
	}
	if refreshCookie == nil {
		// ログイン失敗
		return s.loginFailed(ctx, cookies, "")
	}

	refresh, err := models.Refreshes(
		models.RefreshWhere.ID.EQ(refreshCookie.Value),
	).One(ctx, s.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return s.loginFailed(ctx, cookies, refreshTokenName)
	}
	if err != nil {
		return nil, []*http.Cookie{}, err
	}
	// ユーザが不正な場合
	if string(refresh.UserID) != loginUserId.Value {
		return s.loginFailed(ctx, cookies, refreshTokenName)
	}
	// 有効期限が切れている
	if time.Now().After(refresh.Period) {
		return s.loginFailed(ctx, cookies, refreshTokenName)
	}

	u, err := models.Users(
		models.UserWhere.ID.EQ(refresh.UserID),
	).One(ctx, s.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return s.loginFailed(ctx, cookies, refreshTokenName)
	}
	if err != nil {
		return nil, []*http.Cookie{}, err
	}

	// 前のリフレッシュトークンは削除してしまう
	if _, err := refresh.Delete(ctx, s.DB); err != nil {
		return nil, []*http.Cookie{}, err
	}
	// リフレッシュトークンにセッショントークンが紐付けられている場合は、セッショントークンを削除する
	if refresh.SessionID.Valid {
		if _, err := models.Sessions(
			models.SessionWhere.ID.EQ(refresh.SessionID.String),
		).DeleteAll(ctx, s.DB); err != nil {
			return nil, []*http.Cookie{}, err
		}
	}

	// リフレッシュトークンを更新し、セッショントークンを新規作成する
	newSessionToken, err := lib.RandomStr(31)
	if err != nil {
		return nil, []*http.Cookie{}, err
	}
	newRefreshToken, err := lib.RandomStr(63)
	if err != nil {
		return nil, []*http.Cookie{}, err
	}

	newSession := models.Session{
		ID:     newSessionToken,
		UserID: u.ID,

		Period: time.Now().Add(s.SessionDBPeriod),
	}
	if err := newSession.Insert(ctx, s.DB, boil.Infer()); err != nil {
		return nil, []*http.Cookie{}, err
	}
	newRefresh := models.Refresh{
		ID:        newRefreshToken,
		UserID:    u.ID,
		HistoryID: refresh.HistoryID, // history_id は引き継ぐ
		SessionID: null.NewString(newSessionToken, true),

		Period: time.Now().Add(s.RefreshDBPeriod),
	}
	if err := newRefresh.Insert(ctx, s.DB, boil.Infer()); err != nil {
		return nil, []*http.Cookie{}, err
	}

	// 新しいCookie設定
	newSessionCookie := &http.Cookie{
		Name:     s.SessionCookie.Name,
		Secure:   s.SessionCookie.Secure,
		HttpOnly: s.SessionCookie.HttpOnly,
		Path:     s.SessionCookie.Path,
		MaxAge:   s.SessionCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(s.SessionCookie.MaxAge) * time.Second),
		SameSite: s.SessionCookie.SameSite,

		Value: newSessionToken,
	}
	newRefreshCookie := &http.Cookie{
		Name:     refreshTokenName,
		Secure:   s.RefreshCookie.Secure,
		HttpOnly: s.RefreshCookie.HttpOnly,
		Path:     s.RefreshCookie.Path,
		MaxAge:   s.RefreshCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(s.RefreshCookie.MaxAge) * time.Second),
		SameSite: s.RefreshCookie.SameSite,

		Value: newRefreshToken,
	}
	newLoginUserCookie := &http.Cookie{
		Name:     s.LoginUserCookie.Name,
		Secure:   s.LoginUserCookie.Secure,
		HttpOnly: s.LoginUserCookie.HttpOnly,
		Path:     s.LoginUserCookie.Path,
		MaxAge:   s.LoginUserCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(s.LoginUserCookie.MaxAge) * time.Second),
		SameSite: s.LoginUserCookie.SameSite,

		Value: string(u.ID),
	}
	newLoginStateCookie := &http.Cookie{
		Name:     s.LoginStateCookie.Name,
		Secure:   s.LoginStateCookie.Secure,
		HttpOnly: s.LoginStateCookie.HttpOnly,
		Path:     s.LoginStateCookie.Path,
		MaxAge:   s.LoginStateCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(s.LoginStateCookie.MaxAge) * time.Second),
		SameSite: s.LoginStateCookie.SameSite,

		Value: "1", // 1 or null
	}

	return u, []*http.Cookie{
		newSessionCookie,
		newRefreshCookie,
		newLoginUserCookie,
		newLoginStateCookie,
	}, nil
}

// ログインできなかったあとのcookieをいい感じに削除する
func (s *Session) loginFailed(ctx context.Context, cookies []*http.Cookie, refreshCookieName string) (*models.User, []*http.Cookie, error) {
	cookies, err := s.logoutC(ctx, cookies, refreshCookieName)
	if err != nil {
		return nil, []*http.Cookie{}, err
	}

	return nil, cookies, NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "login failed")
}

// ログアウトする
// Cookieを削除して、セッションをDBから削除する
func (s *Session) Logout(ctx context.Context, cookies []*http.Cookie, user *models.User) ([]*http.Cookie, error) {
	refreshCookieName := fmt.Sprintf("%s-%s", s.RefreshCookie.Name, user.ID)

	return s.logoutC(ctx, cookies, refreshCookieName)
}

func (s *Session) logoutC(ctx context.Context, cookies []*http.Cookie, refreshCookieName string) ([]*http.Cookie, error) {
	var sessionCookie *http.Cookie = nil
	var refreshCookie *http.Cookie = nil
	var loginUserCookie *http.Cookie = nil
	var loginStateCookie *http.Cookie = nil

	for _, cookie := range cookies {
		switch cookie.Name {
		case s.SessionCookie.Name:
			sessionCookie = cookie
		case refreshCookieName:
			refreshCookie = cookie
		case s.LoginUserCookie.Name:
			loginUserCookie = cookie
		case s.LoginStateCookie.Name:
			loginStateCookie = cookie
		}
	}

	deleteSetCookie := []*http.Cookie{}

	// セッションcookie、セッションを削除
	if sessionCookie != nil {
		if _, err := models.Sessions(
			models.SessionWhere.ID.EQ(sessionCookie.Value),
		).DeleteAll(ctx, s.DB); err != nil {
			return []*http.Cookie{}, err
		}

		deleteSetCookie = append(deleteSetCookie, &http.Cookie{
			Name:     s.SessionCookie.Name,
			Secure:   s.SessionCookie.Secure,
			HttpOnly: s.SessionCookie.HttpOnly,
			Path:     s.SessionCookie.Path,
			MaxAge:   -1,
			Expires:  time.Now(),
			SameSite: s.SessionCookie.SameSite,

			Value: "",
		})
	}
	// リフレッシュcookie、リフレッシュを削除
	if refreshCookie != nil && refreshCookieName != "" {
		if _, err := models.Refreshes(
			models.RefreshWhere.ID.EQ(refreshCookie.Value),
		).DeleteAll(ctx, s.DB); err != nil {
			return []*http.Cookie{}, err
		}

		deleteSetCookie = append(deleteSetCookie, &http.Cookie{
			Name:     refreshCookieName,
			Secure:   s.RefreshCookie.Secure,
			HttpOnly: s.RefreshCookie.HttpOnly,
			Path:     s.RefreshCookie.Path,
			MaxAge:   -1,
			Expires:  time.Now(),
			SameSite: s.RefreshCookie.SameSite,

			Value: "",
		})
	}
	if loginUserCookie != nil {
		deleteSetCookie = append(deleteSetCookie, &http.Cookie{
			Name:     s.LoginUserCookie.Name,
			Secure:   s.LoginUserCookie.Secure,
			HttpOnly: s.LoginUserCookie.HttpOnly,
			Path:     s.LoginUserCookie.Path,
			MaxAge:   -1,
			Expires:  time.Now(),
			SameSite: s.LoginUserCookie.SameSite,

			Value: "",
		})
	}
	if loginStateCookie != nil {
		deleteSetCookie = append(deleteSetCookie, &http.Cookie{
			Name:     s.LoginStateCookie.Name,
			Secure:   s.LoginStateCookie.Secure,
			HttpOnly: s.LoginStateCookie.HttpOnly,
			Path:     s.LoginStateCookie.Path,
			MaxAge:   -1,
			Expires:  time.Now(),
			SameSite: s.LoginStateCookie.SameSite,

			Value: "0", // 1 or null
		})
	}

	return deleteSetCookie, nil
}

// ユーザのセッションを登録する
func (s *Session) NewRegisterSession(ctx context.Context, user *models.User, ua *UserData, ip string) (*RegisterSession, error) {
	sessionToken, err := lib.RandomStr(31)
	if err != nil {
		return nil, err
	}
	refreshToken, err := lib.RandomStr(63)
	if err != nil {
		return nil, err
	}
	id := ulid.Make()
	idBin, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}

	// DBに入れる
	ss := models.Session{
		ID:     sessionToken,
		UserID: user.ID,

		Period: time.Now().Add(s.SessionDBPeriod),
	}
	if err := ss.Insert(ctx, s.DB, boil.Infer()); err != nil {
		return nil, err
	}
	rr := models.Refresh{
		ID:        refreshToken,
		UserID:    user.ID,
		HistoryID: idBin,
		SessionID: null.NewString(sessionToken, true),

		Period: time.Now().Add(s.RefreshDBPeriod),
	}
	if err := rr.Insert(ctx, s.DB, boil.Infer()); err != nil {
		return nil, err
	}

	// ログイン履歴を追加
	history := models.LoginHistory{
		UserID: user.ID,

		RefreshID: idBin,

		Device:   null.NewString(ua.Device, true),
		Os:       null.NewString(ua.OS, true),
		Browser:  null.NewString(ua.Browser, true),
		IsMobile: null.NewBool(ua.IsMobile, true),

		IP: net.ParseIP(ip),
	}
	if err := history.Insert(ctx, s.DB, boil.Infer()); err != nil {
		return nil, err
	}

	return &RegisterSession{
		SessionToken: sessionToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	}, nil
}

// ログイン用のCookie作成
func (s *RegisterSession) InsertCookie(c *Config) []*http.Cookie {
	refreshCookieName := fmt.Sprintf("%s-%s", c.RefreshCookie.Name, s.UserID)

	// セッショントークン
	sessionCookie := &http.Cookie{
		Name:     c.SessionCookie.Name,
		Secure:   c.SessionCookie.Secure,
		HttpOnly: c.SessionCookie.HttpOnly,
		Path:     c.SessionCookie.Path,
		MaxAge:   c.SessionCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(c.SessionCookie.MaxAge) * time.Second),
		SameSite: c.SessionCookie.SameSite,

		Value: s.SessionToken,
	}
	// リフレッシュトークン
	refreshCookie := &http.Cookie{
		Name:     refreshCookieName,
		Secure:   c.RefreshCookie.Secure,
		HttpOnly: c.RefreshCookie.HttpOnly,
		Path:     c.RefreshCookie.Path,
		MaxAge:   c.RefreshCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(c.RefreshCookie.MaxAge) * time.Second),
		SameSite: c.RefreshCookie.SameSite,

		Value: s.RefreshToken,
	}
	// ログインしているユーザ
	loginUserCookie := &http.Cookie{
		Name:     c.LoginUserCookie.Name,
		Secure:   c.LoginUserCookie.Secure,
		HttpOnly: c.LoginUserCookie.HttpOnly,
		Path:     c.LoginUserCookie.Path,
		MaxAge:   c.LoginUserCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(c.LoginUserCookie.MaxAge) * time.Second),
		SameSite: c.LoginUserCookie.SameSite,

		Value: s.UserID,
	}
	// ログイン状態（JSで見るよう）
	loginStateCookie := &http.Cookie{
		Name:     c.LoginStateCookie.Name,
		Secure:   c.LoginStateCookie.Secure,
		HttpOnly: c.LoginStateCookie.HttpOnly,
		Path:     c.LoginStateCookie.Path,
		MaxAge:   c.LoginStateCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(c.LoginStateCookie.MaxAge) * time.Second),
		SameSite: c.LoginStateCookie.SameSite,

		Value: "1", // 1 or null
	}

	return []*http.Cookie{
		sessionCookie,
		refreshCookie,
		loginUserCookie,
		loginStateCookie,
	}
}

// ログインするアカウントを変更する
// すでにセッションがログイン存在している場合は、そのセッションは削除する
// エラーでもcookieはセットする必要があります
// このメソッドではLoginUser cookieを更新するだけです
func (s *Session) SwitchAccount(ctx context.Context, cookies []*http.Cookie, userID string) ([]*http.Cookie, error) {
	// ユーザIDのユーザが存在するかチェック
	userExists, err := models.Users(
		models.UserWhere.ID.EQ(userID),
	).Exists(ctx, s.DB)
	if err != nil {
		return []*http.Cookie{}, err
	}
	if !userExists {
		return []*http.Cookie{}, NewHTTPError(http.StatusBadRequest, "user not found")
	}

	newUserRefreshTokenName := fmt.Sprintf("%s-%s", s.RefreshCookie.Name, userID)
	newCookie := []*http.Cookie{}
	var newRefreshCookie *http.Cookie = nil
	var sessionCookie *http.Cookie = nil
	var loggedInUser *http.Cookie = nil
	for _, cookie := range cookies {
		switch cookie.Name {
		case newUserRefreshTokenName:
			newRefreshCookie = cookie
		case s.SessionCookie.Name:
			sessionCookie = cookie
		case s.LoginUserCookie.Name:
			loggedInUser = cookie
		}
	}
	// 新規にログインするリフレッシュCookieが存在しないのでそもそもログインできない
	if newRefreshCookie == nil {
		return []*http.Cookie{}, NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "login failed")
	}
	// セッショントークンが無い場合はログインできない
	if sessionCookie == nil {
		return []*http.Cookie{}, NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "login failed")
	}
	// 変更先のユーザが現在ログインしているユーザと同じ場合は特に何もしない
	if loggedInUser.Value == userID {
		return []*http.Cookie{}, nil
	}

	// リフレッシュトークンの値が不正な場合はログインしない
	refresh, err := models.Refreshes(
		models.RefreshWhere.ID.EQ(newRefreshCookie.Value),
	).One(ctx, s.DB)
	if errors.Is(err, sql.ErrNoRows) {
		refreshCookie := &http.Cookie{
			Name:     newUserRefreshTokenName,
			Secure:   s.RefreshCookie.Secure,
			HttpOnly: s.RefreshCookie.HttpOnly,
			Path:     s.RefreshCookie.Path,
			MaxAge:   -1,
			Expires:  time.Now(),
			SameSite: s.RefreshCookie.SameSite,

			Value: "",
		}
		return []*http.Cookie{refreshCookie}, NewHTTPError(http.StatusBadRequest, "refresh token is invalid")
	}
	if err != nil {
		return []*http.Cookie{}, err
	}
	// リフレッシュトークンのユーザが違う場合はエラー
	if string(refresh.UserID) != userID {
		refreshCookie := &http.Cookie{
			Name:     newUserRefreshTokenName,
			Secure:   s.RefreshCookie.Secure,
			HttpOnly: s.RefreshCookie.HttpOnly,
			Path:     s.RefreshCookie.Path,
			MaxAge:   -1,
			Expires:  time.Now(),
			SameSite: s.RefreshCookie.SameSite,

			Value: "",
		}
		return []*http.Cookie{refreshCookie}, err
	}

	newCookie = append(newCookie, &http.Cookie{
		Name:     s.LoginUserCookie.Name,
		Secure:   s.LoginUserCookie.Secure,
		HttpOnly: s.LoginUserCookie.HttpOnly,
		Path:     s.LoginUserCookie.Path,
		MaxAge:   s.LoginUserCookie.MaxAge,
		Expires:  time.Now().Add(time.Duration(s.LoginUserCookie.MaxAge) * time.Second),
		SameSite: s.LoginUserCookie.SameSite,

		Value: userID,
	})

	if _, err := models.Sessions(
		models.SessionWhere.ID.EQ(sessionCookie.Value),
	).DeleteAll(ctx, s.DB); err != nil {
		return []*http.Cookie{}, err
	}

	newCookie = append(newCookie, &http.Cookie{
		Name:     s.SessionCookie.Name,
		Secure:   s.SessionCookie.Secure,
		HttpOnly: s.SessionCookie.HttpOnly,
		Path:     s.SessionCookie.Path,
		MaxAge:   -1,
		Expires:  time.Now(),
		SameSite: s.SessionCookie.SameSite,

		Value: "",
	})

	return newCookie, nil
}

// ログイン可能なアカウントを返す
// cookieのリフレッシュトークンからユーザを出しています
// 有効期限が切れたリフレッシュトークンは省きます
func (s *Session) LoggedInAccounts(ctx context.Context, cookies []*http.Cookie) ([]*models.User, error) {
	var refreshTokens []interface{}
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, s.RefreshCookie.Name) && cookie.Value != "" {
			refreshTokens = append(refreshTokens, cookie.Value)
		}
	}
	if len(refreshTokens) == 0 {
		return []*models.User{}, nil
	}

	// SELECT user.* FROM user
	// INNER JOIN refresh
	//     ON user.id = refresh.user_id
	// WHERE refresh.id IN ?
	// AND refresh.period > NOW()
	// ORDER BY user.id DESC;
	users, err := models.Users(
		qm.Select("user.*"),
		qm.InnerJoin("refresh ON user.id = refresh.user_id"),
		qm.WhereIn("refresh.id IN ?", refreshTokens...),
		qm.And("refresh.period > NOW()"),
		qm.OrderBy("user.id DESC"),
	).All(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// スタッフのみ
func (s *Session) RequireStuff(ctx context.Context, user *models.User) error {
	isStuff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(user.ID),
	).Exists(ctx, s.DB)
	if err != nil {
		return err
	}
	if !isStuff {
		return NewHTTPError(http.StatusForbidden, "require staff")
	}
	return nil
}
