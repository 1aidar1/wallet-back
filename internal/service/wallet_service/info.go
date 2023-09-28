package wallet_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *WalletService) Info(ctx context.Context, req dto.WalletInfoReq) (entity.Wallet, *errcode.ErrCode) {

	if err := s.consumer.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_WalletInfo,
	}); err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Wallet{}, err
	}

	wallet, err := s.repo.Find(ctx, req.WalletId, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Wallet{}, err
	}
	return wallet, nil
}
