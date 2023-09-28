package consumer_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *ConsumerService) List(ctx context.Context) ([]entity.Consumer, *errcode.ErrCode) {
	list, err := s.repo.List(ctx, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return nil, err
	}
	return list, nil
}
