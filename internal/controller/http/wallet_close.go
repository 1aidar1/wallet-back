package http

import (
	"context"
	"time"

	"git.example.kz/wallet/wallet-back/internal/controller"
	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"github.com/gin-gonic/gin"
)

type WalletCloseRequest struct {
	ConsumerId                    string                      `json:"consumer_id"`
	WalletIdentifier              entity.WalletIdentifierType `json:"wallet_identifier"`
	CorrespondingWalletIdentifier entity.WalletIdentifierType `json:"corresponding_wallet_identifier"`
	Description                   string                      `json:"description"`
}

type WalletCloseResponse struct {
	Status    string `json:"status"`
	ErrorCode string `json:"error_code"`
}

func (h *Handler) WalletClose(c *gin.Context) {
	var body WalletCloseRequest
	err := c.BindJSON(&body)
	if err != nil {
		e := dictionary.Get().Errors.Get(err.Error())
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, e := h.getWalletIdAndLock(ctx, body.WalletIdentifier, body.CorrespondingWalletIdentifier)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	defer h.wallet.Unlock(c, walletId...)

	status, e := h.wallet.Close(c, dto.WalletCloseReq{
		ConsumerId:            body.ConsumerId,
		WalletId:              walletId[0],
		CorrespondingWalletId: walletId[1],
		Description:           body.Description,
	})
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	c.JSON(200, WalletCloseResponse{
		Status:    status.String(),
		ErrorCode: "",
	})
}
