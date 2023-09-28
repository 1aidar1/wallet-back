package dto

import (
	"errors"
	"github.com/segmentio/asm/ascii"
)

// *** Clear ***

type ClearReq struct {
	ConsumerId    string
	TransactionId string
	WalletId      string
	Amount        int
}

func (r *ClearReq) Validate() error {
	if r.Amount < 0 {
		return errors.New("invalid_operation_amount")
	}
	return nil
}

// *** Refill ***

type RefillReq struct {
	ConsumerId       string
	ServiceProvideId string
	WalletId         string
	Amount           int
	Description      string
	OrderId          string
	TransactionId    *string // Выставляем, если хотим добавить операцию к транзакции, иначе создается под капотом.
}

func (r *RefillReq) Validate() error {
	if !ascii.Valid([]byte(r.Description)) {
		return errors.New("bad_description")
	}
	if r.Amount < 0 {
		return errors.New("invalid_operation_amount")
	}
	return nil
}

// *** Withdraw ***

type WithdrawReq struct {
	ConsumerId       string
	ServiceProvideId string
	WalletId         string
	Amount           int
	Description      string  // Transaction data
	OrderId          string  // Transaction data
	TransactionId    *string // Выставляем, если хотим добавить операцию к транзакции, иначе создается под капотом.
}

func (r *WithdrawReq) Validate() error {
	if !ascii.Valid([]byte(r.Description)) {
		return errors.New("bad_description")
	}
	if r.Amount < 0 {
		return errors.New("invalid_operation_amount")
	}
	return nil
}

// *** Hold ***

type HoldReq struct {
	ConsumerId       string
	ServiceProvideId string
	WalletId         string
	Amount           int
	Description      string
	OrderId          string
}

func (r *HoldReq) Validate() error {
	if !ascii.Valid([]byte(r.Description)) {
		return errors.New("bad_description")
	}
	if r.Amount < 0 {
		return errors.New("invalid_operation_amount")
	}
	return nil
}

// *** Unhold ***

type UnholdReq struct {
	ConsumerId    string
	WalletId      string
	TransactionId string
}

// *** Refund ***

type RefundReq struct {
	ConsumerId    string
	TransactionId string
	WalletId      string
}

// *** Transfer ***

type TransferReq struct {
	ConsumerId       string
	ServiceProvideId string
	WithdrawWalletId string
	RefillWalletId   string
	Amount           int
	Description      string
	OrderId          string
}

func (r *TransferReq) Validate() error {
	if r.Amount < 0 {
		return errors.New("invalid_operation_amount")
	}
	return nil
}

// *** FindTransaction ***

type FindTransactionReq struct {
	ConsumerId       string
	TransactionId    string
	OrderId          string
	ServiceProvideId string
}
