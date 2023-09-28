package migrate

import (
	"database/sql"
	"embed"
	"fmt"

	"git.example.kz/wallet/wallet-back/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "migrate",
	Run: NewWalletMigrator,
}

//go:embed migrations
var migrations embed.FS

func NewWalletMigrator(cmd *cobra.Command, args []string) {
	conf := config.GetGeneralConfig()
	// vaultConf := config.GetGeneralConfig()
	fmt.Printf("%+v\n", conf)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Db)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithInstance("iofs", source, "postgres", driver)
	if err != nil {
		panic(err)
	}

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file:///migrations",
	// 	"postgres", driver)
	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		panic(err)
	}
	// return migrate.NewMigrateCommand(&opts)
}
