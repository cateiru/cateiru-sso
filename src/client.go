package src

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var PromptEMUNS = []string{"none", "login", "2fa_login"}
var ScopeEMUNS = []string{"email", ""}

// クライアントの一覧を返す
// client_idを指定するとそのクライアントを返す
func (h *Handler) ClientHandler(c echo.Context) error {
	return nil
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

	// -- チェック終わり --

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
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
		url := &url.URL{
			Scheme: h.C.CDNHost.Scheme,
			Host:   h.C.CDNHost.Host,
			Path:   path,
		}
		image = null.NewString(url.String(), true)
	}

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
	if err := client.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (h *Handler) ClientUpdateHandler(c echo.Context) error {
	return nil
}

func (h *Handler) ClientDeleteHandler(c echo.Context) error {
	return nil
}

// ホワイトリストにユーザーを追加する
func (h *Handler) ClientAddAllowUser(c echo.Context) error {
	return nil
}

// ホワイトリストからユーザーを削除する
func (h *Handler) ClientDeleteAllowUser(c echo.Context) error {
	return nil
}

// クライアントにログインしているユーザー一覧を返す
func (h *Handler) ClientLoginUsersHandler(c echo.Context) error {
	return nil
}
