package operation_service

import (
	"context"
	"fmt"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"github.com/jackc/pgx/v5"
)

func (s *OperationService) Transfer(ctx context.Context, req dto.TransferReq) (operation entity.TransferOperation, err *errcode.ErrCode) {
	if e := req.Validate(); e != nil {
		return entity.TransferOperation{}, errcode.New(e.Error())
	}

	if err = s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_OperationTransfer,
	}); err != nil {
		return entity.TransferOperation{}, err
	}

	f := func(tx pgx.Tx) *errcode.ErrCode {
		operation, err = s.TransferWithTX(ctx, req, tx)
		if err != nil {
			return err
		}
		return nil
	}

	fErr := s.atomicRepo.Do(ctx, f)
	return operation, fErr
}

// Трансфер с опрокинутой транзакцией
// Под капотом это две операции внутри одной транзакции, снятие с одного кошелька и пополнение второго.
func (s *OperationService) TransferWithTX(ctx context.Context, req dto.TransferReq, tx pgx.Tx) (entity.TransferOperation, *errcode.ErrCode) {
	// Preload
	fromWallet, err := s.walletRepo.Find(ctx, req.WithdrawWalletId, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.TransferOperation{}, errcode.New("wallet_not_found")
	}
	toWallet, err := s.walletRepo.Find(ctx, req.RefillWalletId, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.TransferOperation{}, errcode.New("wallet_not_found")
	}

	// Check
	if toWallet.Status != dictionary.Get().WalletStatuses.Active.String() {
		return entity.TransferOperation{}, errcode.New("wallet_not_active")
	}
	if fromWallet.Status != dictionary.Get().WalletStatuses.Active.String() {
		return entity.TransferOperation{}, errcode.New("wallet_not_active")
	}
	if fromWallet.CurrencyID != toWallet.CurrencyID {
		return entity.TransferOperation{}, errcode.New("wallet_currency_mismatch")
	}
	if fromWallet.Balance < req.Amount {
		return entity.TransferOperation{}, errcode.New("not_enough_balance")
	}

	// Logic
	transactionId, err := s.transactionRepo.Create(ctx, entity.Transaction{
		ConsumerId:       req.ConsumerId,
		ServiceProvideId: req.ServiceProvideId,
		Type:             dictionary.Get().TransactionTypes.Transfer.String(),
		Description:      req.Description,
		OrderId:          req.OrderId,
	}, tx)
	if err != nil {
		logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
		return entity.TransferOperation{}, err
	}

	withdraw, err := s.withdraw(ctx, dto.WithdrawReq{
		WalletId:      req.WithdrawWalletId,
		Amount:        req.Amount,
		TransactionId: &transactionId,
	}, tx)
	if err != nil {
		logger.GetProdLogger().Error(fmt.Sprintf("withdraw failed err: %s", err.Log()))
		return entity.TransferOperation{}, err
	}

	refill, err := s.refill(ctx, dto.RefillReq{
		WalletId:      req.RefillWalletId,
		Amount:        req.Amount,
		TransactionId: &transactionId,
	}, tx)
	if err != nil {
		logger.GetProdLogger().Error(fmt.Sprintf("refill failed err: %s", err.Log()))
		return entity.TransferOperation{}, err
	}

	// withdraw и refill оба удались
	if withdraw.Status == dictionary.Get().Statuses.Success.String() && refill.Status == dictionary.Get().Statuses.Success.String() {
		operation := entity.TransferOperation{
			TransactionId: transactionId,
			Status:        dictionary.Get().Statuses.Success.String(),
			ErrorCode:     "",
			CreatedAt:     refill.CreatedAt,
		}
		return operation, nil
	} else {
		return entity.TransferOperation{}, errcode.NewDefaultErr()
	}
}
