package main

import (
	"fmt"
	"os"

	"git.example.kz/wallet/wallet-back/cmd/app"
	"git.example.kz/wallet/wallet-back/cmd/migrate"
	"git.example.kz/wallet/wallet-back/cmd/seed"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "root"}

	rootCmd.AddCommand(app.Cmd)
	rootCmd.AddCommand(seed.Cmd)
	rootCmd.AddCommand(migrate.Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
