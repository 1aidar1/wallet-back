package http

import (
	"context"
	"time"

	"git.example.kz/wallet/wallet-back/internal/controller"
	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OperationClearRequest struct {
	ConsumerId                  string `json:"consumer_id" binding:"required"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
	TransactionId               string  `json:"transaction_id" binding:"required"`
	Amount                      float64 `json:"amount"`
}
type OperationClearResponse struct {
	OperationId string     `json:"operation_id,omitempty"`
	Status      string     `json:"status"`
	ErrorCode   string     `json:"error_code"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Validation  string     `json:"validation,omitempty"`
}

func (h *Handler) OperationClear(c *gin.Context) {
	var body OperationClearRequest
	err := c.ShouldBindJSON(&body)
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

	clear, e := h.operation.Clear(ctx, dto.ClearReq{
		ConsumerId:    body.ConsumerId,
		TransactionId: body.TransactionId,
		WalletId:      walletId[0],
		Amount:        utils.RestToBusiness(body.Amount),
	})
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	c.JSON(200, OperationClearResponse{
		OperationId: clear.Id,
		Status:      clear.Status,
		ErrorCode:   clear.ErrorCode,
		CreatedAt:   clear.CreatedAt,
	})
	return
}
