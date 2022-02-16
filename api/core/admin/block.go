package admin

import (
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func SetBanHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	c := common.NewCert(w, r)
	if err := c.Login(ctx, db); err != nil {
		return err
	}
	userId := c.UserId

	// Adminのユーザのみ使用可
	if err := common.AdminOnly(ctx, db, userId); err != nil {
		return err
	}

	var request BanRequest

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	if request.IP != "" {
		ipBan := models.IPBlockList{
			IP: request.IP,
		}

		if err := ipBan.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}
	if request.Mail != "" {
		mailBan := models.MailBlockList{
			Mail: request.Mail,
		}

		if err := mailBan.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	return nil
}

func BanHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	c := common.NewCert(w, r)
	if err := c.Login(ctx, db); err != nil {
		return err
	}
	userId := c.UserId

	// Adminのユーザのみ使用可
	if err := common.AdminOnly(ctx, db, userId); err != nil {
		return err
	}

	mode, err := net.GetQuery(r, "mode")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	switch mode {
	case "ip":
		blockIPs, err := models.GetAllBlocIP(ctx, db)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		net.ResponseOK(w, blockIPs)
	case "mail":
		blockMails, err := models.GetAllBlocMail(ctx, db)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		net.ResponseOK(w, blockMails)
	default:
		return status.NewBadRequestError(errors.New("required mode")).Caller()
	}

	return nil
}

func DeleteBlocks(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	c := common.NewCert(w, r)
	if err := c.Login(ctx, db); err != nil {
		return err
	}
	userId := c.UserId

	// Adminのユーザのみ使用可
	if err := common.AdminOnly(ctx, db, userId); err != nil {
		return err
	}

	mode, err := net.GetQuery(r, "mode")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	element, err := net.GetQuery(r, "element")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	switch mode {
	case "ip":
		if err := models.DeleteBlockIP(ctx, db, element); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	case "mail":
		if err := models.DeleteBlockMail(ctx, db, element); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	default:
		return status.NewBadRequestError(errors.New("required mode")).Caller()
	}

	return nil
}
