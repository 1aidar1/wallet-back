package http

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
)

type DefaultResponse struct {
	Status    string `json:"status"`
	ErrorCode string `json:"error_code"`
}

func defaultHttpError(code string) DefaultResponse {
	return DefaultResponse{
		Status:    dictionary.Get().Statuses.Error.String(),
		ErrorCode: code,
	}
}

func (h *Handler) getWalletIdAndLock(ctx context.Context, data ...entity.WalletIdentifierType) ([]string, *errcode.ErrCode) {
	walletIds := make([]string, len(data))
	for i, d := range data {
		walletId, e := h.wallet.GetId(ctx, d)
		if e != nil {
			return nil, e
		}
		walletIds[i] = walletId
	}
	e := h.wallet.WaitAndLock(ctx, walletIds...)
	if e != nil {
		return nil, e
	}
	return walletIds, nil
}
