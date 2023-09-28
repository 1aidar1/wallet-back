package atomic_repo

import (
	"context"

	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AtomicRepo struct {
	Db *pgxpool.Pool
}

func NewAtomicRepo(db *pgxpool.Pool) *AtomicRepo {
	return &AtomicRepo{
		Db: db,
	}
}

func (r *AtomicRepo) Do(ctx context.Context, f func(tx pgx.Tx) *errcode.ErrCode) (e *errcode.ErrCode) {
	tx, err := r.Db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return errcode.NewDefaultErr().WithErr(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil || e != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	e = f(tx)
	if e != nil {
		return e
	}

	return nil
}
