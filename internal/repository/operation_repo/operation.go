package operation_repo

import (
	"context"
	"fmt"
	"time"

	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OperationRepo struct {
	db *pgxpool.Pool
}

func NewOperationRepo(db *pgxpool.Pool) *OperationRepo {
	return &OperationRepo{
		db: db,
	}
}

type operationModel struct {
	Id                 uuid.UUID  `db:"id"`
	TransactionId      uuid.UUID  `db:"transaction_id"`
	Type               string     `db:"type"`
	WalletId           uuid.UUID  `db:"wallet_id"`
	Amount             int        `db:"amount"`
	Status             string     `db:"status"`
	ErrorCode          string     `db:"error_code"`
	InternalLogMessage string     `db:"internal_log_message"`
	BalanceBefore      int        `db:"balance_before"`
	BalanceAfter       int        `db:"balance_after"`
	HoldBalanceBefore  int        `db:"hold_balance_before"`
	HoldBalanceAfter   int        `db:"hold_balance_after"`
	CreatedAt          *time.Time `db:"created_at"`
}

func (r *OperationRepo) Create(ctx context.Context, operation entity.Operation, tx pgx.Tx) (entity.Operation, *errcode.ErrCode) {
	model := operationModel{
		TransactionId:      uuid.FromStringOrNil(operation.TransactionId),
		Type:               operation.Type,
		WalletId:           uuid.FromStringOrNil(operation.WalletId),
		Amount:             operation.Amount,
		Status:             operation.Status,
		ErrorCode:          operation.ErrorCode,
		InternalLogMessage: operation.InternalLogMessage,
		BalanceBefore:      operation.BalanceBefore,
		BalanceAfter:       operation.BalanceAfter,
		HoldBalanceBefore:  operation.HoldBalanceBefore,
		HoldBalanceAfter:   operation.HoldBalanceAfter,
	}
	var out entity.Operation
	q := `INSERT INTO operations (transaction_id, type, wallet_id, amount, status, error_code, internal_log_message, 
		balance_before, balance_after, hold_balance_before, hold_balance_after) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING *`
	rows, err := r.query(ctx, tx, q,
		model.TransactionId, model.Type, model.WalletId, model.Amount, model.Status, model.ErrorCode,
		model.InternalLogMessage, model.BalanceBefore, model.BalanceAfter, model.HoldBalanceBefore, model.HoldBalanceAfter)
	if err != nil {
		return entity.Operation{}, errcode.NewDefaultErr().WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", operation))
	}
	err = pgxscan.ScanOne(&out, rows)
	if err != nil {
		return entity.Operation{}, errcode.NewDefaultErr().WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", operation))
	}

	return out, nil
}

func (r *OperationRepo) query(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) (pgx.Rows, error) {
	if conn == nil {
		return r.db.Query(ctx, q, args...)
	} else {
		return conn.Query(ctx, q, args...)
	}
}
