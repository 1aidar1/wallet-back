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

type OperationRefundRequest struct {
	ConsumerId                  string `json:"consumer_id" binding:"required"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
	TransactionId               string `json:"transaction_id" binding:"required"`
}

type OperationRefundResponse struct {
	TransactionId string     `json:"transaction_id,omitempty"`
	OperationId   string     `json:"operation_id,omitempty"`
	Status        string     `json:"status"`
	ErrorCode     string     `json:"error_code"`
	NeedApprove   string     `json:"need_approve,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Validation    string     `json:"validation,omitempty"`
}

func (h *Handler) OperationRefund(c *gin.Context) {
	var body OperationRefundRequest
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

	refund, e := h.operation.Refund(ctx, dto.RefundReq{
		ConsumerId:    body.ConsumerId,
		TransactionId: body.TransactionId,
		WalletId:      walletId[0],
	})
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	c.JSON(200, OperationRefundResponse{
		TransactionId: refund.TransactionId,
		OperationId:   refund.Id,
		Status:        refund.Status,
		ErrorCode:     refund.ErrorCode,
		CreatedAt:     refund.CreatedAt,
	})
	return
}
