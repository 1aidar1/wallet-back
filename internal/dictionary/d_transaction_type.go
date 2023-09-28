package dictionary

import "git.example.kz/wallet/wallet-back/internal/entity"

type TransactionTypes struct {
	Pay      entity.TransactionType
	Transfer entity.TransactionType
}

func initTransactionTypesEnums() TransactionTypes {
	//enum'ы с базы
	types := TransactionTypes{
		Pay:      "pay",
		Transfer: "transfer",
	}

	return types
}
