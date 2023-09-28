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

func (s *OperationService) Refund(ctx context.Context, req dto.RefundReq) (entity.Operation, *errcode.ErrCode) {
	if err := s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_OperationRefund,
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
		// проверяем есть ли операция withdraw или clear, на которую можно совершить возврат.
		var refundableOperation entity.Operation
		{
			withdrawOperation := findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Withdraw, transaction.Operations)
			clearingOperation := findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Clear, transaction.Operations)
			if (clearingOperation != entity.Operation{} && withdrawOperation != entity.Operation{}) {
				e := errcode.New("operation_already_done").WithMsg("2 операции списания внутри одной транзакции.")
				logger.GetProdLogger().Error(fmt.Sprintf("err: %s", e.Log()))
				return e
			}
			if (clearingOperation != entity.Operation{}) {
				refundableOperation = clearingOperation
			}
			if (withdrawOperation != entity.Operation{}) {
				refundableOperation = withdrawOperation
			}
		}

		if (refundableOperation == entity.Operation{}) {
			e := errcode.New("nothing_to_refund").WithMsg(req.TransactionId)
			logger.GetProdLogger().Error(e.Log())
			return e
		}
		// проверяем, есть ли успешный возврат.
		if (findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Refund, transaction.Operations) != entity.Operation{}) {
			e := errcode.New("operation_already_done").WithMsg("already found successful refund operation in transaction", req.TransactionId)
			logger.GetProdLogger().Error(e.Log())
			return e
		}

		operation, err = s.operationRepo.Create(ctx, entity.Operation{
			TransactionId:      req.TransactionId,
			Type:               dictionary.Get().OperationTypes.Refund.String(),
			WalletId:           req.WalletId,
			Amount:             refundableOperation.Amount,
			Status:             dictionary.Get().Statuses.Success.String(),
			ErrorCode:          "",
			InternalLogMessage: "",
			BalanceBefore:      wallet.Balance,
			BalanceAfter:       wallet.Balance + refundableOperation.Amount,
			HoldBalanceBefore:  wallet.Hold,
			HoldBalanceAfter:   wallet.Hold,
		}, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		err = s.walletRepo.AddBalance(ctx, req.WalletId, refundableOperation.Amount, 0, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		return nil
	}

	fErr := s.atomicRepo.Do(ctx, f)
	return operation, fErr
}
