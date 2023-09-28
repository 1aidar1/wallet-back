package entity

import (
	"time"
)

type Transaction struct {
	Id               string      `json:"id" db:"id"`
	ConsumerId       string      `json:"consumer_id" db:"consumer_id"`
	ServiceProvideId string      `json:"serviceProvideId" db:"service_provider_id"`
	Type             string      `json:"type" db:"type"`
	Description      string      `json:"description" db:"description"`
	OrderId          string      `json:"order_id" db:"order_id"`
	CreatedAt        time.Time   `json:"created_at" db:"created_at"`
	Operations       []Operation `json:"operations" db:"operations"`
}

//
//func (t	 Transaction) MarshalToRest() json.RawMessage {
//	for _, operation := range t.Operations {
//		operation.Amount =
//	}
//}
