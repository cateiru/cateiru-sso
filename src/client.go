package src

import (
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
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

type ClientAllowUserRuleResponse struct {
	Id uint `json:"id"`

	UserId      null.String `json:"user_id,omitempty"`
	EmailDomain null.String `json:"email_domain,omitempty"`
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
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
			models.ClientWhere.OwnerUserID.EQ(u.ID),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "client not found")
		}
		if err != nil {
			return err
		}

		response := &ClientResponse{
			ClientID: client.ClientID,

			Name:        client.Name,
			Description: client.Description,
			Image:       client.Image,

			IsAllow: client.IsAllow,
			Prompt:  client.Prompt,

			CreatedAt: client.CreatedAt,
			UpdatedAt: client.UpdatedAt,
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
// - description: クライアントの説明
// - image?: クライアントのアイコン
// - is_allow: ホワイトリスト使うか
// - prompt: ログイン求めたりするやつ
// - scopes: スコープ
func (h *Handler) ClientCreateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	name := c.FormValue("name")
	description := c.FormValue("description")
	isAllowForm := c.FormValue("is_allow")
	prompt := c.FormValue("prompt")
	scope := c.FormValue("scopes")

	image := null.NewString("", false)
	imageHeader, err := c.FormFile("image")
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	isAllow := isAllowForm == "true"

	if prompt == "" {
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
		if lib.ValidateScope(s) {
			return NewHTTPError(http.StatusBadRequest, "scope is invalid")
		}
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

	clientId, err := lib.RandomStr(32)
	if err != nil {
		return err
	}
	clientSecret, err := lib.RandomStr(63)
	if err != nil {
		return err
	}

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

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		client := models.Client{
			ClientID: clientId,

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
				ClientID: clientId,
				Scope:    scope,
			}
			if err := clientScope.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	currentClient, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	response := &ClientResponse{
		ClientID: currentClient.ClientID,

		Name:        currentClient.Name,
		Description: currentClient.Description,
		Image:       currentClient.Image,

		IsAllow: currentClient.IsAllow,
		Prompt:  currentClient.Prompt,

		CreatedAt: currentClient.CreatedAt,
		UpdatedAt: currentClient.UpdatedAt,
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
func (h *Handler) ClientUpdateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.FormValue("client_id")
	name := c.FormValue("name")
	description := c.FormValue("description")
	isAllowForm := c.FormValue("is_allow")
	prompt := c.FormValue("prompt")
	scope := c.FormValue("scopes")

	updateSecretForm := c.FormValue("update_secret")

	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	imageHeader, err := c.FormFile("image")
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	if name == "" {
		return NewHTTPError(http.StatusBadRequest, "name is required")
	}

	isAllow := isAllowForm == "true"
	updateSecret := updateSecretForm == "true"

	if prompt == "" {
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
		if lib.ValidateScope(s) {
			return NewHTTPError(http.StatusBadRequest, "scope is invalid")
		}
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

		return nil
	})
	if err != nil {
		return err
	}

	currentClient, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	response := &ClientResponse{
		ClientID: currentClient.ClientID,

		Name:        currentClient.Name,
		Description: currentClient.Description,
		Image:       currentClient.Image,

		IsAllow: currentClient.IsAllow,
		Prompt:  currentClient.Prompt,

		CreatedAt: currentClient.CreatedAt,
		UpdatedAt: currentClient.UpdatedAt,
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

	roleResponse := make([]*ClientAllowUserRuleResponse, 0, len(rules))
	for i, rule := range rules {
		roleResponse[i] = &ClientAllowUserRuleResponse{
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
