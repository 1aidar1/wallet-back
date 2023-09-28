package http

import (
	"net/http"

	"git.example.kz/wallet/wallet-back/config"
	"git.example.kz/wallet/wallet-back/internal/controller"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg       *config.Config
	wallet    controller.WalletService
	operation controller.OperationService
}

type Deps struct {
}

func NewHandler(cfg *config.Config, wallet controller.WalletService, operation controller.OperationService) *Handler {
	return &Handler{
		cfg:       cfg,
		wallet:    wallet,
		operation: operation,
	}
}

func (h *Handler) Init() *gin.Engine {
	// Init gin handler
	router := gin.New()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	//handlerV1 := v1.NewV1Handler(h.cfg)
	router.Use(CORSMiddleware(), JSONMiddleware())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			operations := v1.Group("/operations")
			//operations.Use(h.LockMiddleware())
			{
				operations.POST("/refill", h.OperationRefill)
				operations.POST("/withdraw", h.OperationWithdraw)
				operations.POST("/confirm", h.OperationConfirm)
				operations.POST("/hold", h.OperationHold)
				operations.POST("/clear", h.OperationClear)
				operations.POST("/unhold", h.OperationUnhold)
				operations.POST("/transfer", h.OperationTransfer)
				operations.POST("/refund", h.OperationRefund)
			}
			wallet := v1.Group("/wallet")
			{
				wallet.POST("/create", h.WalletCreate)
				wallet.POST("/close", h.WalletClose)
				wallet.POST("/block", h.WalletBlock)
				wallet.POST("/unblock", h.WalletUnblock)
				wallet.POST("/identification", h.WalletIdentification)
				wallet.GET("/info", h.WalletInfo)
				wallet.GET("/history", h.WalletHistory)
				wallet.GET("/statistics", h.WalletStatistics)
				wallet.GET("/transaction_info", h.TransactionInfo)
			}
		}
	}
}
