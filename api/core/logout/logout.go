package logout

import (
	"context"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	return Logout(ctx, db, w, r)
}

// ログアウトする
//
// ログインに必要なcookieを削除とセッションDBを削除します
func Logout(ctx context.Context, db *database.Database, w http.ResponseWriter, r *http.Request) error {
	sessionToken, err := net.GetCookie(r, "session-token")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}
	refreshToken, err := net.GetCookie(r, "refresh-token")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	// DBとcookieを削除
	// tokenがない場合は削除しない（有効期限で勝手に切れるのを待つしかない）
	//
	// refresh-tokenを削除する
	if len(refreshToken) != 0 {
		// もし、何らかの理由でrefresh-tokenはあるのにsession-tokenがcookie内になかった場合、
		// refresh-tokenのDBからsession-tokenを参照して削除する
		if len(sessionToken) == 0 {
			entity, err := models.GetRefreshToken(ctx, db, refreshToken)
			if err != nil {
				return status.NewInternalServerErrorError(err).Caller()
			}
			sessionToken = entity.SessionToken
		}

		if err := models.DeleteRefreshToken(ctx, db, refreshToken); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		if err := net.DeleteCookie(w, r, "refresh-token"); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}
	// session-tokenを削除する
	if len(sessionToken) != 0 {
		if err := models.DeleteSessionToken(ctx, db, sessionToken); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		if err := net.DeleteCookie(w, r, "session-token"); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	return nil
}
