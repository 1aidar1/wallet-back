package grpc_server

import (
	"fmt"
	"net"

	"git.example.kz/wallet/wallet-back/config"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"git.example.kz/wallet/wallet-back/pkg/proto/wallet_storage"
	ds "git.example.kz/wallet/wallet-back/pkg/proto/wallet_storage_ds"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	GrpcPort            int
	WalletStorageServer wallet_storage.WalletStorageServer
	ArmServer           ds.WalletStorageArmServer
}

func (s *GrpcServer) StartGrpcServer(cfg *config.Config, l logger.LoggerInterface) {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GrpcPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			Recover(),
			// transport_log.ServerInterceptor(transport_log.ParticipantServiceData, l, transport_log.ResponseLimit(40000)),
		),
	)
	wallet_storage.RegisterWalletStorageServer(grpcServer, s.WalletStorageServer)
	ds.RegisterWalletStorageArmServer(grpcServer, s.ArmServer)

	if err = grpcServer.Serve(conn); err != nil {
		panic(err)
	}
}
