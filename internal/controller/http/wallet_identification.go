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

type WalletIdentificationRequest struct {
	ConsumerId                  string `json:"consumer_id" binding:"required"`
	Status                      string `json:"status" binding:"required"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
}

type WalletIdentificationResponse struct {
	Status    string `json:"status"`
	ErrorCode string `json:"error_code"`
}

func (h *Handler) WalletIdentification(c *gin.Context) {
	var body WalletIdentificationRequest
	err := c.BindJSON(&body)
	if err != nil {
		e := dictionary.Get().Errors.Get("validation_err")
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()

	walletId, e := h.getWalletIdAndLock(ctx, body.WalletIdentifierType)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	defer h.wallet.Unlock(c, walletId...)

	req := dto.WalletIdentificationReq{
		ConsumerId:     body.ConsumerId,
		WalletId:       walletId[0],
		Identification: body.Status,
	}
	status, e := h.wallet.Identification(c, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	c.JSON(200, WalletIdentificationResponse{
		Status:    status.String(),
		ErrorCode: "",
	})
}
