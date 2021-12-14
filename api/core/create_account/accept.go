package createaccount

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

// client check tokenを使用して、buffer tokenをcookieに格納します
//
// このメソッドを叩く前にユーザに一度問いましょう（複数ウィンドウで同時にやるとどれか1つしか適用されません）
func CreateAcceptHandler(w http.ResponseWriter, r *http.Request) error {
	clientCheckToken, err := net.GetQuery(r, "token")
	// クエリがない場合、tokenが空の場合は400を返す
	if err != nil || len(clientCheckToken) == 0 {
		return status.NewBadRequestError(err).Caller("core/create_account/accept.go", 17).Wrap()
	}

	ctx := r.Context()
	bufferToken, err := AcceptVerify(ctx, clientCheckToken)
	if err != nil {
		return err
	}

	// secure属性はproductionのみにする（テストが通らないため）
	secure := false
	if utils.DEPLOY_MODE == "production" {
		secure = true
	}
	// ブラウザ上でcookieを追加できるように、HttpOnlyはfalseにする
	cookie := net.NewCookie(os.Getenv("COOKIE_DOMAIN"), secure, http.SameSiteDefaultMode, false)

	cookieExp := net.NewSession()
	cookie.Set(w, "buffer-token", bufferToken, cookieExp)

	return nil
}

func AcceptVerify(ctx context.Context, clientCheckToken string) (string, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/accept.go", 100).Wrap()
	}
	defer db.Close()

	certificationEntry, err := models.GetMailCertificationByCheckToken(ctx, db, clientCheckToken)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/accept.go", 100).Wrap()
	}

	// entryがない場合は400を返す
	if certificationEntry == nil {
		return "", status.NewBadRequestError(errors.New("deleted entry")).Caller("core/create_account/accept.go", 98).Wrap()
	}

	// 認証済みではない場合、403を返す
	if !certificationEntry.Verify {
		return "", status.NewForbiddenError(errors.New("not verify")).Caller(
			"core/create_account/accept.go", 103).Wrap()
	}

	// 有効期限が切れている場合は、400を返す
	if common.CheckExpired(&certificationEntry.Period) {
		return "", status.NewBadRequestError(errors.New("Expired")).Caller(
			"core/create_account/accept.go", 67).AddCode(net.TimeOutError).Wrap()
	}

	// bufferTokenを設定する
	bufferToken := utils.CreateID(20)
	buffer := &models.CreateAccountBuffer{
		BufferToken: bufferToken,
		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 60,
		},
		UserMailPW: certificationEntry.UserMailPW,
	}
	if err := buffer.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/accept.go", 100).Wrap()
	}

	// 元のメール認証用Entryは削除する
	// そのため、このAPIは一回のみのアクセスとなります
	if err := models.DeleteMailCertification(ctx, db, certificationEntry.MailToken); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/accept.go", 133).Wrap()
	}

	return bufferToken, nil
}
