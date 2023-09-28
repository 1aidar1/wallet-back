package wallet_storage_server

import (
	"context"
	"time"

	"git.example.kz/wallet/wallet-back/internal/controller"
	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"git.example.kz/wallet/wallet-back/pkg/proto/wallet_storage"
	"git.example.kz/wallet/wallet-back/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WalletStorageServer struct {
	wallet    controller.WalletService
	operation controller.OperationService
	logger    logger.LoggerInterface
	wallet_storage.UnimplementedWalletStorageServer
}

var _ wallet_storage.WalletStorageServer = &WalletStorageServer{}

func NewWalletStorageServer(wallet controller.WalletService, operation controller.OperationService, l logger.LoggerInterface) *WalletStorageServer {
	return &WalletStorageServer{
		wallet:    wallet,
		operation: operation,
		logger:    l,
	}
}

func (h *WalletStorageServer) Create(c context.Context, req *wallet_storage.CreateRequest) (*wallet_storage.CreateResponse, error) {
	wallet, err := h.wallet.Create(c, dto.WalletCreateReq{
		ConsumerId:   req.GetConsumerId(),
		CurrencyCode: req.GetCurrency(),
		CountryCode:  req.GetCountry(),
		Phone:        req.GetAdditionalIdentifiers().GetPhone(),
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.CreateResponse{
		Identifiers: &wallet_storage.WalletIdentifierResponse{
			Id:    wallet.ID,
			Phone: wallet.Phone,
		},
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(dictionary.Get().Statuses.Success),
			ErrorCode: "",
		},
	}, nil
}

func (h *WalletStorageServer) Info(c context.Context, req *wallet_storage.InfoRequest) (*wallet_storage.InfoResponse, error) {
	walletId, err := h.wallet.GetId(c, identifierRequestToEntity(req.GetWalletIdentifier()))
	if err != nil {
		return nil, err
	}
	wallet, err := h.wallet.Info(c, dto.WalletInfoReq{
		ConsumerId: req.GetConsumerId(),
		WalletId:   walletId,
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.InfoResponse{
		Identifiers: &wallet_storage.WalletIdentifierResponse{
			Id:    wallet.ID,
			Phone: wallet.Phone,
		},
		Currency:             dictionary.Get().Currencies.ById(wallet.CurrencyID).Code,
		Country:              dictionary.Get().Countries.ById(wallet.CountryID).Code,
		WalletStatus:         internalWalletStatusToGrpcWalletStatus(entity.WalletStatus(wallet.Status)),
		Balance:              utils.BusinessToRest(wallet.Balance),
		HoldBalance:          utils.BusinessToRest(wallet.Hold),
		IdentificationStatus: internalIdentificationToGrpcIdentification(entity.IdentificationType(wallet.Identification)),
		CreatedAt:            timestamppb.New(*wallet.CreatedAt),
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(dictionary.Get().Statuses.Success),
			ErrorCode: "",
		},
	}, nil
}

func (h *WalletStorageServer) History(c context.Context, req *wallet_storage.HistoryRequest) (*wallet_storage.HistoryResponse, error) {
	walletId, e := h.wallet.GetId(c, identifierRequestToEntity(req.GetIdentifiers()))
	if e != nil {
		return nil, e
	}
	history, e := h.wallet.History(c, dto.WalletHistoryReq{
		ConsumerId: req.GetConsumerId(),
		WalletId:   walletId,
		DateStart:  req.GetDateStart().AsTime(),
		DateEnd:    req.GetDateEnd().AsTime(),
		Page:       int(req.GetPaging().GetPage()),
		PerPage:    int(req.GetPaging().GetItemsNumberPerPage()),
	})
	if e != nil {
		return nil, e
	}
	transactions := make([]*wallet_storage.Transaction, len(history))
	for i, t := range history {
		operations := make([]*wallet_storage.Transaction_Operation, len(t.Operations))
		for j, operation := range history[i].Operations {
			operations[j] = &wallet_storage.Transaction_Operation{
				Id:                operation.Id,
				Type:              internalOperationTypeToGrpcOperationType(entity.OperationType(operation.Type)),
				WalletId:          operation.WalletId,
				Amount:            utils.BusinessToRest(operation.Amount),
				Status:            internalStatusToGrpcStatus(entity.OperationStatus(operation.Status)),
				ErrorCode:         operation.ErrorCode,
				BalanceBefore:     utils.BusinessToRestP(operation.BalanceBefore),
				BalanceAfter:      utils.BusinessToRestP(operation.BalanceAfter),
				HoldBalanceBefore: utils.BusinessToRestP(operation.HoldBalanceBefore),
				HoldBalanceAfter:  utils.BusinessToRestP(operation.HoldBalanceAfter),
				CreatedAt:         timestamppb.New(operation.CreatedAt),
			}
		}
		transactions[i] = &wallet_storage.Transaction{
			Id:                t.Id,
			ConsumerId:        t.ConsumerId,
			Type:              internalTransactionTypeToGrpcTransactionType(entity.TransactionType(t.Type)),
			Description:       t.Description,
			ServiceProviderId: t.ServiceProvideId,
			OrderId:           t.OrderId,
			CreatedAt:         timestamppb.New(t.CreatedAt),
			Operations:        operations,
		}
	}

	return &wallet_storage.HistoryResponse{
		TotalItemsCount: int64(len(transactions)),
		Transactions:    transactions,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(dictionary.Get().Statuses.Success),
			ErrorCode: "",
		},
	}, nil
}

func (h *WalletStorageServer) HistoryByProvider(c context.Context, req *wallet_storage.HistoryByProviderRequest) (*wallet_storage.HistoryResponse, error) {
	walletId, e := h.wallet.GetId(c, identifierRequestToEntity(req.GetIdentifiers()))
	if e != nil {
		return nil, e
	}
	history, e := h.wallet.HistoryByProvider(c, dto.WalletHistoryReq{
		ConsumerId: req.GetConsumerId(),
		WalletId:   walletId,
		ProviderId: &req.ProviderId,
		DateStart:  req.GetDateStart().AsTime(),
		DateEnd:    req.GetDateEnd().AsTime(),
		PerPage:    int(req.GetPaging().GetItemsNumberPerPage()),
		Page:       int(req.GetPaging().GetPage()),
	})
	if e != nil {
		return nil, e
	}
	transactions := make([]*wallet_storage.Transaction, len(history))
	for i, t := range history {
		operations := make([]*wallet_storage.Transaction_Operation, len(t.Operations))
		for j, operation := range history[i].Operations {
			operations[j] = &wallet_storage.Transaction_Operation{
				Id:                operation.Id,
				Type:              internalOperationTypeToGrpcOperationType(entity.OperationType(operation.Type)),
				WalletId:          operation.WalletId,
				Amount:            utils.BusinessToRest(operation.Amount),
				Status:            internalStatusToGrpcStatus(entity.OperationStatus(operation.Status)),
				ErrorCode:         operation.ErrorCode,
				BalanceBefore:     utils.BusinessToRestP(operation.BalanceBefore),
				BalanceAfter:      utils.BusinessToRestP(operation.BalanceAfter),
				HoldBalanceBefore: utils.BusinessToRestP(operation.HoldBalanceBefore),
				HoldBalanceAfter:  utils.BusinessToRestP(operation.HoldBalanceAfter),
				CreatedAt:         timestamppb.New(operation.CreatedAt),
			}
		}
		transactions[i] = &wallet_storage.Transaction{
			Id:                t.Id,
			ConsumerId:        t.ConsumerId,
			Type:              internalTransactionTypeToGrpcTransactionType(entity.TransactionType(t.Type)),
			Description:       t.Description,
			ServiceProviderId: t.ServiceProvideId,
			OrderId:           t.OrderId,
			CreatedAt:         timestamppb.New(t.CreatedAt),
			Operations:        operations,
		}
	}

	return &wallet_storage.HistoryResponse{
		TotalItemsCount: int64(len(transactions)),
		Transactions:    transactions,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(dictionary.Get().Statuses.Success),
			ErrorCode: "",
		},
	}, nil
}

