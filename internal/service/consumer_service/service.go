package consumer_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/jackc/pgx/v5"
)

const (
	// Update

	Method_OperationClear       = "clear"
	Method_OperationRefill      = "refill"
	Method_OperationWithdraw    = "withdraw"
	Method_OperationHold        = "hold"
	Method_OperationUnhold      = "unhold"
	Method_OperationRefund      = "refund"
	Method_OperationTransfer    = "transfer"
	Method_WalletCreate         = "wallet_create"
	Method_WalletClose          = "wallet_close"
	Method_WalletBlock          = "wallet_block"
	Method_WalletUnblock        = "wallet_unblock"
	Method_WalletIdentification = "wallet_identification"

	//Read

	Method_WalletInfo              = "wallet_info"
	Method_WalletHistory           = "wallet_history"
	Method_WalletHistoryByProvider = "wallet_history_by_provider"
	Method_WalletStatistics        = "wallet_statistics"
	Method_TransactionInfo         = "transaction_info"
)

var Methods = []string{
	Method_OperationClear,
	Method_OperationRefill,
	Method_OperationWithdraw,
	Method_OperationHold,
	Method_OperationUnhold,
	Method_OperationRefund,
	Method_OperationTransfer,
	Method_WalletCreate,
	Method_WalletClose,
	Method_WalletBlock,
	Method_WalletUnblock,
	Method_WalletIdentification,
	//	read
	Method_WalletInfo,
	Method_WalletHistory,
	Method_WalletHistoryByProvider,
	Method_WalletStatistics,
	Method_TransactionInfo,
}

type ConsumerRepository interface {
	Create(ctx context.Context, req dto.ConsumerCreateReq, tx pgx.Tx) (entity.Consumer, *errcode.ErrCode)
	Read(ctx context.Context, consumerId string, tx pgx.Tx) (entity.Consumer, *errcode.ErrCode)
	Update(ctx context.Context, id string, req dto.ConsumerUpdateReq, tx pgx.Tx) *errcode.ErrCode
	Delete(ctx context.Context, id string, tx pgx.Tx) *errcode.ErrCode
	List(ctx context.Context, tx pgx.Tx) ([]entity.Consumer, *errcode.ErrCode)
}

type ConsumerService struct {
	repo ConsumerRepository
}

func NewConsumerService(repo ConsumerRepository) *ConsumerService {
	return &ConsumerService{
		repo: repo,
	}
}
