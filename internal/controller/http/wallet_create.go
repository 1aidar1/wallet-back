package http

import (
	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"github.com/gin-gonic/gin"
)

type WalletCreateRequest struct {
	ConsumerId                  string `json:"consumer_id"`
	Currency                    string `json:"currency" binding:"required"`
	Country                     string `json:"country" binding:"required"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
}

type WalletCreateResponse struct {
	entity.WalletIdentifierType `json:"wallet_identifier"`
	Status                      string `json:"status"`
	ErrorCode                   string `json:"error_code"`
}

func (h *Handler) WalletCreate(c *gin.Context) {
	var body WalletCreateRequest
	err := c.BindJSON(&body)
	if err != nil {
		e := dictionary.Get().Errors.Get("validation_err")
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	req := dto.WalletCreateReq{
		ConsumerId:   body.ConsumerId,
		CurrencyCode: body.Currency,
		CountryCode:  body.Country,
		Phone:        body.Phone,
	}
	wallet, e := h.wallet.Create(c, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	c.JSON(200, WalletCreateResponse{
		WalletIdentifierType: entity.WalletIdentifierType{
			Id:    wallet.ID,
			Phone: wallet.Phone,
		},
		Status:    dictionary.Get().Statuses.Success.String(),
		ErrorCode: "",
	})
}