func (h *WalletStorageServer) Statistics(c context.Context, req *wallet_storage.StatisticsRequest) (*wallet_storage.StatisticsResponse, error) {
	walletId, e := h.wallet.GetId(c, identifierRequestToEntity(req.GetWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	wallet, e := h.wallet.Statistics(c, dto.WalletStatisticsReq{
		ConsumerId: req.GetConsumerId(),
		WalletId:   walletId,
		DateStart:  req.GetPeriod().GetFrom().AsTime(),
		DateEnd:    req.GetPeriod().GetTo().AsTime(),
		Status:     dictionary.Get().Statuses.Success,
	})
	if e != nil {
		return nil, e
	}

	var refill, withdraw *wallet_storage.StatisticsResponse_Statistics
	for key, _ := range wallet {
		switch key {
		case dictionary.Get().OperationTypes.Refill.String():
			refill = &wallet_storage.StatisticsResponse_Statistics{
				PaymentsAmount: utils.BusinessToRest(wallet[key].Amount),
				PaymentsCount:  int64(wallet[key].Count),
			}
		case dictionary.Get().OperationTypes.Withdraw.String():
			withdraw = &wallet_storage.StatisticsResponse_Statistics{
				PaymentsAmount: utils.BusinessToRest(wallet[key].Amount),
				PaymentsCount:  int64(wallet[key].Count),
			}
		}
	}
	out := &wallet_storage.StatisticsResponse{
		Refill:   refill,
		Withdraw: withdraw,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(dictionary.Get().Statuses.Success),
			ErrorCode: "",
		},
	}
	return out, nil
}

func (h *WalletStorageServer) Close(c context.Context, req *wallet_storage.CloseRequest) (*wallet_storage.StatusResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, err := h.getWalletIdAndLock(
		ctx,
		identifierRequestToEntity(req.GetWalletIdentifier()),
		identifierRequestToEntity(req.GetCorrespondingWalletIdentifier()),
	)
	if err != nil {
		return nil, err
	}
	defer h.wallet.Unlock(c, walletId...)

	status, err := h.wallet.Close(c, dto.WalletCloseReq{
		ConsumerId:            req.ConsumerId,
		WalletId:              walletId[0],
		CorrespondingWalletId: walletId[1],
		Description:           req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.StatusResponse{
		Status:    internalStatusToGrpcStatus(status),
		ErrorCode: "",
	}, nil
}

func (h *WalletStorageServer) Block(c context.Context, req *wallet_storage.BlockRequest) (*wallet_storage.StatusResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, err := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if err != nil {
		return nil, err
	}
	defer h.wallet.Unlock(c, walletId...)

	status, err := h.wallet.Block(c, dto.WalletBlockReq{
		ConsumerId:  req.GetConsumerId(),
		WalletId:    walletId[0],
		Description: req.GetConsumerId(),
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.StatusResponse{
		Status:    internalStatusToGrpcStatus(status),
		ErrorCode: "",
	}, nil
}

func (h *WalletStorageServer) Unblock(c context.Context, req *wallet_storage.UnblockRequest) (*wallet_storage.StatusResponse, error) {

	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, err := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if err != nil {
		return nil, err
	}
	defer h.wallet.Unlock(c, walletId...)

	status, err := h.wallet.Unblock(c, dto.WalletUnblockReq{
		ConsumerId:  req.GetConsumerId(),
		WalletId:    walletId[0],
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.StatusResponse{
		Status:    internalStatusToGrpcStatus(status),
		ErrorCode: "",
	}, nil
}

func (h *WalletStorageServer) Identification(c context.Context, req *wallet_storage.IdentificationRequest) (*wallet_storage.StatusResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()

	walletId, err := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if err != nil {
		return nil, err
	}
	defer h.wallet.Unlock(c, walletId...)

	status, err := h.wallet.Identification(c, dto.WalletIdentificationReq{
		ConsumerId:     req.GetConsumerId(),
		WalletId:       walletId[0],
		Identification: grpcIdentificationToInternal(req.GetStatus()).String(),
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.StatusResponse{
		Status:    internalStatusToGrpcStatus(status),
		ErrorCode: "",
	}, nil
}

func (h *WalletStorageServer) Refill(c context.Context, req *wallet_storage.RefillRequest) (*wallet_storage.RefillResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, err := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if err != nil {
		return nil, err
	}
	defer h.wallet.Unlock(c, walletId...)

	refill, err := h.operation.Refill(ctx, dto.RefillReq{
		ServiceProvideId: req.GetServiceProviderId(),
		ConsumerId:       req.ConsumerId,
		WalletId:         walletId[0],
		Amount:           utils.RestToBusiness(req.Amount),
		Description:      req.Description,
		OrderId:          req.OrderId,
	})
	if err != nil {
		return nil, err
	}
	return &wallet_storage.RefillResponse{
		TransactionId: refill.TransactionId,
		OperationId:   refill.Id,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(refill.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*refill.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) Withdraw(c context.Context, req *wallet_storage.WithdrawRequest) (*wallet_storage.WithdrawResponse, error) {

	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, e := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	defer h.wallet.Unlock(c, walletId...)

	withdraw, e := h.operation.Withdraw(ctx, dto.WithdrawReq{
		ConsumerId:       req.GetConsumerId(),
		ServiceProvideId: req.GetServiceProviderId(),
		WalletId:         walletId[0],
		Amount:           utils.RestToBusiness(req.GetAmount()),
		Description:      req.GetDescription(),
		OrderId:          req.GetOrderId(),
		TransactionId:    nil,
	})
	if e != nil {
		return nil, e
	}
	return &wallet_storage.WithdrawResponse{
		TransactionId: withdraw.TransactionId,
		OperationId:   withdraw.Id,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(withdraw.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*withdraw.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) Hold(c context.Context, req *wallet_storage.HoldRequest) (*wallet_storage.HoldResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()

	walletId, e := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	defer h.wallet.Unlock(c, walletId...)

	hold, e := h.operation.Hold(ctx, dto.HoldReq{
		ConsumerId:       req.GetConsumerId(),
		ServiceProvideId: req.GetServiceProviderId(),
		WalletId:         walletId[0],
		Amount:           utils.RestToBusiness(req.GetAmount()),
		Description:      req.GetDescription(),
		OrderId:          req.GetOrderId(),
	})
	if e != nil {
		return nil, e
	}
	return &wallet_storage.HoldResponse{
		TransactionId: hold.TransactionId,
		OperationId:   hold.Id,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(hold.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*hold.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) Unhold(c context.Context, req *wallet_storage.UnholdRequest) (*wallet_storage.UnholdResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()

	walletId, e := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	defer h.wallet.Unlock(c, walletId...)

	unhold, e := h.operation.Unhold(ctx, dto.UnholdReq{
		ConsumerId:    req.GetConsumerId(),
		WalletId:      walletId[0],
		TransactionId: req.GetTransactionId(),
	})
	if e != nil {
		return nil, e
	}
	return &wallet_storage.UnholdResponse{
		TransactionId: unhold.TransactionId,
		OperationId:   unhold.Id,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(unhold.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*unhold.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) Clear(c context.Context, req *wallet_storage.ClearRequest) (*wallet_storage.ClearResponse, error) {

	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, e := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	defer h.wallet.Unlock(c, walletId...)

	clear, e := h.operation.Clear(ctx, dto.ClearReq{
		ConsumerId:    req.GetConsumerId(),
		TransactionId: req.GetTransactionId(),
		WalletId:      walletId[0],
		Amount:        utils.RestToBusiness(req.GetAmount()),
	})
	if e != nil {
		return nil, e
	}
	return &wallet_storage.ClearResponse{
		OperationId: clear.Id,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(clear.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*clear.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) Refund(c context.Context, req *wallet_storage.RefundRequest) (*wallet_storage.RefundResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, e := h.getWalletIdAndLock(ctx, identifierRequestToEntity(req.GetWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	defer h.wallet.Unlock(c, walletId...)

	refund, e := h.operation.Refund(ctx, dto.RefundReq{
		ConsumerId:    req.GetConsumerId(),
		TransactionId: req.GetTransactionId(),
		WalletId:      walletId[0],
	})
	if e != nil {
		return nil, e
	}
	return &wallet_storage.RefundResponse{
		TransactionId: refund.TransactionId,
		OperationId:   refund.Id,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(refund.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*refund.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) Transfer(c context.Context, req *wallet_storage.TransferRequest) (*wallet_storage.TransferResponse, error) {
	ctx, cancel := context.WithTimeout(c, controller.MAX_LIFE_SECONDS*time.Second)
	defer cancel()
	walletId, e := h.getWalletIdAndLock(ctx,
		identifierRequestToEntity(req.GetWithdrawWalletIdentifier()),
		identifierRequestToEntity(req.GetRefillWalletIdentifier()))
	if e != nil {
		return nil, e
	}
	defer h.wallet.Unlock(c, walletId...)

	transfer, e := h.operation.Transfer(ctx, dto.TransferReq{
		ConsumerId:       req.GetConsumerId(),
		ServiceProvideId: req.GetServiceProviderId(),
		WithdrawWalletId: walletId[0],
		RefillWalletId:   walletId[1],
		Amount:           utils.RestToBusiness(req.GetAmount()),
		Description:      req.GetDescription(),
		OrderId:          req.GetOrderId(),
	})
	if e != nil {
		return nil, e
	}
	return &wallet_storage.TransferResponse{
		TransactionId: transfer.TransactionId,
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(entity.OperationStatus(transfer.Status)),
			ErrorCode: "",
		},
		CreatedAt: timestamppb.New(*transfer.CreatedAt),
	}, nil
}

func (h *WalletStorageServer) TransactionInfo(c context.Context, req *wallet_storage.TransactionInfoRequest) (*wallet_storage.TransactionInfoResponse, error) {
	transaction, e := h.operation.FindTransaction(c, dto.FindTransactionReq{
		ConsumerId:       req.GetConsumerId(),
		TransactionId:    req.GetTransactionId(),
		OrderId:          req.GetOrder().GetOrderId(),
		ServiceProvideId: req.GetOrder().GetServiceProviderId(),
	})
	if e != nil {
		return nil, e
	}

	operations := make([]*wallet_storage.Transaction_Operation, len(transaction.Operations))
	for i, op := range transaction.Operations {
		balanceBefore := utils.BusinessToRest(op.BalanceBefore)
		balanceAfter := utils.BusinessToRest(op.BalanceAfter)
		holdBefore := utils.BusinessToRest(op.HoldBalanceBefore)
		holdAfter := utils.BusinessToRest(op.HoldBalanceAfter)
		operations[i] = &wallet_storage.Transaction_Operation{
			Id:                op.Id,
			Type:              internalOperationTypeToGrpcOperationType(entity.OperationType(op.Type)),
			WalletId:          op.WalletId,
			Amount:            utils.BusinessToRest(op.Amount),
			Status:            internalStatusToGrpcStatus(entity.OperationStatus(op.Status)),
			ErrorCode:         op.ErrorCode,
			BalanceBefore:     &balanceBefore,
			BalanceAfter:      &balanceAfter,
			HoldBalanceBefore: &holdBefore,
			HoldBalanceAfter:  &holdAfter,
			CreatedAt:         timestamppb.New(*op.CreatedAt),
		}
	}
	return &wallet_storage.TransactionInfoResponse{
		RequestStatus: &wallet_storage.StatusResponse{
			Status:    internalStatusToGrpcStatus(dictionary.Get().Statuses.Success),
			ErrorCode: "",
		}, Transaction: &wallet_storage.Transaction{
			Id:                transaction.Id,
			ConsumerId:        transaction.ConsumerId,
			Type:              internalTransactionTypeToGrpcTransactionType(entity.TransactionType(transaction.Type)),
			Description:       transaction.Description,
			OrderId:           transaction.OrderId,
			ServiceProviderId: transaction.ServiceProvideId,
			CreatedAt:         timestamppb.New(transaction.CreatedAt),
			Operations:        operations,
		},
	}, nil
}

func (h *WalletStorageServer) mustEmbedUnimplementedWalletStorageServer() {
	//TODO implement me
	panic("implement me")
}

func identifierRequestToEntity(req *wallet_storage.WalletIdentifierRequest) entity.WalletIdentifierType {
	return entity.WalletIdentifierType{
		Id:    req.GetId(),
		Phone: req.GetPhone(),
	}
}

func (h *WalletStorageServer) getWalletIdAndLock(ctx context.Context, data ...entity.WalletIdentifierType) ([]string, *errcode.ErrCode) {
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

func internalStatusToGrpcStatus(status entity.OperationStatus) wallet_storage.OperationStatus {
	var ans wallet_storage.OperationStatus
	switch status {
	case dictionary.Get().Statuses.Success:
		ans = wallet_storage.OperationStatus_SUCCESS
	case dictionary.Get().Statuses.Error:
		ans = wallet_storage.OperationStatus_ERROR
	case dictionary.Get().Statuses.Process:
		ans = wallet_storage.OperationStatus_PROCESS
	}
	return ans
}

func internalWalletStatusToGrpcWalletStatus(status entity.WalletStatus) wallet_storage.WalletStatus {
	var ans wallet_storage.WalletStatus
	switch status {
	case dictionary.Get().WalletStatuses.Active:
		ans = wallet_storage.WalletStatus_WALLET_STATUS_ACTIVE
	case dictionary.Get().WalletStatuses.Closed:
		ans = wallet_storage.WalletStatus_WALLET_STATUS_CLOSED
	case dictionary.Get().WalletStatuses.Blocked:
		ans = wallet_storage.WalletStatus_WALLET_STATUS_BLOCKED
	}
	return ans
}

func internalIdentificationToGrpcIdentification(status entity.IdentificationType) wallet_storage.WalletIdentificationStatus {
	var ans wallet_storage.WalletIdentificationStatus
	switch status {
	case dictionary.Get().WalletIdentification.None:
		ans = wallet_storage.WalletIdentificationStatus_WALLET_IDENTIFICATION_NONE
	case dictionary.Get().WalletIdentification.Full:
		ans = wallet_storage.WalletIdentificationStatus_WALLET_IDENTIFICATION_FULL
	case dictionary.Get().WalletIdentification.Basic:
		ans = wallet_storage.WalletIdentificationStatus_WALLET_IDENTIFICATION_BASIC
	}
	return ans
}
func grpcIdentificationToInternal(status wallet_storage.WalletIdentificationStatus) entity.IdentificationType {
	var ans entity.IdentificationType
	switch status {
	case wallet_storage.WalletIdentificationStatus_WALLET_IDENTIFICATION_BASIC:
		ans = dictionary.Get().WalletIdentification.Basic
	case wallet_storage.WalletIdentificationStatus_WALLET_IDENTIFICATION_FULL:
		ans = dictionary.Get().WalletIdentification.Full
	case wallet_storage.WalletIdentificationStatus_WALLET_IDENTIFICATION_NONE:
		ans = dictionary.Get().WalletIdentification.None
	}
	return ans
}

func internalTransactionTypeToGrpcTransactionType(t entity.TransactionType) wallet_storage.TransactionType {
	var ans wallet_storage.TransactionType
	switch t {
	case dictionary.Get().TransactionTypes.Transfer:
		ans = wallet_storage.TransactionType_TRANSACTION_TYPE_TRANSFER
	case dictionary.Get().TransactionTypes.Pay:
		ans = wallet_storage.TransactionType_TRANSACTION_TYPE_PAY
	}
	return ans
}

func internalOperationTypeToGrpcOperationType(t entity.OperationType) wallet_storage.OperationType {
	var ans wallet_storage.OperationType
	switch t {
	case dictionary.Get().OperationTypes.Hold:
		ans = wallet_storage.OperationType_OPERATION_TYPE_HOLD
	case dictionary.Get().OperationTypes.Unhold:
		ans = wallet_storage.OperationType_OPERATION_TYPE_UNHOLD
	case dictionary.Get().OperationTypes.Refill:
		ans = wallet_storage.OperationType_OPERATION_TYPE_REFILL
	case dictionary.Get().OperationTypes.Withdraw:
		ans = wallet_storage.OperationType_OPERATION_TYPE_WITHDRAW
	case dictionary.Get().OperationTypes.Clear:
		ans = wallet_storage.OperationType_OPERATION_TYPE_CLEAR
	case dictionary.Get().OperationTypes.Refund:
		ans = wallet_storage.OperationType_OPERATION_TYPE_REFUND
	}
	return ans
}
