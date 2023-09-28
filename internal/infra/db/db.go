package db

import (
	"context"
	"fmt"
	"time"

	"git.example.kz/wallet/wallet-back/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(cfg *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Db)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, ";")
	if err != nil {
		panic(err)
	}

	return pool
}
