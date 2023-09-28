package dto

import (
	"time"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
)

// *** CreateWallet ***

type WalletCreateReq struct {
	ConsumerId   string
	CurrencyCode string
	CountryCode  string
	Phone        string
}

func (r *WalletCreateReq) Validate() *errcode.ErrCode {
	if c := dictionary.Get().Countries.ByCode(r.CountryCode); c == entity.EmptyCountry {
		return errcode.New("no_such_country")
	}
	if c := dictionary.Get().Currencies.ByCode(r.CurrencyCode); c == entity.EmptyCurrency {
		return errcode.New("no_such_currency")
	}
	return nil
}

// *** UpdateWalletStatus ***

type UpdateWalletStatus struct {
	WalletId    string
	NewStatus   string
	Description string
}

// *** BlockWallet ***

type WalletBlockReq struct {
	ConsumerId  string
	WalletId    string
	Description string
}

// *** UnblockWallet ***

type WalletUnblockReq struct {
	ConsumerId  string
	WalletId    string
	Description string
}

// *** CloseWallet ***

type WalletCloseReq struct {
	ConsumerId            string
	WalletId              string
	CorrespondingWalletId string
	Description           string
}

// *** WalletIdentification ***

type WalletIdentificationReq struct {
	ConsumerId     string
	WalletId       string
	Identification string
}

// *** WalletInfo ***

type WalletInfoReq struct {
	ConsumerId string
	WalletId   string
}

func (req *WalletIdentificationReq) Validate() *errcode.ErrCode {
	identifications := dictionary.Get().WalletIdentification
	if req.Identification != identifications.None.String() && req.Identification != identifications.Full.String() && req.Identification != identifications.Basic.String() {
		return errcode.New("no_such_identification")
	}
	return nil
}

// *** WalletHistory ***

type WalletHistoryReq struct {
	ConsumerId string
	WalletId   string
	ProviderId *string
	DateStart  time.Time
	DateEnd    time.Time
	PerPage    int
	Page       int
}

func (r *WalletHistoryReq) Validate() *errcode.ErrCode {
	if r.DateStart.After(r.DateEnd) {
		return errcode.New("validation_err")
	}
	//if r.DateEnd.Sub(r.DateStart).Hours() > time.Hour.Hours()*24*365 {
	//	return errcode.New("exceed_max_timespan")
	//}
	if r.PerPage > 100 || r.PerPage <= 0 {
		r.PerPage = 100
	}
	if r.Page <= 0 {
		r.Page = 1
	}
	return nil
}

// *** Statistics ***

type WalletStatisticsReq struct {
	ConsumerId string
	WalletId   string
	DateStart  time.Time
	DateEnd    time.Time
	Status     entity.OperationStatus
}

func (r *WalletStatisticsReq) Validate() *errcode.ErrCode {
	now := time.Now()
	year, month, _ := now.Date()
	if (r.DateStart == time.Time{}) {
		beginningOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
		r.DateStart = beginningOfMonth
	}
	if (r.DateEnd == time.Time{}) {
		beginningOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())
		r.DateStart = beginningOfNextMonth
	}
	return nil
}

type WalletStatisticsRes map[string]Stats

type Stats struct {
	Amount int `json:"payments_amount"`
	Count  int `json:"payments_count"`
}
