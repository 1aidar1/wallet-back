package http

import (
	"time"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/utils"
	"github.com/gin-gonic/gin"
)

type WalletStatisticsRequest struct {
	ConsumerId                  string `json:"consumer_id" binding:"required"`
	Period                      Period `json:"period"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
}
type Period struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type WalletStatisticsResponse struct {
	Statistics map[string]Stats `json:"statistics"`
	Status     string           `json:"status"`
	ErrorCode  string           `json:"error_code"`
}
type Stats struct {
	Amount float64 `json:"payments_amount"`
	Count  int     `json:"payments_count"`
}

func (h *Handler) WalletStatistics(c *gin.Context) {
	var body WalletStatisticsRequest
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
	req := dto.WalletStatisticsReq{
		ConsumerId: body.ConsumerId,
		WalletId:   walletId,
		DateStart:  body.Period.From,
		DateEnd:    body.Period.To,
		Status:     dictionary.Get().Statuses.Success,
	}
	wallet, e := h.wallet.Statistics(c, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	m := make(map[string]Stats)
	for key, _ := range wallet {
		stat := Stats{
			Amount: utils.BusinessToRest(wallet[key].Amount),
			Count:  wallet[key].Count,
		}
		m[key] = stat
	}

	c.JSON(200, WalletStatisticsResponse{
		Statistics: m,
		Status:     dictionary.Get().Statuses.Success.String(),
		ErrorCode:  "",
	})
}
