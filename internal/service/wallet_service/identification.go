package wallet_service

import (
	"context"
	"fmt"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"github.com/jackc/pgx/v5"
)

func (s *WalletService) Identification(ctx context.Context, req dto.WalletIdentificationReq) (entity.OperationStatus, *errcode.ErrCode) {
	if e := req.Validate(); e != nil {
		return dictionary.Get().Statuses.Error, errcode.New(e.Error())
	}

	if err := s.consumer.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_WalletIdentification,
	}); err != nil {
		return dictionary.Get().Statuses.Error, err
	}

	outStatus := dictionary.Get().Statuses.Error
	f := func(tx pgx.Tx) *errcode.ErrCode {
		// Preload
		wallet, err := s.repo.Find(ctx, req.WalletId, tx)
		if err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("wallet_find err: %s id: %s", err.Log(), req.WalletId))
			return errcode.New("wallet_not_found")
		}
		if wallet.Status != dictionary.Get().WalletStatuses.Active.String() {
			return errcode.New("wallet_not_active")
		}
		if wallet.Identification == req.Identification {
			outStatus = dictionary.Get().Statuses.Success
			return nil
		}

		if err = s.repo.UpdateIdentification(ctx, req.WalletId, req.Identification, tx); err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		outStatus = dictionary.Get().Statuses.Success
		return nil
	}
	fErr := s.atomic.Do(ctx, f)
	return outStatus, fErr
}
