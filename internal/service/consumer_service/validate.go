package consumer_service

import (
	"context"
	"fmt"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"git.example.kz/wallet/wallet-back/pkg/utils"
)

func (s *ConsumerService) Validate(ctx context.Context, req dto.ConsumerValidateReq) *errcode.ErrCode {
	// check if method allowed
	logger.GetProdLogger().Info(fmt.Sprintf("cons validate %+v\n", req))
	consumer, err := s.repo.Read(ctx, req.ConsumerId, nil)
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return err
	}
	if utils.Contains(consumer.WhiteListMethods, req.Method) {
		return nil
	} else {
		return errcode.New("method_not_allowed")
	}
}
