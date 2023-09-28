package wallet_service

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
)

func (s *WalletService) Create(ctx context.Context, req dto.WalletCreateReq) (entity.Wallet, *errcode.ErrCode) {
	if err := req.Validate(); err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Wallet{}, err
	}
	if err := s.consumer.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_WalletCreate,
	}); err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Wallet{}, err
	}

	wallet, err := s.repo.Create(ctx, entity.Wallet{
		CurrencyID:     dictionary.Get().Currencies.ByCode(req.CurrencyCode).ID,
		CountryID:      dictionary.Get().Countries.ByCode(req.CountryCode).ID,
		Balance:        0,
		Phone:          req.Phone,
		Hold:           0,
		Status:         dictionary.Get().WalletStatuses.Active.String(),
		Identification: dictionary.Get().WalletIdentification.None.String(),
	})
	if err != nil {
		logger.GetProdLogger().Error(err.Log())
		return entity.Wallet{}, err
	}
	return wallet, nil
}
