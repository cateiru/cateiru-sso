package createaccount

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
	"golang.org/x/net/websocket"
)

// メールが認証されているかを1秒おきにチェックし、
// 認証された場合`true`を返します。
//
// もし、認証されずcloseされた場合は、メールに送信されたURLから続きを始めるため、OpenNewWindowをtrueにします。
func MailVerifyObserve(w http.ResponseWriter, r *http.Request) error {
	clientCheckToken, err := net.GetQuery(r, "cct")
	ctx := r.Context()
	// クエリパラメータがない場合は400を返す
	if err != nil {
		return status.NewBadRequestError(err).Caller("core/create_account/create_observe.go", 25).Wrap()
	}

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller(
			"core/create_account/temporary_account.go", 30).Wrap()
	}
	defer db.Close()

	// プロトコルをWSにアップデーーーーート！！！
	s := websocket.Server{
		Handler: websocket.Handler(
			func(ws *websocket.Conn) {
				isVerified := false

				defer closeWS(ctx, db, clientCheckToken, isVerified)

				quit := make(chan bool)

				go send(ctx, db, ws, quit, clientCheckToken, &isVerified)
				receive(ctx, db, ws, quit, clientCheckToken)
			}),
	}
	s.ServeHTTP(w, r)

	return nil
}

// websocket受信
//
// クライアントからの受信はしない
func receive(ctx context.Context, db *database.Database, ws *websocket.Conn, quit chan bool, token string) {
	for {
		var response []byte
		if err := websocket.Message.Receive(ws, &response); err != nil {
			if err == io.EOF {
				quit <- true
				logging.Sugar.Debugf("close websocket. token: %s", token)
			} else {
				logging.Sugar.Errorf("websocket err: %v", err.Error())
			}
			return
		}

	}
}

// WS送信
// 1秒おきにDBを参照し、認証されているかを確認する
// 認証された場合、trueを返す
func send(ctx context.Context, db *database.Database, ws *websocket.Conn, quit chan bool, token string, isVerified *bool) {
	for {
		select {
		case <-quit:
			return
		default:
			entry, err := models.GetMailCertificationByCheckToken(ctx, db, token)
			if err != nil {
				logging.Sugar.Error(err)
				ws.Close()
				return
			}
			// 要素がない状態はそのまま終了
			// 有効期限切れでデータが削除されたときなど
			if entry == nil {
				ws.Close()
				return
			}

			// 認証された場合、`true`を送信
			// wsのcloseは原則としてclient側から行い、こちら側で閉じるのはエラーのときのみにする
			if entry.Verify {
				*isVerified = true
				if err := websocket.Message.Send(ws, "true"); err != nil {
					logging.Sugar.Error(err)
					ws.Close()
				}

				return
			}

			// 1秒おきにチェックする
			time.Sleep(1 * time.Second)
		}
	}
}

// WSを閉じるときにする処理
func closeWS(ctx context.Context, db *database.Database, token string, isVerified bool) {
	if !isVerified {
		// 認証されずcloseされた場合は、メールに送信されたURLから続きを始めるため、OpenNewWindowをtrueにする
		entry, err := models.GetMailCertificationByCheckToken(ctx, db, token)
		if err != nil {
			logging.Sugar.Error(err)
			return
		}
		entry.OpenNewWindow = true

		if err := entry.Add(ctx, db); err != nil {
			logging.Sugar.Error(err)
			return
		}
	}
}
