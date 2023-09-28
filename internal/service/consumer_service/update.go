package consumer_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *ConsumerService) Update(ctx context.Context, consumerId string, req dto.ConsumerUpdateReq) *errcode.ErrCode {
	err := s.repo.Update(ctx, consumerId, req, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return err
	}
	return nil
}
