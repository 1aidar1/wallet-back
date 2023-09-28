package dictionary

import "git.example.kz/wallet/wallet-back/internal/entity"

type OperationTypes struct {
	Refill   entity.OperationType
	Withdraw entity.OperationType
	Hold     entity.OperationType
	Clear    entity.OperationType
	Unhold   entity.OperationType
	Refund   entity.OperationType
}

func initOperationTypesEnums() OperationTypes {
	//enum'ы с базы
	types := OperationTypes{
		Refill:   "refill",
		Withdraw: "withdraw",
		Hold:     "hold",
		Clear:    "clear",
		Unhold:   "unhold",
		Refund:   "refund",
	}

	return types
}
