package src

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// リクエストされたOIDCクライアントが有効か、認証可能かなどの情報を返す
// jsonでpayloadを取得する
func (h *Handler) OIDCRequireHandler(c echo.Context) error {
	ctx := c.Request().Context()

	authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
	if err != nil {
		return err
	}

	previewResponse, err := authenticationRequest.GetPreviewResponse(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, previewResponse)
}

// 認証用のWebAuthnチャレンジを返す
func (h *Handler) OIDCBeginWebAuthnHandler(c echo.Context) error {
	return nil
}

// OIDCリクエスト時にWebAuthnで認証する
func (h *Handler) OIDCWebAuthnHandler(c echo.Context) error {
	return nil
}

// OIDCのリクエスト時にパスワードで認証する
func (h *Handler) OIDCPasswordHandler(c echo.Context) error {
	return nil
}

// ODICのリクエスト時にOTPで認証する
func (h *Handler) OIDCOTPHandler(c echo.Context) error {
	return nil
}

func (h *Handler) OIDCLoginHandler(c echo.Context) error {
	return nil
}
