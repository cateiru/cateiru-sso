package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/exp/slices"
)

// https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthRequest
type AuthenticationRequest struct {
	Scopes []string

	// コードフローを決定する値
	// Authorization Code Flow の場合は `code`
	ResponseType lib.ResponseType

	// レスポンスが返される Redirection URI
	RedirectUri *url.URL

	// リクエストとコールバックの間で維持されるランダムな値
	// SPAなのでjs側で保持する予定。そのためこの値はサーバー側で使う予定はない
	State null.String

	// パラメータを返す手段
	ResponseMode lib.ResponseMode

	// Client セッションと ID Token を紐づける文字列であり、リプレイアタック対策に用いられる
	Nonce null.String

	// Authorization Server が認証および同意のためのユーザーインタフェースを End-User にどのように表示するかを指定するための ASCII 値
	// Authorization Server は User Agent の機能を検知して適切な表示を行うようにしても良い
	//
	// - page: Authorization Server は認証および同意 UI を User Agent の全画面に表示すべきである (SHOULD). display パラメータが指定されていない場合, この値がデフォルトとなる
	// - popup: Authorization Server は認証および同意 UI を User Agent のポップアップウィンドウに表示すべきである (SHOULD). User Agent のポップアップウィンドウはログインダイアログに適切なサイズで, 親ウィンドウ全体を覆うことのないようにすべきである
	// - touch: Authorization Server は認証および同意 UI をタッチインタフェースを持つデバイスに適した形で表示すべきである (SHOULD)
	// - wap: Authorization Server は認証および同意 UI を "feature phone" に適した形で表示すべきである (SHOULD)
	Display lib.Display

	// Authorization Server が End-User に再認証および同意を再度要求するかどうか指定するための, スペース区切りの ASCII 文字列のリスト. 以下の値が定義されている
	// - none: Authorization Server はいかなる認証および同意 UI をも表示してはならない
	// - login: Authorization Server は End-User を再認証するべきである
	// - consent: Authorization Server は Client にレスポンスを返す前に End-User に同意を要求するべきである
	// - select_account: Authorization Server は End-User にアカウント選択を促すべきである
	Prompts []lib.Prompt

	// Authentication Age の最大値. End-User が OP によって明示的に認証されてからの経過時間の最大許容値 (秒)
	MaxAge uint64

	// ロケール。一旦これはja_JPのみを想定するが、将来的には他の言語も対応すると想定してサーバにも持ってくる
	UiLocales   []string
	IdTokenHint null.String
	LoginHint   null.String
	AcrValues   null.String

	Client *models.Client

	AllowRules  []*models.ClientAllowRule
	RefererHost string
}

// プレビュー用のレスポンスを返す
func (a *AuthenticationRequest) GetPreviewResponse(ctx context.Context, loginSessionPeriod time.Duration, db *sql.DB, sessionToken string) (*PreviewResponse, error) {

	orgName := null.NewString("", false)
	orgImage := null.NewString("", false)

	if a.Client.OrgID.Valid {
		// orgは見つからないことはないはずなので、見つからなかったら500エラーにする
		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(a.Client.OrgID.String),
		).One(ctx, db)
		if err != nil {
			return nil, err
		}
		orgName = null.NewString(org.Name, true)
		orgImage = org.Image
	}

	// userは見つからないことはないはずなので、見つからなかったら500エラーにする
	user, err := models.Users(
		models.UserWhere.ID.EQ(a.Client.OwnerUserID),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	var loginSession *LoginSession = nil
	registerLoginSession := func() error {
		// max_age が設定されている場合はその秒数で有効期限を設定する
		period := loginSessionPeriod
		if a.MaxAge != 0 {
			period = time.Duration(a.MaxAge) * time.Second
		}

		loginSession, err = a.GetLoginSession(ctx, period, db)
		if err != nil {
			return err
		}
		return nil
	}

	// prompt = login の場合、ログインセッションを作成する
	// - トークンがすでにログイン済みだった場合はプレビューを返す
	// - トークンが有効期限切れなどで存在しない場合は再度トークンを作り直す
	if slices.Contains(a.Prompts, lib.PromptLogin) {
		if sessionToken != "" {
			// セッションがある場合はDBから引いてきて、有効かつログイン済みの場合はそのまま通す
			loginSession, err := models.OauthLoginSessions(
				models.OauthLoginSessionWhere.Token.EQ(sessionToken),
				models.OauthLoginSessionWhere.Period.GT(time.Now()),
			).One(ctx, db)
			if errors.Is(err, sql.ErrNoRows) {
				// セッション切れなどでトークンが有効ではなかった場合は再度ログインセッションを作る
				if err := registerLoginSession(); err != nil {
					return nil, err
				}
			}
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
			// QUESTION: ただのエラーで良いんだっけ？
			if loginSession != nil && !loginSession.LoginOk {
				return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "no login", "", "")
			}
		} else {
			if err := registerLoginSession(); err != nil {
				return nil, err
			}
		}
	}

	return &PreviewResponse{
		ClientId:          a.Client.ClientID,
		ClientName:        a.Client.Name,
		ClientDescription: a.Client.Description,
		Image:             a.Client.Image,

		OrgName:       orgName,
		OrgImage:      orgImage,
		OrgMemberOnly: a.Client.OrgMemberOnly,

		Scopes:       a.Scopes,
		RedirectUri:  a.RedirectUri.String(),
		ResponseType: string(a.ResponseType),

		RegisterUserName:  user.UserName,
		RegisterUserImage: user.Avatar,

		Prompts: a.Prompts,

		LoginSession: loginSession,
	}, nil
}

