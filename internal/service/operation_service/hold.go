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

func (s *OperationService) Hold(ctx context.Context, req dto.HoldReq) (entity.Operation, *errcode.ErrCode) {
	if err := req.Validate(); err != nil {
		return entity.Operation{}, errcode.New(err.Error())
	}

	if err := s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_OperationHold,
	}); err != nil {
		return entity.Operation{}, err
	}

	var operation entity.Operation
	f := func(tx pgx.Tx) *errcode.ErrCode {
		wallet, err := s.walletRepo.Find(ctx, req.WalletId, tx)
		if err != nil {
			return err
		}

		//Check
		if wallet.Status != dictionary.Get().WalletStatuses.Active.String() {
			return errcode.New("wallet_not_active")
		}
		if req.Amount <= 0 {
			return errcode.New("invalid_operation_amount")
		}
		if wallet.Balance < req.Amount {
			return errcode.New("not_enough_balance")
		}

		//Logic
		transactionId, err := s.transactionRepo.Create(ctx, entity.Transaction{
			ConsumerId:       req.ConsumerId,
			ServiceProvideId: req.ServiceProvideId,
			Type:             dictionary.Get().TransactionTypes.Pay.String(),
			Description:      req.Description,
			OrderId:          req.OrderId,
		}, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		operation, err = s.operationRepo.Create(ctx, entity.Operation{
			TransactionId:      transactionId,
			Type:               dictionary.Get().OperationTypes.Hold.String(),
			WalletId:           req.WalletId,
			Amount:             req.Amount,
			Status:             dictionary.Get().Statuses.Success.String(),
			ErrorCode:          "",
			InternalLogMessage: "",
			BalanceBefore:      wallet.Balance,
			BalanceAfter:       wallet.Balance - req.Amount,
			HoldBalanceBefore:  wallet.Hold,
			HoldBalanceAfter:   wallet.Hold + req.Amount,
		}, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		err = s.walletRepo.AddBalance(ctx, req.WalletId, -req.Amount, req.Amount, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		return nil
	}

	fErr := s.atomicRepo.Do(ctx, f)
	return operation, fErr
}
