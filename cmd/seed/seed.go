package seed

import (
	"context"
	_ "embed"

	"git.example.kz/wallet/wallet-back/config"
	"git.example.kz/wallet/wallet-back/internal/infra/db"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"github.com/spf13/cobra"
)

//go:embed seed.sql
var seed []byte

var Cmd = &cobra.Command{
	Use: "seed",
	Run: NewWalletSeeder,
}

func NewWalletSeeder(cmd *cobra.Command, args []string) {
	cfg := config.GetGeneralConfig()

	_db := db.NewDB(cfg)
	_, err := _db.Exec(context.Background(), string(seed))
	if err != nil {
		logger.GetProdLogger().Error("seed failed", err)
	} else {
		logger.GetProdLogger().Info("seed success")

	}
	return
}
