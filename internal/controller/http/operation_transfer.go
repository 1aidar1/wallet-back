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

type OperationTransferRequest struct {
	ConsumerId       string                      `json:"consumer_id"`
	ServiceProvideId string                      `json:"service_provider_id"`
	WithdrawWallet   entity.WalletIdentifierType `json:"withdraw_wallet_identifier"`
	RefillWallet     entity.WalletIdentifierType `json:"refill_wallet_identifier"`
	Amount           float64                     `json:"amount" binding:"required,gte=0"`
	Description      string                      `json:"description"`
	OrderId          string                      `json:"order_id"`
}

type OperationTransferResponse struct {
	TransactionId string     `json:"transaction_id,omitempty"`
	Status        string     `json:"status"`
	ErrorCode     string     `json:"error_code"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Validation    string     `json:"validation,omitempty"`
}

func (h *Handler) OperationTransfer(c *gin.Context) {
	var body OperationTransferRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := dictionary.Get().Errors.Get("validation_err")
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, e := h.getWalletIdAndLock(ctx, body.WithdrawWallet, body.RefillWallet)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	defer h.wallet.Unlock(c, walletId...)

	req := dto.TransferReq{
		ConsumerId:       body.ConsumerId,
		ServiceProvideId: body.ServiceProvideId,
		WithdrawWalletId: walletId[0],
		RefillWalletId:   walletId[1],
		Amount:           utils.RestToBusiness(body.Amount),
		Description:      body.Description,
		OrderId:          body.OrderId,
	}
	transfer, e := h.operation.Transfer(ctx, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	c.JSON(200, OperationTransferResponse{
		TransactionId: transfer.TransactionId,
		Status:        transfer.Status,
		ErrorCode:     transfer.ErrorCode,
		CreatedAt:     transfer.CreatedAt,
	})
	return
}
