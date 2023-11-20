package src

import (
	"errors"
	"net/http"
	"path/filepath"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/labstack/echo/v4"
)

// 静的画像を返す
// このAPIはBasic認証がついているのでFastlyなどのCDNからアクセスしてキャッシュします
func (h *Handler) InternalAvatarHandler(c echo.Context) error {
	ctx := c.Request().Context()

	key := c.Param("key")
	id := c.Param("id")

	if key == "" || id == "" {
		return NewHTTPError(http.StatusBadRequest, "key and id are required")
	}

	path := filepath.Join(key, id)
	data, contentType, err := h.Storage.Read(ctx, path)
	if errors.Is(err, lib.ErrBucketNotFound) || errors.Is(err, lib.ErrObjectNotFound) {
		return NewHTTPError(http.StatusNotFound, "image not found")
	}
	if err != nil {
		return err
	}

	// キャッシュ設定
	// ブラウザキャッシュは行わずにCDNのみのキャッシュにする
	// こうすることで、画像が更新されてもリロードで変えることができる
	c.Response().Header().Set("Cache-Control", "s-maxage=31536000")

	return c.Blob(http.StatusOK, contentType, data)
}

// TODO
func (h *Handler) InternalWorkerHandler(c echo.Context) error {
	return nil
}
