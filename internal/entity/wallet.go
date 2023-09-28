package entity

import (
	"time"
)

type Wallet struct {
	ID             string
	Phone          string
	CurrencyID     string
	CountryID      string
	Balance        int
	Hold           int
	Status         string
	Identification string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type WalletIdentifierType struct {
	Id    string `json:"id"` // wallet id
	Phone string `json:"phone"`
}
