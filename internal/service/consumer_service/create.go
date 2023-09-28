package consumer_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *ConsumerService) Create(ctx context.Context, req dto.ConsumerCreateReq) (entity.Consumer, *errcode.ErrCode) {
	consumer, err := s.repo.Create(ctx, req, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Consumer{}, err
	}
	return consumer, nil
}
