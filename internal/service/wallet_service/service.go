package wallet_service

import (
	"context"
	"time"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, *errcode.ErrCode)
	UpdateStatus(ctx context.Context, wallet dto.UpdateWalletStatus, tx pgx.Tx) *errcode.ErrCode
	UpdateIdentification(ctx context.Context, walletId, identification string, tx pgx.Tx) *errcode.ErrCode
	Find(ctx context.Context, id string, tx pgx.Tx) (entity.Wallet, *errcode.ErrCode)
	FindByPhone(ctx context.Context, phone string, tx pgx.Tx) (entity.Wallet, *errcode.ErrCode)
	Stat(ctx context.Context, req dto.WalletStatisticsReq) (dto.WalletStatisticsRes, *errcode.ErrCode)
	History(ctx context.Context, history dto.WalletHistoryReq, tx pgx.Tx) ([]entity.WalletHistory, *errcode.ErrCode)
}

type OperationService interface {
	TransferWithTX(ctx context.Context, transfer dto.TransferReq, tx pgx.Tx) (entity.TransferOperation, *errcode.ErrCode)
}
type AtomicWrapper interface {
	Do(ctx context.Context, f func(tx pgx.Tx) *errcode.ErrCode) *errcode.ErrCode
}
type ConsumerService interface {
	Validate(ctx context.Context, req dto.ConsumerValidateReq) *errcode.ErrCode
}

type WalletService struct {
	operations OperationService
	repo       WalletRepository
	atomic     AtomicWrapper
	consumer   ConsumerService
	redis      *redis.Client // использую напрямую client т.к. существует miniredis
	lockTime   time.Duration // сколько по времени лочим wallet внутри редиса, если вдруг не разлочим.
}

func NewWalletService(op OperationService, repo WalletRepository, tx AtomicWrapper, consumer ConsumerService, redis *redis.Client) *WalletService {
	return &WalletService{
		operations: op,
		repo:       repo,
		atomic:     tx,
		redis:      redis,
		consumer:   consumer,
		lockTime:   time.Second * 15,
	}
}

func (s *WalletService) GetId(ctx context.Context, data entity.WalletIdentifierType) (string, *errcode.ErrCode) {

	if data.Id != "" {
		return data.Id, nil
	} else if data.Phone != "" {
		wallet, err := s.repo.FindByPhone(ctx, data.Phone, nil)
		if err != nil {
			logger.GetProdLogger().Error(err.Log())
			return "", errcode.New("wallet_not_found")
		}
		return wallet.ID, nil
	}
	return "", errcode.New("wallet_not_found")
}
