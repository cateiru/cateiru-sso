package src

import (
	"errors"
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

	u, err := h.Session.SimpleLogin(ctx, c, true)
	if errors.Is(err, ErrorLoginFailed) {
		// 未ログインの場合は200でトークンを返す
		response, err := authenticationRequest.GetLoginSession(ctx, h.C.OauthLoginSessionPeriod, h.DB)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	}
	if err != nil {
		return err
	}

	ok, err := authenticationRequest.CheckUserAuthenticationPossible(ctx, h.DB, u)
	if err != nil {
		return err
	}
	if !ok {
		return NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "user is not allowed", "", "")
	}

	previewResponse, err := authenticationRequest.GetPreviewResponse(ctx, h.C.OauthLoginSessionPeriod, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, previewResponse)
}

func (h *Handler) OIDCLoginHandler(c echo.Context) error {
	ctx := c.Request().Context()

	authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
	if err != nil {
		return err
	}

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	ok, err := authenticationRequest.CheckUserAuthenticationPossible(ctx, h.DB, u)
	if err != nil {
		return err
	}
	if !ok {
		return NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "user is not allowed", "", "")
	}

	response, err := authenticationRequest.Submit(ctx, h.DB, u, h.C.OauthLoginSessionPeriod)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// キャンセルはユーザーの認証や、トークンチェックをしない
func (h *Handler) OIDCCancelHandler(c echo.Context) error {
	ctx := c.Request().Context()

	authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
	if err != nil {
		return err
	}

	response, err := authenticationRequest.Cancel(ctx, h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
