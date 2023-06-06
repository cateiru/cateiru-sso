package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
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

var PromptEMUNS = []string{"login", "2fa_login"}
var ScopeEMUNS = []string{"email", ""}

type ClientResponse struct {
	ClientID string `json:"client_id"`

	Name        string      `json:"name"`
	Description null.String `json:"description,omitempty"`
	Image       null.String `json:"image,omitempty"`

	IsAllow bool        `json:"is_allow"`
	Prompt  null.String `json:"prompt,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientDetailResponse struct {
	ClientSecret string `json:"client_secret"`

	RedirectUrls []string `json:"redirect_urls"`
	ReferrerUrls []string `json:"referrer_urls"`

	Scopes []string `json:"scopes"`

	ClientResponse
}

type ClientAllowUserRuleResponse struct {
	Id uint `json:"id"`

	UserId      null.String `json:"user_id,omitempty"`
	EmailDomain null.String `json:"email_domain,omitempty"`
}

// クライアントの詳細を取得する
func getClientDetails(ctx context.Context, db *sql.DB, clientId string, u *models.User) (*ClientDetailResponse, error) {
	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
		models.ClientWhere.OwnerUserID.EQ(u.ID),
	).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, echo.NewHTTPError(http.StatusNotFound, "client not found")
	}
	if err != nil {
		return nil, err
	}

	redirectUrlRecords, err := models.ClientRedirects(
		models.ClientRedirectWhere.ClientID.EQ(client.ClientID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	redirectUrls := make([]string, len(redirectUrlRecords))
	for i, redirect := range redirectUrlRecords {
		redirectUrls[i] = redirect.URL
	}

	referrerUrlRecords, err := models.ClientReferrers(
		models.ClientReferrerWhere.ClientID.EQ(client.ClientID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	referrerUrls := make([]string, len(referrerUrlRecords))
	for i, referrer := range referrerUrlRecords {
		// referrerはホストのみを見るので
		referrerUrls[i] = referrer.Host
	}

	scopesRecords, err := models.ClientScopes(
		models.ClientScopeWhere.ClientID.EQ(client.ClientID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	scopes := make([]string, len(scopesRecords))
	for i, scope := range scopesRecords {
		scopes[i] = scope.Scope
	}

	return &ClientDetailResponse{
		ClientSecret: client.ClientSecret,

		RedirectUrls: redirectUrls,
		ReferrerUrls: referrerUrls,
		Scopes:       scopes,

		ClientResponse: ClientResponse{
			ClientID: client.ClientID,

			Name:        client.Name,
			Description: client.Description,
			Image:       client.Image,

			IsAllow: client.IsAllow,
			Prompt:  client.Prompt,

			CreatedAt: client.CreatedAt,
			UpdatedAt: client.UpdatedAt,
		},
	}, nil
}

// クライアントの一覧を返す
// client_idを指定するとそのクライアントを返す
func (h *Handler) ClientHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.QueryParam("client_id")

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// client_idが指定されている場合はそのIDのクライアントを返す
	if clientId != "" {
		response, err := getClientDetails(ctx, h.DB, clientId, u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	}

	clients, err := models.Clients(
		models.ClientWhere.OwnerUserID.EQ(u.ID),
		qm.Limit(h.C.ClientMaxCreated),
		qm.OrderBy("updated_at DESC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	response := make([]*ClientResponse, len(clients))
	for i, client := range clients {
		response[i] = &ClientResponse{
			ClientID: client.ClientID,

			Name:        client.Name,
			Description: client.Description,
			Image:       client.Image,

			IsAllow: client.IsAllow,
			Prompt:  client.Prompt,

			CreatedAt: client.CreatedAt,
			UpdatedAt: client.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, response)
}

// クライアントを作成する
// フォーム要素:
// - name: クライアント名
// - description?: クライアントの説明
// - image?: クライアントのアイコン
// - is_allow: ホワイトリスト使うか
// - prompt?: ログイン求めたりするやつ
// - scopes: スコープ
// - redirect_url: リダイレクトURL
//   - redirect_url_count: リダイレクトURLの数
//   - redirect_url_[index]: リダイレクトURL
//
// - referrer_url?: リファラURL
//   - referrer_url_count: リファラURLの数
//   - referrer_url_[index]: リファラURL
func (h *Handler) ClientCreateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	name := c.FormValue("name")
	description := c.FormValue("description")
	isAllowForm := c.FormValue("is_allow")
	prompt := c.FormValue("prompt")
	scope := c.FormValue("scopes")

	redirectUrlForms, err := h.FormValues(c, "redirect_url")
	if err != nil {
		return err
	}
	referrerUrlForms, err := h.FormValues(c, "referrer_url", true)
	if err != nil {
		return err
	}

	imageHeader, err := c.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	isAllow := isAllowForm == "true"

	if prompt != "" {
		// promptの値が正しいかどうかチェックする
		promptOk := false
		for _, p := range PromptEMUNS {
			if p == prompt {
				promptOk = true
				break
			}
		}
		if !promptOk {
			return NewHTTPError(http.StatusBadRequest, "prompt is invalid")
		}
	}

	if scope == "" {
		return NewHTTPError(http.StatusBadRequest, "scope is required")
	}
	scopes := strings.Split(scope, " ")
	if len(scopes) == 0 || scopes[0] == "" {
		return NewHTTPError(http.StatusBadRequest, "scope is invalid")
	}
	for _, s := range scopes {
		if !lib.ValidateScope(s) {
			return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("scope `%s` is invalid", s))
		}
	}

	if len(redirectUrlForms) > h.C.ClientRedirectURLMaxCreated {
		return NewHTTPError(http.StatusBadRequest, "too many redirect urls")
	}
	if len(referrerUrlForms) > h.C.ClientReferrerURLMaxCreated {
		return NewHTTPError(http.StatusBadRequest, "too many referrer urls")
	}

	redirectUrls := make([]url.URL, len(redirectUrlForms))
	for i, redirectUrlForm := range redirectUrlForms {
		redirectUrl, err := url.ParseRequestURI(redirectUrlForm)
		if err != nil {
			return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("referrer_url `%s` is invalid", redirectUrlForm))
		}
		redirectUrls[i] = *redirectUrl
	}
	referrerUrls := make([]url.URL, len(referrerUrlForms))
	for i, referrerUrlForm := range referrerUrlForms {
		referrerUrl, err := url.ParseRequestURI(referrerUrlForm)
		if err != nil {
			return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("referrer_url `%s` is invalid", referrerUrlForm))
		}
		referrerUrls[i] = *referrerUrl
	}

	// -- チェック終わり --

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	clientCount, err := models.Clients(
		models.ClientWhere.OwnerUserID.EQ(u.ID),
	).Count(ctx, h.DB)
	if err != nil {
		return err
	}
	// 新規作成するので現在あるクライアント数が上限-1以上であればエラー
	if (clientCount) >= int64(h.C.ClientMaxCreated-1) {
		return NewHTTPError(http.StatusBadRequest, "too many clients")
	}

	clientId := ulid.Make()

	clientSecret, err := lib.RandomStr(63)
	if err != nil {
		return err
	}

	image := null.NewString("", false)

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
		path := filepath.Join("client_icon", clientId.String())
		if err := h.Storage.Write(ctx, path, file, contentType); err != nil {
			return err
		}

		// ローカル環境では /[bucket-name]/avatar/[image] となるので
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

		image = null.NewString(url.String(), true)
	}

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		client := models.Client{
			ClientID: clientId.String(),

			Name:        name,
			Description: null.NewString(description, description != ""),
			Image:       image,
			IsAllow:     isAllow,
			Prompt:      null.NewString(prompt, prompt != ""),

			OwnerUserID: u.ID,

			ClientSecret: clientSecret,
		}
		if err := client.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		// FIXME: Bulk InsertしたいけどSQLBoilerにはないので
		for _, scope := range scopes {
			clientScope := models.ClientScope{
				ClientID: clientId.String(),
				Scope:    scope,
			}
			if err := clientScope.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		for _, redirectUrl := range redirectUrls {
			clientRedirectUrl := models.ClientRedirect{
				ClientID: clientId.String(),
				URL:      redirectUrl.String(),
				Host:     redirectUrl.Host,
			}
			if err := clientRedirectUrl.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		// リファラーURLはOptionalなのである場合のみ
		for _, referrerUrl := range referrerUrls {
			clientReferrerUrl := models.ClientReferrer{
				ClientID: clientId.String(),
				URL:      referrerUrl.String(),
				Host:     referrerUrl.Host,
			}
			if err := clientReferrerUrl.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	response, err := getClientDetails(ctx, h.DB, clientId.String(), u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// Clientを更新する
// フォーム要素:
// - client_id: クライアントID
// - name: クライアント名
// - description: クライアントの説明
// - image?: クライアントのアイコン。指定しない場合は、更新しない
//   - 画像を削除するには`ClientDeleteImageHandler`を使う
//
// - is_allow: ホワイトリスト使うか
// - prompt: ログイン求めたりするやつ
// - scopes: スコープ
// - update_secret: クライアントシークレットを更新するか
// - redirect_url: リダイレクトURL
//   - redirect_url_count: リダイレクトURLの数
//   - redirect_url_[index]: リダイレクトURL
//
// - referrer_url?: リファラURL
//   - referrer_url_count: リファラURLの数
//   - referrer_url_[index]: リファラURL
func (h *Handler) ClientUpdateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.FormValue("client_id")
	name := c.FormValue("name")
	description := c.FormValue("description")
	isAllowForm := c.FormValue("is_allow")
	prompt := c.FormValue("prompt")
	scope := c.FormValue("scopes")

	redirectUrlForms, err := h.FormValues(c, "redirect_url")
	if err != nil {
		return err
	}
	referrerUrlForms, err := h.FormValues(c, "referrer_url", true)
	if err != nil {
		return err
	}

	updateSecretForm := c.FormValue("update_secret")

	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	imageHeader, err := c.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	isAllow := isAllowForm == "true"
	updateSecret := updateSecretForm == "true"

	if prompt != "" {
		// promptの値が正しいかどうかチェックする
		promptOk := false
		for _, p := range PromptEMUNS {
			if p == prompt {
				promptOk = true
				break
			}
		}
		if !promptOk {
			return NewHTTPError(http.StatusBadRequest, "prompt is invalid")
		}
	}

	if scope == "" {
		return NewHTTPError(http.StatusBadRequest, "scope is required")
	}
	scopes := strings.Split(scope, " ")
	if len(scopes) == 0 && scopes[0] != "" {
		return NewHTTPError(http.StatusBadRequest, "scope is invalid")
	}
	for _, s := range scopes {
		if !lib.ValidateScope(s) {
			return NewHTTPError(http.StatusBadRequest, "scope is invalid")
		}
	}

	if len(redirectUrlForms) > h.C.ClientRedirectURLMaxCreated {
		return NewHTTPError(http.StatusBadRequest, "too many redirect urls")
	}
	if len(referrerUrlForms) > h.C.ClientReferrerURLMaxCreated {
		return NewHTTPError(http.StatusBadRequest, "too many referrer urls")
	}

	redirectUrls := make([]url.URL, len(redirectUrlForms))
	for i, redirectUrlForm := range redirectUrlForms {
		redirectUrl, err := url.ParseRequestURI(redirectUrlForm)
		if err != nil {
			return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("referrer_url `%s` is invalid", redirectUrlForm))
		}
		redirectUrls[i] = *redirectUrl
	}
	referrerUrls := make([]url.URL, len(referrerUrlForms))
	for i, referrerUrlForm := range referrerUrlForms {
		referrerUrl, err := url.ParseRequestURI(referrerUrlForm)
		if err != nil {
			return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("referrer_url `%s` is invalid", referrerUrlForm))
		}
		referrerUrls[i] = *referrerUrl
	}

	// -- チェック終わり --

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// 画像をアップロードする前に一度、clientIdのクライアントが存在するかを確認する
	existClient, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
		models.ClientWhere.OwnerUserID.EQ(u.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existClient {
		return NewHTTPError(http.StatusNotFound, "client not found")
	}

	// 画像をアップロードする（ある場合）
	image := null.NewString("", false)
	if imageHeader != nil {
		file, err := imageHeader.Open()
		if err != nil {
			return err
		}
		contentType := imageHeader.Header.Get("Content-Type")
		if !lib.ValidateContentType(contentType) {
			return NewHTTPError(http.StatusBadRequest, "invalid Content-Type")
		}
		path := filepath.Join("client_icon", clientId)
		if err := h.Storage.Write(ctx, path, file, contentType); err != nil {
			return err
		}
		// ローカル環境では /[bucket-name]/avatar/[image] となるので
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

		image = null.NewString(url.String(), true)
	}

	// トランザクション
	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
			models.ClientWhere.OwnerUserID.EQ(u.ID),
		).One(ctx, tx)
		if errors.Is(err, sql.ErrNoRows) {
			return NewHTTPError(http.StatusNotFound, "client not found")
		}
		if err != nil {
			return err
		}

		if updateSecret {
			clientSecret, err := lib.RandomStr(63)
			if err != nil {
				return err
			}
			client.ClientSecret = clientSecret
		}

		client.Image = image
		client.Name = name
		client.Description = null.NewString(description, description != "")
		client.IsAllow = isAllow
		client.Prompt = null.NewString(prompt, prompt != "")

		if _, err := client.Update(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}

		// スコープを一度すべて削除してから追加する
		if _, err := client.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(clientId),
		).DeleteAll(ctx, h.DB); err != nil {
			return err
		}
		// FIXME: Bulk InsertしたいけどSQLBoilerにはないので
		for _, scope := range scopes {
			clientScope := models.ClientScope{
				ClientID: clientId,
				Scope:    scope,
			}
			if err := clientScope.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		// リダイレクトURLを一度すべて削除してから追加する
		if _, err := client.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(clientId),
		).DeleteAll(ctx, h.DB); err != nil {
			return err
		}
		for _, redirectUrl := range redirectUrls {
			clientRedirectUrl := models.ClientRedirect{
				ClientID: clientId,
				URL:      redirectUrl.String(),
				Host:     redirectUrl.Host,
			}
			if err := clientRedirectUrl.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		// リダイレクトURLを一度すべて削除してから追加する
		if _, err := client.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(clientId),
		).DeleteAll(ctx, h.DB); err != nil {
			return err
		}
		// リファラーURLはOptionalなのである場合のみ
		for _, referrerUrl := range referrerUrls {
			clientReferrerUrl := models.ClientReferrer{
				ClientID: clientId,
				URL:      referrerUrl.String(),
				Host:     referrerUrl.Host,
			}
			if err := clientReferrerUrl.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	response, err := getClientDetails(ctx, h.DB, clientId, u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// クライアント削除
func (h *Handler) ClientDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.QueryParam("client_id")
	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	setImage := false

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
			models.ClientWhere.OwnerUserID.EQ(u.ID),
		).One(ctx, tx)
		if errors.Is(err, sql.ErrNoRows) {
			return NewHTTPError(http.StatusNotFound, "client not found")
		}
		if err != nil {
			return err
		}

		if client.Image.Valid {
			setImage = true
		}

		if _, err := client.Delete(ctx, tx); err != nil {
			return err
		}

		// スコープすべて削除
		if _, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(client.ClientID),
		).DeleteAll(ctx, tx); err != nil {
			return err
		}

		// ホワイトリストルールを削除する
		if _, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(client.ClientID),
		).DeleteAll(ctx, tx); err != nil {
			return err
		}

		// リダイレクトURLを削除する
		if _, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(client.ClientID),
		).DeleteAll(ctx, tx); err != nil {
			return err
		}

		// リファラーURLを削除する
		if _, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(client.ClientID),
		).DeleteAll(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	// 画像が設定されている場合は削除する
	if setImage {
		path := filepath.Join("client_icon", clientId)
		if err := h.Storage.Delete(ctx, path); err != nil {
			return err
		}

		// ローカル環境では /[bucket-name]/avatar/[image] となるので
		p, err := url.JoinPath(h.C.CDNHost.Path, path)
		if err != nil {
			return err
		}

		// CDNをパージ
		url := &url.URL{
			Scheme: h.C.CDNHost.Scheme,
			Host:   h.C.CDNHost.Host,
			Path:   p,
		}
		if err := h.CDN.Purge(url.String()); err != nil {
			return err
		}
	}

	return nil
}

// クライアント画像の削除
func (h *Handler) ClientDeleteImageHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.QueryParam("client_id")
	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
		models.ClientWhere.OwnerUserID.EQ(u.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "client not found")
	}
	if err != nil {
		return err
	}

	if !client.Image.Valid {
		return NewHTTPError(http.StatusNotFound, "image is not set")
	}

	path := filepath.Join("client_icon", client.ClientID)
	if err := h.Storage.Delete(ctx, path); err != nil {
		return err
	}

	// ローカル環境では /[bucket-name]/avatar/[image] となるので
	p, err := url.JoinPath(h.C.CDNHost.Path, path)
	if err != nil {
		return err
	}

	// CDNをパージ
	url := &url.URL{
		Scheme: h.C.CDNHost.Scheme,
		Host:   h.C.CDNHost.Host,
		Path:   p,
	}
	if err := h.CDN.Purge(url.String()); err != nil {
		return err
	}

	client.Image = null.NewString("", false)

	if _, err := client.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// AllowUserを返す
func (h *Handler) ClientAllowUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.QueryParam("client_id")
	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
		models.ClientWhere.OwnerUserID.EQ(u.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !client {
		return NewHTTPError(http.StatusNotFound, "client not found")
	}

	rules, err := models.ClientAllowRules(
		models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		qm.Limit(100),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	roleResponse := make([]ClientAllowUserRuleResponse, len(rules))
	for i, rule := range rules {
		roleResponse[i] = ClientAllowUserRuleResponse{
			Id: rule.ID,

			UserId:      rule.UserID,
			EmailDomain: rule.EmailDomain,
		}
	}

	return c.JSON(http.StatusOK, roleResponse)
}

// ホワイトリストにユーザーを追加する
func (h *Handler) ClientAddAllowUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.FormValue("client_id")
	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	userId := c.FormValue("user_id")
	emailDomain := c.FormValue("email_domain")
	// どちらか必須
	if userId != "" || emailDomain != "" {
		return NewHTTPError(http.StatusBadRequest, "user_id or email_domain is required")
	}
	// 片方しか設定できない
	if userId == emailDomain {
		return NewHTTPError(http.StatusBadRequest, "user_id and email_domain cannot be set at the same time")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// クライアントのis_allowがfalseでもホワイトリストに追加削除はできる
	existClient, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
		models.ClientWhere.OwnerUserID.EQ(u.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existClient {
		return NewHTTPError(http.StatusNotFound, "client not found")
	}

	clientAllowRule := &models.ClientAllowRule{
		ClientID: clientId,

		UserID:      null.NewString(userId, userId != ""),
		EmailDomain: null.NewString(emailDomain, emailDomain != ""),
	}
	if err := clientAllowRule.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// ホワイトリストからユーザーを削除する
func (h *Handler) ClientDeleteAllowUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if id == "" {
		return NewHTTPError(http.StatusBadRequest, "id is required")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "id is invalid")
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	allowRule, err := models.ClientAllowRules(
		models.ClientAllowRuleWhere.ID.EQ(uint(idInt)),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusNotFound, "allow rule not found")
	}
	if err != nil {
		return err
	}

	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(allowRule.ClientID),
	).One(ctx, h.DB)
	if err != nil {
		// レコードが存在しないことは無いのですべてエラーにする
		return err
	}

	if client.OwnerUserID != u.ID {
		return NewHTTPError(http.StatusForbidden, "you are not owner")
	}

	if _, err := allowRule.Delete(ctx, h.DB); err != nil {
		return err
	}

	return nil
}

// クライアントにログインしているユーザー一覧を返す
// TODO: クライアントのセッション実装したらやる
func (h *Handler) ClientLoginUsersHandler(c echo.Context) error {
	return nil
}
