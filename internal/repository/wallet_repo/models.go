package wallet_repo

import (
	"strings"
	"time"

	"git.example.kz/wallet/wallet-back/internal/entity"
)

type WalletHistoryTxModel struct {
	Id               string                        `db:"id" json:"id"`
	ConsumerId       string                        `db:"consumer_id" json:"consumer_id"`
	Type             string                        `db:"type" json:"type"`
	Description      string                        `db:"description" json:"description"`
	ServiceProvideId string                        `db:"service_provider_id" json:"serviceProvideId"`
	OrderId          string                        `db:"order_id" json:"order_id"`
	CreatedAt        time.Time                     `db:"created_at" json:"created_at"`
	Operations       []WalletHistoryOperationModel `db:"operations" json:"operations"`
}

type WalletHistoryOperationModel struct {
	Id                 string      `json:"id"`
	Type               string      `json:"type"`
	WalletId           string      `json:"wallet_id"`
	Amount             int         `json:"amount"`
	Status             string      `json:"status"`
	ErrorCode          string      `json:"error_code"`
	InternalLogMessage string      `json:"-"`
	BalanceBefore      int         `json:"balance_before"`
	BalanceAfter       int         `json:"balance_after"`
	HoldBalanceBefore  int         `json:"hold_balance_before"`
	HoldBalanceAfter   int         `json:"hold_balance_after"`
	CreatedAt          Iso8601Time `json:"created_at"`
}
type Iso8601Time time.Time

func (c *Iso8601Time) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", value) //parse time
	if err != nil {
		return err
	}
	*c = Iso8601Time(t) //set result using the pointer
	return nil
}

func (c Iso8601Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format(time.RFC3339Nano) + `"`), nil
}

func transformToEntitySlice(txModels []WalletHistoryTxModel) []entity.WalletHistory {
	txEntities := make([]entity.WalletHistory, len(txModels))
	for i, txModel := range txModels {
		txEntities[i] = transformToEntity(txModel)
	}
	return txEntities
}
func transformToEntity(walletTxModel WalletHistoryTxModel) entity.WalletHistory {
	var walletTx entity.WalletHistory

	// copy simple fields
	walletTx.Id = walletTxModel.Id
	walletTx.ConsumerId = walletTxModel.ConsumerId
	walletTx.Type = walletTxModel.Type
	walletTx.Description = walletTxModel.Description
	walletTx.OrderId = walletTxModel.OrderId
	walletTx.ServiceProvideId = walletTxModel.ServiceProvideId
	walletTx.CreatedAt = walletTxModel.CreatedAt

	// copy and transform operations
	for _, opModel := range walletTxModel.Operations {
		op := entity.WalletHistoryOperation{
			Id:                 opModel.Id,
			Type:               opModel.Type,
			WalletId:           opModel.WalletId,
			Amount:             opModel.Amount,
			Status:             opModel.Status,
			ErrorCode:          opModel.ErrorCode,
			InternalLogMessage: opModel.InternalLogMessage,
			BalanceBefore:      &opModel.BalanceBefore,
			BalanceAfter:       &opModel.BalanceAfter,
			HoldBalanceBefore:  &opModel.HoldBalanceBefore,
			HoldBalanceAfter:   &opModel.HoldBalanceAfter,
			CreatedAt:          time.Time(opModel.CreatedAt),
		}
		walletTx.Operations = append(walletTx.Operations, op)
	}

	return walletTx
}
