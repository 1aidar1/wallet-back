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

func (s *OperationService) Refill(ctx context.Context, req dto.RefillReq) (entity.Operation, *errcode.ErrCode) {
	if err := req.Validate(); err != nil {
		return entity.Operation{}, errcode.New(err.Error())
	}

	if err := s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_OperationRefill,
	}); err != nil {
		return entity.Operation{}, err
	}
	var operation entity.Operation
	f := func(tx pgx.Tx) *errcode.ErrCode {
		var err *errcode.ErrCode
		operation, err = s.refill(ctx, req, tx)
		if err != nil {
			return err
		}
		return nil
	}

	fErr := s.atomicRepo.Do(ctx, f)
	return operation, fErr
}
func (s *OperationService) refill(ctx context.Context, req dto.RefillReq, tx pgx.Tx) (entity.Operation, *errcode.ErrCode) {
	if e := req.Validate(); e != nil {
		return entity.Operation{}, errcode.New(e.Error())
	}

	wallet, err := s.walletRepo.Find(ctx, req.WalletId, tx)
	if err != nil {
		return entity.Operation{}, err
	}

	//Check
	if wallet.Status != dictionary.Get().WalletStatuses.Active.String() {
		return entity.Operation{}, errcode.New("wallet_not_active")
	}
	if req.Amount <= 0 {
		return entity.Operation{}, errcode.New("invalid_operation_amount")
	}
	//если transaction id выставлен, то добавляем операцию к существующей транзакции (сначала проверяем, что она существует)
	//иначе создаем новую транзакцию
	var transactionId string
	if req.TransactionId == nil {
		transactionId, err = s.transactionRepo.Create(ctx, entity.Transaction{
			ConsumerId:       req.ConsumerId,
			ServiceProvideId: req.ServiceProvideId,
			Type:             dictionary.Get().TransactionTypes.Pay.String(),
			Description:      req.Description,
			OrderId:          req.OrderId,
		}, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return entity.Operation{}, err
		}
	} else {
		trx, err := s.transactionRepo.Find(ctx, *req.TransactionId, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return entity.Operation{}, err
		}
		transactionId = trx.Id
	}
	operation, err := s.operationRepo.Create(ctx, entity.Operation{
		TransactionId:      transactionId,
		Type:               dictionary.Get().OperationTypes.Refill.String(),
		WalletId:           req.WalletId,
		Amount:             req.Amount,
		Status:             dictionary.Get().Statuses.Success.String(),
		ErrorCode:          "",
		InternalLogMessage: "",
		BalanceBefore:      wallet.Balance,
		BalanceAfter:       wallet.Balance + req.Amount,
		HoldBalanceBefore:  wallet.Hold,
		HoldBalanceAfter:   wallet.Hold,
	}, tx)
	if err != nil {
		logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
		return entity.Operation{}, err
	}
	err = s.walletRepo.AddBalance(ctx, req.WalletId, req.Amount, 0, tx)
	if err != nil {
		logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
		return entity.Operation{}, err
	}
	return operation, nil
}
