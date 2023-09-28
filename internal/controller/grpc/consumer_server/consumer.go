package consumer_server

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/controller"
	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	ds "git.example.kz/wallet/wallet-back/pkg/proto/wallet_storage_ds"
	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type WalletArmServer struct {
	consumer controller.ConsumerService
	logger   logger.LoggerInterface
	ds.UnimplementedWalletStorageArmServer
}

func NewWalletArmServer(consumer controller.ConsumerService, l logger.LoggerInterface) *WalletArmServer {
	return &WalletArmServer{
		consumer: consumer,
		logger:   l,
	}
}

func (w *WalletArmServer) ConsumerList(c context.Context, empty *empty.Empty) (*ds.ConsumerItems, error) {
	list, err := w.consumer.List(c)
	if err != nil {
		return nil, err
	}
	items := make([]*ds.ConsumerItems_Item, len(list))
	for i, item := range list {
		items[i] = &ds.ConsumerItems_Item{
			Id:               item.ID,
			Code:             item.Code,
			Slug:             item.Slug,
			Secret:           item.Secret,
			WhiteListMethods: item.WhiteListMethods,
			CreatedAt:        timestamppb.New(item.CreatedAt),
		}
	}
	return &ds.ConsumerItems{
		Items: items,
		RequestStatus: &ds.StatusResponse{
			Status:    0,
			ErrorCode: "",
		},
	}, nil
}

func (w *WalletArmServer) ConsumerCreate(c context.Context, req *ds.ConsumerCreateRequest) (*ds.ConsumerCreateResponse, error) {
	consumer, err := w.consumer.Create(c, dto.ConsumerCreateReq{
		Code:             req.GetCode(),
		Slug:             req.GetSlug(),
		Secret:           req.GetSecret(),
		WhiteListMethods: req.GetWhiteListMethods(),
	})
	if err != nil {
		return nil, err
	}

	return &ds.ConsumerCreateResponse{
		Id: consumer.ID,
		RequestStatus: &ds.StatusResponse{
			Status:    0,
			ErrorCode: "",
		},
	}, nil
}

func (w *WalletArmServer) ConsumerRead(c context.Context, id *ds.EntityId) (*ds.ConsumerItems, error) {
	consumer, err := w.consumer.Read(c, id.GetId())
	if err != nil {
		return nil, err
	}

	return &ds.ConsumerItems{
		Items: []*ds.ConsumerItems_Item{
			{
				Id:               consumer.ID,
				Code:             consumer.Code,
				Slug:             consumer.Slug,
				Secret:           consumer.Secret,
				WhiteListMethods: consumer.WhiteListMethods,
				CreatedAt:        timestamppb.New(consumer.CreatedAt),
			},
		},
		RequestStatus: &ds.StatusResponse{
			Status:    0,
			ErrorCode: "",
		},
	}, nil
}

func (w *WalletArmServer) ConsumerUpdate(c context.Context, request *ds.ConsumerUpdateRequest) (*ds.StatusResponse, error) {
	err := w.consumer.Update(c, request.GetId(), dto.ConsumerUpdateReq{
		Code:             request.GetCode(),
		Slug:             request.GetSlug(),
		Secret:           request.GetSecret(),
		WhiteListMethods: request.GetWhiteListMethods(),
	})
	if err != nil {
		return nil, err
	}
	return &ds.StatusResponse{
		Status:    0,
		ErrorCode: "",
	}, nil
}

func (w *WalletArmServer) ConsumerDelete(c context.Context, id *ds.EntityId) (*ds.StatusResponse, error) {
	err := w.consumer.Delete(c, id.GetId())
	if err != nil {
		return nil, err
	}
	return &ds.StatusResponse{
		Status:    0,
		ErrorCode: "",
	}, nil
}

func (w *WalletArmServer) MethodList(c context.Context, empty *empty.Empty) (*ds.MethodListResponse, error) {
	return &ds.MethodListResponse{
		Methods: w.consumer.MethodList(c),
	}, nil
}

func (w *WalletArmServer) mustEmbedUnimplementedWalletStorageArmServer() {
	//TODO implement me
	panic("implement me")
}
