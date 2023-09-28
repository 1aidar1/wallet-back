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

func (s *OperationService) Clear(ctx context.Context, req dto.ClearReq) (entity.Operation, *errcode.ErrCode) {
	if err := req.Validate(); err != nil {
		return entity.Operation{}, errcode.New(err.Error())
	}
	if err := s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_OperationClear,
	}); err != nil {
		return entity.Operation{}, err
	}

	var operation entity.Operation
	f := func(tx pgx.Tx) *errcode.ErrCode {
		wallet, err := s.walletRepo.Find(ctx, req.WalletId, tx)
		if err != nil {
			logger.GetProdLogger().Error(err.Log())
			return err
		}

		//Check
		if wallet.Status != dictionary.Get().WalletStatuses.Active.String() {
			return errcode.New("wallet_not_active")
		}

		transaction, err := s.transactionRepo.Find(ctx, req.TransactionId, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		// проверяем есть ли операция hold, чтобы сделать clear по ней.
		holdOperation := findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Hold, transaction.Operations)
		if (holdOperation == entity.Operation{}) {
			e := errcode.New("no_hold_operation").WithMsg("nu successful hold operation in transaction", req.TransactionId)
			logger.GetProdLogger().Error(e.Log())
			return e
		}
		// проверяем небыло ли успешных unhold или clear на наш hold
		if (findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Clear, transaction.Operations) != entity.Operation{} ||
			findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Unhold, transaction.Operations) != entity.Operation{}) {
			e := errcode.New("operation_already_done").WithMsg("already found successful clear operation in transaction", req.TransactionId)
			logger.GetProdLogger().Error(e.Log())
			return e
		}

		operation, err = s.operationRepo.Create(ctx, entity.Operation{
			TransactionId:      req.TransactionId,
			Type:               dictionary.Get().OperationTypes.Clear.String(),
			WalletId:           req.WalletId,
			Amount:             holdOperation.Amount,
			Status:             dictionary.Get().Statuses.Success.String(),
			ErrorCode:          "",
			InternalLogMessage: "",
			BalanceBefore:      wallet.Balance,
			BalanceAfter:       wallet.Balance,
			HoldBalanceBefore:  wallet.Hold,
			HoldBalanceAfter:   wallet.Hold - holdOperation.Amount,
		}, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		err = s.walletRepo.AddBalance(ctx, req.WalletId, 0, -holdOperation.Amount, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		return nil
	}

	fErr := s.atomicRepo.Do(ctx, f)
	return operation, fErr
}
