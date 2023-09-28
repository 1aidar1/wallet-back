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

type OperationWithdrawRequest struct {
	ConsumerId                  string `json:"consumer_id"`
	ServiceProvideId            string `json:"service_provider_id"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
	Amount                      float64 `json:"amount"`
	Description                 string  `json:"description"`
	OrderId                     string  `json:"order_id"`
}

type OperationWithdrawResponse struct {
	TransactionId string     `json:"transaction_id,omitempty"`
	OperationId   string     `json:"operation_id,omitempty"`
	Status        string     `json:"status"`
	ErrorCode     string     `json:"error_code"`
	NeedApprove   bool       `json:"need_approve,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Validation    string     `json:"validation,omitempty"`
}

func (h *Handler) OperationWithdraw(c *gin.Context) {
	var body OperationWithdrawRequest
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

	req := dto.WithdrawReq{
		ConsumerId:       body.ConsumerId,
		ServiceProvideId: body.ServiceProvideId,
		WalletId:         walletId[0],
		Amount:           utils.RestToBusiness(body.Amount),
		Description:      body.Description,
		OrderId:          body.OrderId,
		TransactionId:    nil,
	}
	withdraw, e := h.operation.Withdraw(ctx, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	c.JSON(200, OperationWithdrawResponse{
		TransactionId: withdraw.TransactionId,
		OperationId:   withdraw.Id,
		Status:        withdraw.Status,
		NeedApprove:   false,
		ErrorCode:     withdraw.ErrorCode,
		CreatedAt:     withdraw.CreatedAt,
	})
	return
}
