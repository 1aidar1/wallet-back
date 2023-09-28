package controller

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
)

type WalletService interface {
	WalletLockService
	GetId(ctx context.Context, data entity.WalletIdentifierType) (string, *errcode.ErrCode)
	Create(ctx context.Context, wallet dto.WalletCreateReq) (entity.Wallet, *errcode.ErrCode)
	Close(ctx context.Context, wallet dto.WalletCloseReq) (entity.OperationStatus, *errcode.ErrCode)
	Block(ctx context.Context, block dto.WalletBlockReq) (entity.OperationStatus, *errcode.ErrCode)
	Unblock(ctx context.Context, unblock dto.WalletUnblockReq) (entity.OperationStatus, *errcode.ErrCode)
	Info(ctx context.Context, req dto.WalletInfoReq) (entity.Wallet, *errcode.ErrCode)
	History(ctx context.Context, history dto.WalletHistoryReq) ([]entity.WalletHistory, *errcode.ErrCode)
	HistoryByProvider(ctx context.Context, req dto.WalletHistoryReq) ([]entity.WalletHistory, *errcode.ErrCode)
	Identification(ctx context.Context, identification dto.WalletIdentificationReq) (entity.OperationStatus, *errcode.ErrCode)
	Statistics(ctx context.Context, req dto.WalletStatisticsReq) (dto.WalletStatisticsRes, *errcode.ErrCode)
}
type WalletLockService interface {
	WaitAndLock(ctx context.Context, walletIds ...string) *errcode.ErrCode
	Lock(ctx context.Context, walletIds ...string) (bool, *errcode.ErrCode)
	Unlock(ctx context.Context, walletIds ...string) *errcode.ErrCode
	AreUnlocked(ctx context.Context, walletIds ...string) (bool, *errcode.ErrCode)
}

type OperationService interface {
	Refill(ctx context.Context, refill dto.RefillReq) (entity.Operation, *errcode.ErrCode)
	Withdraw(ctx context.Context, withdraw dto.WithdrawReq) (entity.Operation, *errcode.ErrCode)
	Hold(ctx context.Context, hold dto.HoldReq) (entity.Operation, *errcode.ErrCode)
	Unhold(ctx context.Context, unhold dto.UnholdReq) (entity.Operation, *errcode.ErrCode)
	Clear(ctx context.Context, clear dto.ClearReq) (entity.Operation, *errcode.ErrCode)
	Refund(ctx context.Context, clear dto.RefundReq) (entity.Operation, *errcode.ErrCode)
	Transfer(ctx context.Context, transfer dto.TransferReq) (entity.TransferOperation, *errcode.ErrCode)
	FindTransaction(ctx context.Context, req dto.FindTransactionReq) (entity.Transaction, *errcode.ErrCode)
}

type ConsumerService interface {
	Create(ctx context.Context, req dto.ConsumerCreateReq) (entity.Consumer, *errcode.ErrCode)
	List(ctx context.Context) ([]entity.Consumer, *errcode.ErrCode)
	Read(ctx context.Context, id string) (entity.Consumer, *errcode.ErrCode)
	Update(ctx context.Context, consumerId string, req dto.ConsumerUpdateReq) *errcode.ErrCode
	Delete(ctx context.Context, id string) *errcode.ErrCode
	MethodList(ctx context.Context) []string
}
