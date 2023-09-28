package dictionary

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

var (
	dictOnce         sync.Once
	globalDictionary *Dictionary
)

type Dictionary struct {
	Statuses                Statuses
	OperationTypes          OperationTypes
	TransactionTypes        TransactionTypes
	Currencies              Currencies
	Countries               Countries
	WalletStatuses          WalletStatuses
	WalletIdentification    WalletIdentification
	Errors                  CodeErrors
	DefaultServiceProvideId string
}

// TODO: update dictionary with some interval
func CreateOrUpdate(pool *pgxpool.Pool) *Dictionary {
	dictOnce.Do(func() {
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			panic(err)
		}
		defer conn.Release()
		globalDictionary = &Dictionary{
			Statuses:                initStatusesEnums(),
			OperationTypes:          initOperationTypesEnums(),
			TransactionTypes:        initTransactionTypesEnums(),
			Currencies:              initCurrencies(conn),
			Countries:               initCountries(conn),
			WalletStatuses:          initWalletStatusesEnums(),
			WalletIdentification:    initWalletIdentificationEnums(),
			Errors:                  initErrors(conn),
			DefaultServiceProvideId: "server",
		}
	})
	return globalDictionary
}
func Get() *Dictionary {
	return globalDictionary
}
