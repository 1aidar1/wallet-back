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

type WalletBlockRequest struct {
	ConsumerId                  string `json:"consumer_id" binding:"required"`
	Description                 string `json:"description"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
}

type WalletBlockResponse struct {
	Status    string `json:"status"`
	ErrorCode string `json:"error_code"`
}

func (h *Handler) WalletBlock(c *gin.Context) {
	var body WalletBlockRequest
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

	status, e := h.wallet.Block(c, dto.WalletBlockReq{
		ConsumerId:  body.ConsumerId,
		WalletId:    walletId[0],
		Description: body.Description,
	})
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	c.JSON(200, WalletBlockResponse{
		Status:    status.String(),
		ErrorCode: "",
	})
}
