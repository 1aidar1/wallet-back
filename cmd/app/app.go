package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"git.example.kz/wallet/wallet-back/config"
	grpc_server "git.example.kz/wallet/wallet-back/internal/controller/grpc"
	"git.example.kz/wallet/wallet-back/internal/controller/grpc/consumer_server"
	"git.example.kz/wallet/wallet-back/internal/controller/grpc/wallet_storage_server"
	delivery "git.example.kz/wallet/wallet-back/internal/controller/http"
	"git.example.kz/wallet/wallet-back/internal/dictionary"
	"git.example.kz/wallet/wallet-back/internal/infra/db"
	"git.example.kz/wallet/wallet-back/internal/infra/redis"
	"git.example.kz/wallet/wallet-back/internal/repository/atomic_repo"
	"git.example.kz/wallet/wallet-back/internal/repository/consumer_repo"
	"git.example.kz/wallet/wallet-back/internal/repository/operation_repo"
	"git.example.kz/wallet/wallet-back/internal/repository/transaction_repo"
	"git.example.kz/wallet/wallet-back/internal/repository/wallet_repo"
	"git.example.kz/wallet/wallet-back/internal/service/consumer_service"
	"git.example.kz/wallet/wallet-back/internal/service/operation_service"
	"git.example.kz/wallet/wallet-back/internal/service/wallet_service"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "app",
	Run: app,
}

func app(cmd *cobra.Command, args []string) {
	cfg := config.GetGeneralConfig()
	_logger := logger.GetProdLogger()

	_db := db.NewDB(cfg)
	_redis := redis.NewRedis(cfg)

	dictionary.CreateOrUpdate(_db)
	//fmt.Printf("%+v\n", _dictionary)

	walletRepo := wallet_repo.NewWalletRepo(_db)
	operationRepo := operation_repo.NewOperationRepo(_db)
	transactionRepo := transaction_repo.NewTransactionRepo(_db)
	consumerRepo := consumer_repo.NewConsumerRepo(_db)
	txRepo := atomic_repo.NewAtomicRepo(_db)

	consumerService := consumer_service.NewConsumerService(consumerRepo)
	operationService := operation_service.NewOperationService(operationRepo, transactionRepo, walletRepo, txRepo, consumerService)
	walletService := wallet_service.NewWalletService(operationService, walletRepo, txRepo, consumerService, _redis.Client)
	server := grpc_server.GrpcServer{
		GrpcPort:            cfg.App.GrpcPort,
		WalletStorageServer: wallet_storage_server.NewWalletStorageServer(walletService, operationService, _logger),
		ArmServer:           consumer_server.NewWalletArmServer(consumerService, _logger),
	}
	go server.StartGrpcServer(cfg, _logger)

	// http Server
	handler := delivery.NewHandler(cfg, walletService, operationService)
	mux := handler.Init()
	httpServer := &http.Server{
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		Addr:         fmt.Sprint(":", cfg.App.HttpPort),
		//Addr:         ":8081",
	}

	// Waiting signal
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	go func() {

		_logger.Info(fmt.Sprintf("listening on :%d", cfg.App.HttpPort))
		if err := httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				fmt.Println(err)
				os.Exit(123)
			}
		}
	}()

	_logger.Info(fmt.Sprintf("shutting down... got signal: %+v", <-stop))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
