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

type OperationRefillRequest struct {
	ConsumerId                  string `json:"consumer_id"`
	ServiceProvideId            string `json:"service_provider_id"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
	Amount                      float64 `json:"amount"`
	Description                 string  `json:"description"`
	OrderId                     string  `json:"order_id"`
}

type OperationRefillResponse struct {
	TransactionId string     `json:"transaction_id,omitempty"`
	OperationId   string     `json:"operation_id,omitempty"`
	Status        string     `json:"status"`
	ErrorCode     string     `json:"error_code"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Validation    string     `json:"validation,omitempty"`
}

func (h *Handler) OperationRefill(c *gin.Context) {
	var body OperationRefillRequest
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

	req := dto.RefillReq{
		ServiceProvideId: body.ServiceProvideId,
		ConsumerId:       body.ConsumerId,
		WalletId:         walletId[0],
		Amount:           utils.RestToBusiness(body.Amount),
		Description:      body.Description,
		OrderId:          body.OrderId,
	}
	refill, e := h.operation.Refill(ctx, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	c.JSON(200, OperationRefillResponse{
		TransactionId: refill.TransactionId,
		OperationId:   refill.Id,
		Status:        refill.Status,
		ErrorCode:     refill.ErrorCode,
		CreatedAt:     refill.CreatedAt,
	})
	return
}
