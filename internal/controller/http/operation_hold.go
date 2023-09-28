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

type OperationHoldRequest struct {
	ConsumerId                  string `json:"consumer_id"`
	ServiceProvideId            string `json:"service_provider_id"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
	Amount                      float64 `json:"amount"`
	Description                 string  `json:"description"`
	OrderId                     string  `json:"order_id"`
}

type OperationHoldResponse struct {
	TransactionId string     `json:"transaction_id,omitempty"`
	OperationId   string     `json:"operation_id,omitempty"`
	Status        string     `json:"status"`
	ErrorCode     string     `json:"error_code"`
	NeedApprove   string     `json:"need_approve,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
}

func (h *Handler) OperationHold(c *gin.Context) {
	var body OperationHoldRequest
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

	req := dto.HoldReq{
		ConsumerId:       body.ConsumerId,
		ServiceProvideId: body.ServiceProvideId,
		WalletId:         walletId[0],
		Amount:           utils.RestToBusiness(body.Amount),
		Description:      body.Description,
		OrderId:          body.OrderId,
	}

	hold, e := h.operation.Hold(ctx, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	c.JSON(200, OperationHoldResponse{
		TransactionId: hold.TransactionId,
		OperationId:   hold.Id,
		Status:        hold.Status,
		ErrorCode:     hold.ErrorCode,
		CreatedAt:     hold.CreatedAt,
	})
	return
}