// ログインが必要な場合のセッションを返す
func (a *AuthenticationRequest) GetLoginSession(ctx context.Context, period time.Duration, db *sql.DB) (*LoginSession, error) {
	token, err := lib.RandomStr(31)
	if err != nil {
		return nil, err
	}

	limit := time.Now().Add(period)

	oauthLoginSession := models.OauthLoginSession{
		Token:        token,
		ClientID:     a.Client.ClientID,
		ReferrerHost: null.NewString(a.RefererHost, a.RefererHost != ""),
		Period:       limit,
	}
	if err := oauthLoginSession.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return &LoginSession{
		LoginSessionToken: token,
		LimitDate:         limit,
	}, nil
}

// ユーザーが認証可能かチェックする
func (a *AuthenticationRequest) CheckUserAuthenticationPossible(ctx context.Context, db *sql.DB, user *models.User) (bool, error) {
	ok := false

	// ルールが存在しない場合はすべてが認証可能
	if len(a.AllowRules) == 0 {
		ok = true
	}

	for _, rule := range a.AllowRules {
		// ユーザーが一致している場合
		if rule.UserID.Valid && rule.UserID.String == user.ID {
			ok = true
			break
		}

		// メールドメインが後方一致している場合
		if rule.EmailDomain.Valid && strings.HasSuffix(user.Email, rule.EmailDomain.String) {
			ok = true
			break
		}
	}

	// クライアントが組織所属のものかつメンバーオンリーの場合は
	// ユーザーをチェックする
	if a.Client.OrgID.Valid && a.Client.OrgMemberOnly {
		memberExist, err := models.OrganizationUsers(
			models.OrganizationUserWhere.OrganizationID.EQ(a.Client.OrgID.String),
			models.OrganizationUserWhere.UserID.EQ(user.ID),
		).Exists(ctx, db)
		if err != nil {
			return false, err
		}

		if memberExist {
			ok = true
		} else {
			// `OrgMemberOnly` が true の場合にユーザーがそのorgに所属していない場合は強制false
			ok = false
		}
	}

	return ok, nil
}

// TODO: test
func (a *AuthenticationRequest) Submit(ctx context.Context, db *sql.DB) (*OauthResponse, error) {
	return &OauthResponse{}, nil
}

// TODO: test
func (a *AuthenticationRequest) Cancel(ctx context.Context, db *sql.DB) (*OauthResponse, error) {
	return &OauthResponse{}, nil
}

func SetLoggedInOauthLoginSession(ctx context.Context, db *sql.DB, token string) error {
	oauthLoginSession, err := models.OauthLoginSessions(
		models.OauthLoginSessionWhere.Token.EQ(token),
		models.OauthLoginSessionWhere.Period.GT(time.Now()),
	).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		// トークンが不正だった場合は無視する
		return nil
	}
	if err != nil {
		return err
	}

	oauthLoginSession.LoginOk = true

	if _, err := oauthLoginSession.Update(ctx, db, boil.Infer()); err != nil {
		return err
	}

	return nil
}
