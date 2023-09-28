package operation_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *OperationService) FindTransaction(ctx context.Context, req dto.FindTransactionReq) (entity.Transaction, *errcode.ErrCode) {

	if err := s.consumerService.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_TransactionInfo,
	}); err != nil {
		return entity.Transaction{}, err
	}

	if req.TransactionId != "" {
		transaction, err := s.transactionRepo.Find(ctx, req.TransactionId, nil)
		if err != nil {
			logger.GetProdLogger().Error(err.Log())
			return entity.Transaction{}, err
		}
		return transaction, nil
	} else if req.OrderId != "" && req.ServiceProvideId != "" {
		transaction, err := s.transactionRepo.FindByOrder(ctx, req.OrderId, req.ServiceProvideId, nil)
		if err != nil {
			logger.GetProdLogger().Error(err.Log())
			return entity.Transaction{}, err
		}
		return transaction, nil
	} else {
		return entity.Transaction{}, errcode.New("transaction_not_found")
	}
}
