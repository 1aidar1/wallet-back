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

func (s *OperationService) Unhold(ctx context.Context, req dto.UnholdReq) (entity.Operation, *errcode.ErrCode) {
	if err := s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_OperationUnhold,
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
		holdOperation := findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Hold, transaction.Operations)
		if (holdOperation == entity.Operation{}) {
			e := errcode.New("no_hold_operation").WithMsg("nu successful hold operation in transaction", req.TransactionId)
			logger.GetProdLogger().Error(e.Log())
			return e
		}
		//check if there is no successful req operation and no successful clear operation in transaction
		if (findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Unhold, transaction.Operations) != entity.Operation{} ||
			findOperation(dictionary.Get().Statuses.Success, dictionary.Get().OperationTypes.Clear, transaction.Operations) != entity.Operation{}) {
			e := errcode.New("operation_already_done").WithMsg("already found successful unhold operation in transaction", req.TransactionId)
			logger.GetProdLogger().Error(e.Log())
			return e
		}

		operation, err = s.operationRepo.Create(ctx, entity.Operation{
			TransactionId:      req.TransactionId,
			Type:               dictionary.Get().OperationTypes.Unhold.String(),
			WalletId:           req.WalletId,
			Amount:             holdOperation.Amount,
			Status:             dictionary.Get().Statuses.Success.String(),
			ErrorCode:          "",
			InternalLogMessage: "",
			BalanceBefore:      wallet.Balance,
			BalanceAfter:       wallet.Balance + holdOperation.Amount,
			HoldBalanceBefore:  wallet.Hold,
			HoldBalanceAfter:   wallet.Hold - holdOperation.Amount,
		}, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		err = s.walletRepo.AddBalance(ctx, req.WalletId, holdOperation.Amount, -holdOperation.Amount, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		return nil
	}

	fErr := s.atomicRepo.Do(ctx, f)
	return operation, fErr
}

// поиск операции по статусу и типу
func findOperation(status entity.OperationStatus, optype entity.OperationType, operations []entity.Operation) entity.Operation {
	for _, operation := range operations {
		if operation.Status == status.String() && operation.Type == optype.String() {
			return operation
		}
	}
	return entity.Operation{}
}
