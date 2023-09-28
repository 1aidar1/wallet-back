package wallet_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *WalletService) Statistics(ctx context.Context, req dto.WalletStatisticsReq) (dto.WalletStatisticsRes, *errcode.ErrCode) {
	if e := req.Validate(); e != nil {
		return dto.WalletStatisticsRes{}, errcode.New(e.Error())
	}
	if err := s.consumer.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_WalletStatistics,
	}); err != nil {
		return dto.WalletStatisticsRes{}, err
	}

	_, err := s.repo.Find(ctx, req.WalletId, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return nil, err
	}
	stat, err := s.repo.Stat(ctx, req)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return dto.WalletStatisticsRes{}, err
	}
	return stat, nil
}
