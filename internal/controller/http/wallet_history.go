package http

import (
	"time"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/utils"
	"github.com/gin-gonic/gin"
)

type WalletHistoryRequest struct {
	ConsumerId string    `json:"consumer_id"`
	DateStart  time.Time `json:"date_start" binding:"required"`
	DateEnd    time.Time `json:"date_end" binding:"required"`
	Pagination struct {
		ItemsPerPage int `json:"items_per_page"`
		Page         int `json:"page"`
	} `json:"pagination"`
	entity.WalletIdentifierType `json:"wallet_identifier"`
}

type WalletHistoryResponse struct {
	TotalItemsCount int             `json:"total_items_count"`
	Items           []WalletHistory `json:"items"`
}
type WalletHistory struct {
	Id               string                   `json:"id"`
	ConsumerId       string                   `json:"consumer_id"`
	Type             string                   `json:"type"`
	Description      string                   `json:"description"`
	ServiceProvideId string                   `json:"serviceProvideId"`
	OrderId          string                   `json:"order_id"`
	CreatedAt        time.Time                `json:"created_at"`
	Operations       []WalletHistoryOperation `json:"operations"`
}

type WalletHistoryOperation struct {
	Id                string    `json:"id"`
	Type              string    `json:"type"`
	WalletId          string    `json:"wallet_id"`
	Amount            float64   `json:"amount"`
	Status            string    `json:"status"`
	ErrorCode         string    `json:"error_code"`
	BalanceBefore     *float64  `json:"balance_before,omitempty"`
	BalanceAfter      *float64  `json:"balance_after,omitempty"`
	HoldBalanceBefore *float64  `json:"hold_balance_before,omitempty"`
	HoldBalanceAfter  *float64  `json:"hold_balance_after,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}

func (h *Handler) WalletHistory(c *gin.Context) {
	var body WalletHistoryRequest
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
	req := dto.WalletHistoryReq{
		ConsumerId: body.ConsumerId,
		WalletId:   walletId,
		DateStart:  body.DateStart,
		DateEnd:    body.DateEnd,
		Page:       body.Pagination.Page,
		PerPage:    body.Pagination.ItemsPerPage,
	}
	history, e := h.wallet.History(c, req)
	if e != nil {
		e = dictionary.Get().Errors.Get(e.Code)
		c.JSON(e.GetHttp(), defaultHttpError(e.Code))
		return
	}
	transactions := make([]WalletHistory, len(history))
	for i, t := range history {
		operations := make([]WalletHistoryOperation, len(t.Operations))
		for j, operation := range history[i].Operations {
			operations[j] = WalletHistoryOperation{
				Id:                operation.Id,
				Type:              operation.Type,
				WalletId:          operation.WalletId,
				Amount:            utils.BusinessToRest(operation.Amount),
				Status:            operation.Status,
				ErrorCode:         operation.ErrorCode,
				BalanceBefore:     utils.BusinessToRestP(operation.BalanceBefore),
				BalanceAfter:      utils.BusinessToRestP(operation.BalanceAfter),
				HoldBalanceBefore: utils.BusinessToRestP(operation.HoldBalanceBefore),
				HoldBalanceAfter:  utils.BusinessToRestP(operation.HoldBalanceAfter),
				CreatedAt:         operation.CreatedAt,
			}
		}
		transactions[i] = WalletHistory{
			Id:               t.Id,
			ConsumerId:       t.ConsumerId,
			Type:             t.Type,
			Description:      t.Description,
			ServiceProvideId: t.ServiceProvideId,
			OrderId:          t.OrderId,
			CreatedAt:        t.CreatedAt,
			Operations:       operations,
		}

	}
	c.JSON(200, WalletHistoryResponse{
		TotalItemsCount: len(history),
		Items:           transactions,
	})
}
