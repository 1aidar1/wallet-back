package http

import (
	"time"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"github.com/gin-gonic/gin"
)

type WalletInfoRequest struct {
	ConsumerId                  string `json:"consumer_id"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
}

type WalletInfoResponse struct {
	entity.WalletIdentifierType `json:"wallet_identifier"`
	Currency                    string     `json:"currency"`
	Country                     string     `json:"country"`
	WalletStatus                string     `json:"wallet_status"`
	Balance                     float64    `json:"balance"`
	HoldBalance                 float64    `json:"hold_balance"`
	IdentificationStatus        string     `json:"identification_status"`
	CreatedAt                   *time.Time `json:"created_at"`
}

func (h *Handler) WalletInfo(c *gin.Context) {
	var body WalletInfoRequest
	err := c.BindJSON(&body)
	if err != nil {
		e := dictionary.Get().Errors.Get("validation_err")
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	walletId, e := h.wallet.GetId(c, body.WalletIdentifierType)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	wallet, e := h.wallet.Info(c, dto.WalletInfoReq{
		ConsumerId: body.ConsumerId,
		WalletId:   walletId,
	})
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}

	c.JSON(200, WalletInfoResponse{
		WalletIdentifierType: entity.WalletIdentifierType{
			Id:    wallet.ID,
			Phone: wallet.Phone,
		},
		Currency:             dictionary.Get().Currencies.ById(wallet.CurrencyID).Code,
		Country:              dictionary.Get().Countries.ById(wallet.CountryID).Code,
		WalletStatus:         wallet.Status,
		Balance:              float64(wallet.Balance) / 100,
		HoldBalance:          float64(wallet.Hold) / 100,
		IdentificationStatus: wallet.Identification,
		CreatedAt:            wallet.CreatedAt,
	})
}
