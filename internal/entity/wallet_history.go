package entity

import "time"

type WalletHistory struct {
	Id               string                   `db:"id" json:"id"`
	ConsumerId       string                   `db:"consumer_id" json:"consumer_id"`
	Type             string                   `db:"type" json:"type"`
	Description      string                   `db:"description" json:"description"`
	ServiceProvideId string                   `db:"service_provider_id" json:"serviceProvideId"`
	OrderId          string                   `db:"order_id" json:"order_id"`
	CreatedAt        time.Time                `db:"created_at" json:"created_at"`
	Operations       []WalletHistoryOperation `db:"operations" json:"operations"`
}

type WalletHistoryOperation struct {
	Id                 string    `json:"id"`
	Type               string    `json:"type"`
	WalletId           string    `json:"wallet_id"`
	Amount             int       `json:"amount"`
	Status             string    `json:"status"`
	ErrorCode          string    `json:"error_code"`
	InternalLogMessage string    `json:"-"`
	BalanceBefore      *int      `json:"balance_before,omitempty"`
	BalanceAfter       *int      `json:"balance_after,omitempty"`
	HoldBalanceBefore  *int      `json:"hold_balance_before,omitempty"`
	HoldBalanceAfter   *int      `json:"hold_balance_after,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
}
