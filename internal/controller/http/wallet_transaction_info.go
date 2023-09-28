package http

import (
	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"github.com/gin-gonic/gin"
)

type TransactionInfoRequest struct {
	ConsumerId                  string `json:"consumer_id" binding:"required"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
	FindBy                      `json:"find_by"`
}
type FindBy struct {
	TransactionId string `json:"transaction_id"`
	Order         `json:"order"`
}
type Order struct {
	ServiceProvideId string `json:"service_provider_id"`
	OrderId          string `json:"order_id"`
}

type TransactionInfoResponse entity.Transaction

func (h *Handler) TransactionInfo(c *gin.Context) {
	var body TransactionInfoRequest
	err := c.BindJSON(&body)
	if err != nil {
		e := dictionary.Get().Errors.Get("validation_err")
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	transaction, e := h.operation.FindTransaction(c, dto.FindTransactionReq{
		ConsumerId:       body.ConsumerId,
		TransactionId:    body.TransactionId,
		OrderId:          body.OrderId,
		ServiceProvideId: body.ServiceProvideId,
	})
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	c.JSON(200, TransactionInfoResponse(transaction))
}
