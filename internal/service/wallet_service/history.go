package wallet_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *WalletService) History(ctx context.Context, req dto.WalletHistoryReq) ([]entity.WalletHistory, *errcode.ErrCode) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := s.consumer.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_WalletHistory,
	}); err != nil {
		return nil, err
	}

	_, err := s.repo.Find(ctx, req.WalletId, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return nil, err
	}
	out, err := s.repo.History(ctx, req, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return nil, err
	}
	for i, _ := range out {
		for j, _ := range out[i].Operations {
			if out[i].Operations[j].WalletId != req.WalletId {
				out[i].Operations[j].BalanceBefore = nil
				out[i].Operations[j].BalanceAfter = nil
				out[i].Operations[j].HoldBalanceBefore = nil
				out[i].Operations[j].HoldBalanceAfter = nil

			}
		}
	}
	return out, err
}
