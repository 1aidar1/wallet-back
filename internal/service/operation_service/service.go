package operation_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository interface {
	Create(ctx context.Context, tr entity.Transaction, tx pgx.Tx) (string, *errcode.ErrCode)
	Find(ctx context.Context, id string, tx pgx.Tx) (entity.Transaction, *errcode.ErrCode)
	FindByOrder(ctx context.Context, orderId, providerCode string, tx pgx.Tx) (entity.Transaction, *errcode.ErrCode)
}
type OperationRepository interface {
	Create(ctx context.Context, operation entity.Operation, tx pgx.Tx) (entity.Operation, *errcode.ErrCode)
}
type WalletRepository interface {
	Find(ctx context.Context, id string, tx pgx.Tx) (entity.Wallet, *errcode.ErrCode)
	AddBalance(ctx context.Context, walletId string, balance, hold int, tx pgx.Tx) *errcode.ErrCode
}
type AtomicWrapper interface {
	Do(ctx context.Context, f func(tx pgx.Tx) *errcode.ErrCode) *errcode.ErrCode
}

type ConsumerService interface {
	Validate(ctx context.Context, req dto.ConsumerValidateReq) *errcode.ErrCode
}

type OperationService struct {
	walletRepo      WalletRepository
	transactionRepo TransactionRepository
	operationRepo   OperationRepository
	atomicRepo      AtomicWrapper
	consumerService ConsumerService
}

func NewOperationService(operation OperationRepository, transaction TransactionRepository, wallet WalletRepository, tx AtomicWrapper, consumer ConsumerService) *OperationService {
	return &OperationService{
		walletRepo:      wallet,
		transactionRepo: transaction,
		operationRepo:   operation,
		atomicRepo:      tx,
		consumerService: consumer,
	}
}
