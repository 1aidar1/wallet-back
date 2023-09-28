package consumer_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *ConsumerService) Read(ctx context.Context, id string) (entity.Consumer, *errcode.ErrCode) {
	consumer, err := s.repo.Read(ctx, id, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Consumer{}, err
	}
	return consumer, nil
}
