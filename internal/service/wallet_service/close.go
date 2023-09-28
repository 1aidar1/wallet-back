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
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *WalletService) Close(ctx context.Context, req dto.WalletCloseReq) (entity.OperationStatus, *errcode.ErrCode) {

	if err := s.consumer.Validate(ctx, dto.ConsumerValidateReq{
		ConsumerId: req.ConsumerId,
		Method:     consumer_service.Method_WalletClose,
	}); err != nil {
		return dictionary.Get().Statuses.Error, err
	}
	outStatus := dictionary.Get().Statuses.Error
	f := func(tx pgx.Tx) *errcode.ErrCode {
		// Preload
		fromWallet, err := s.repo.Find(ctx, req.WalletId, tx)
		if err != nil {
			logger.GetProdLogger().Error(err.Log())
			return errcode.New("wallet_not_found")
		}

		// Check
		//TODO: chto delat esli balance negative
		if fromWallet.Hold != 0 {
			return errcode.New("hold_amount_notzero")
		}
		if fromWallet.Status != dictionary.Get().WalletStatuses.Active.String() {
			return errcode.New("wallet_not_active")
		}

		//Logic
		if fromWallet.Balance > 0 {
			transfer, err := s.operations.TransferWithTX(ctx, dto.TransferReq{
				ConsumerId:       req.ConsumerId,
				ServiceProvideId: dictionary.Get().DefaultServiceProvideId,
				WithdrawWalletId: req.WalletId,
				RefillWalletId:   req.CorrespondingWalletId,
				Amount:           fromWallet.Balance,
				Description:      req.Description,
				OrderId:          uuid.New().String(),
			}, tx)
			if err != nil {
				logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
				return err
			}
			if transfer.Status != dictionary.Get().Statuses.Success.String() {
				logger.GetProdLogger().Error(fmt.Sprintf("transfer status failed"))
				return errcode.NewDefaultErr()
			}
		}
		if err = s.repo.UpdateStatus(ctx, dto.UpdateWalletStatus{
			WalletId:    req.WalletId,
			NewStatus:   dictionary.Get().WalletStatuses.Closed.String(),
			Description: req.Description,
		}, tx); err != nil {
			logger.GetProdLogger().Error(fmt.Sprintf("err: %s", err.Log()))
			return err
		}
		// Return
		outStatus = dictionary.Get().Statuses.Success
		return nil

	}

	fErr := s.atomic.Do(ctx, f)
	return outStatus, fErr

}
