package transaction_repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

type transactionModel struct {
	Id               string           `db:"id"`
	ConsumerId       string           `db:"consumer_id"`
	ServiceProvideId string           `db:"service_provider_id"`
	Type             string           `db:"type"`
	Description      string           `db:"description"`
	OrderId          string           `db:"order_id"`
	CreatedAt        *time.Time       `db:"created_at"`
	Operations       []operationModel `db:"-"`
}

type operationModel struct {
	Id                 string     `db:"id"`
	TransactionId      string     `db:"transaction_id"`
	Type               string     `db:"type"`
	WalletId           string     `db:"wallet_id"`
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

func (r *TransactionRepo) Create(ctx context.Context, tr entity.Transaction, tx pgx.Tx) (string, *errcode.ErrCode) {
	model := transactionModel{
		ConsumerId:       tr.ConsumerId,
		ServiceProvideId: tr.ServiceProvideId,
		Type:             tr.Type,
		Description:      tr.Description,
		OrderId:          tr.OrderId,
	}
	var transactionId string
	err := r.queryRow(ctx, tx, "INSERT INTO transactions (consumer_id,service_provider_id,description,type, order_id) VALUES ($1,$2,$3,$4,$5) returning id",
		model.ConsumerId, model.ServiceProvideId, model.Description, model.Type, model.OrderId).Scan(&transactionId)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation && e.ConstraintName == transactions_unique_order_id_constraint {
			return "", errcode.New("order_already_exists").WithErr(err)
		}
		return "", errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", tr))
	}
	return transactionId, nil
}

func (r *TransactionRepo) Find(ctx context.Context, id string, tx pgx.Tx) (entity.Transaction, *errcode.ErrCode) {

	rows, err := r.query(ctx, tx, "SELECT * FROM transactions WHERE id=$1", id)
	if err != nil {
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err query from repo. request: %+v", tx))
	}

	var transaction entity.Transaction
	err = pgxscan.ScanOne(&transaction, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Transaction{}, errcode.New("transaction_not_found").WithErr(err).WithMsg(fmt.Sprintf("id: %s", id))
		}
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err scan from repo. request: %+v", id))
	}

	var operations []entity.Operation
	rows, err = r.query(ctx, tx, "SELECT * FROM operations WHERE transaction_id=$1",
		id)
	if err != nil {
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err query from repo. request: %+v", tx))
	}
	err = pgxscan.ScanAll(&operations, rows)
	if err != nil {
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err scan from repo. request: %+v", tx))
	}
	transaction.Operations = operations

	return transaction, nil
}
func (r *TransactionRepo) FindByOrder(ctx context.Context, orderId, providerCode string, tx pgx.Tx) (entity.Transaction, *errcode.ErrCode) {

	rows, err := r.query(ctx, tx, "SELECT * FROM transactions WHERE order_id=$1 and service_provider_id=$2", orderId, providerCode)
	if err != nil {
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err query from repo. request: %+v", tx))
	}

	var transaction entity.Transaction
	err = pgxscan.ScanOne(&transaction, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Transaction{}, errcode.New("transaction_not_found").WithErr(err).WithMsg(fmt.Sprintf("order_id: %s, provider: %s", orderId, providerCode))
		}
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err scan from repo. order_id: %s, provider: %s:", orderId, providerCode))
	}

	var operations []entity.Operation
	rows, err = r.query(ctx, tx, "SELECT * FROM operations WHERE transaction_id=$1",
		transaction.Id)
	if err != nil {
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err query from repo. request: %+v", tx))
	}
	err = pgxscan.ScanAll(&operations, rows)
	if err != nil {
		return entity.Transaction{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err scan from repo. request: %+v", tx))
	}
	transaction.Operations = operations

	return transaction, nil
}

func (r *TransactionRepo) query(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) (pgx.Rows, error) {
	if conn == nil {
		return r.db.Query(ctx, q, args...)
	} else {
		return conn.Query(ctx, q, args...)
	}
}

func (r *TransactionRepo) queryRow(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) pgx.Row {
	if conn == nil {
		return r.db.QueryRow(ctx, q, args...)

	} else {
		return conn.QueryRow(ctx, q, args...)
	}
}
