package consumer_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *ConsumerService) Delete(ctx context.Context, id string) *errcode.ErrCode {
	err := s.repo.Delete(ctx, id, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return err
	}
	return nil
}
