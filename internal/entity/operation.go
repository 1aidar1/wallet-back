package entity

import "time"

type Operation struct {
	Id                 string     `json:"id" db:"id"`
	TransactionId      string     `json:"transaction_id" db:"transaction_id"`
	Type               string     `json:"type" db:"type"`
	WalletId           string     `json:"wallet_id" db:"wallet_id"`
	Amount             int        `json:"amount" db:"amount"`
	Status             string     `json:"status" db:"status"`
	ErrorCode          string     `json:"error_code" db:"error_code"`
	InternalLogMessage string     `json:"internal_log_message" db:"internal_log_message"`
	BalanceBefore      int        `json:"balance_before" db:"balance_before"`
	BalanceAfter       int        `json:"balance_after" db:"balance_after"`
	HoldBalanceBefore  int        `json:"hold_balance_before" db:"hold_balance_before"`
	HoldBalanceAfter   int        `json:"hold_balance_after" db:"hold_balance_after"`
	CreatedAt          *time.Time `json:"created_at" db:"created_at"`
}

type TransferOperation struct {
	TransactionId string
	Status        string
	ErrorCode     string
	CreatedAt     *time.Time
}
