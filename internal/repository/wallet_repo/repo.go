package wallet_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/utils"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepo struct {
	db *pgxpool.Pool
}

func NewWalletRepo(db *pgxpool.Pool) *WalletRepo {
	return &WalletRepo{
		db: db,
	}
}

type walletModel struct {
	ID             uuid.UUID      `db:"id"`
	Phone          sql.NullString `db:"phone"`
	CurrencyID     uuid.UUID      `db:"currency_id"`
	CountryID      uuid.UUID      `db:"country_id"`
	Balance        int            `db:"balance"`
	Hold           int            `db:"hold"`
	Status         string         `db:"status"`
	Identification string         `db:"identification"`
	UpdatedAt      *time.Time     `db:"updated_at"`
	CreatedAt      *time.Time     `db:"created_at"`
}

func (r *WalletRepo) Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, *errcode.ErrCode) {
	model := walletModel{
		CurrencyID:     uuid.FromStringOrNil(wallet.CurrencyID),
		CountryID:      uuid.FromStringOrNil(wallet.CountryID),
		Balance:        wallet.Balance,
		Hold:           wallet.Hold,
		Phone:          utils.NewNullString(wallet.Phone),
		Status:         wallet.Status,
		Identification: wallet.Identification,
	}
	q := `INSERT INTO wallets (currency_id, country_id, balance, hold, status,identification,phone) 
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *`

	rows, err := r.query(ctx, nil, q,
		model.CurrencyID, model.CountryID, model.Balance, model.Hold, model.Status, model.Identification, model.Phone)
	if err != nil {
		return entity.Wallet{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", wallet))
	}
	var created walletModel
	err = pgxscan.ScanOne(&created, rows)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation && e.ConstraintName == wallets_unique_phone_constraint {
			return entity.Wallet{}, errcode.New("phone_already_registered").WithErr(err)
		}
		return entity.Wallet{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", wallet))
	}
	return entity.Wallet{
		ID:             created.ID.String(),
		Phone:          created.Phone.String,
		CurrencyID:     created.CurrencyID.String(),
		CountryID:      created.CountryID.String(),
		Balance:        created.Balance,
		Hold:           created.Hold,
		Status:         created.Status,
		Identification: created.Identification,
		CreatedAt:      created.CreatedAt,
		UpdatedAt:      created.UpdatedAt,
	}, nil
}

func (r *WalletRepo) Update(ctx context.Context, wallet entity.Wallet, tx pgx.Tx) *errcode.ErrCode {
	model := walletModel{
		ID:      uuid.FromStringOrNil(wallet.ID),
		Balance: wallet.Balance,
		Hold:    wallet.Hold,
		Status:  wallet.Status,
	}
	q := "UPDATE wallets SET balance=$1, hold=$2, status=$3 where id=$4"

	_, err := r.exec(ctx, tx, q, model.Balance, model.Hold, model.Status, model.ID)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", model))
	}

	return nil
}

func (r *WalletRepo) UpdateStatus(ctx context.Context, wallet dto.UpdateWalletStatus, tx pgx.Tx) *errcode.ErrCode {

	q := "UPDATE wallets SET status=$1 where id=$2"

	_, err := r.exec(ctx, tx, q, wallet.NewStatus, wallet.WalletId)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from update status"))
	}

	_, err = r.exec(ctx, tx, "INSERT INTO wallet_status_history (wallet_id, status,description) VALUES ($1,$2,$3)",
		wallet.WalletId, wallet.NewStatus, wallet.Description)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("model: %+v", wallet))
	}
	return nil

}

func (r *WalletRepo) UpdateIdentification(ctx context.Context, walletId, identification string, tx pgx.Tx) *errcode.ErrCode {

	q := "UPDATE wallets SET identification=$1 where id=$2"

	_, err := r.exec(ctx, tx, q, identification, walletId)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from update identification"))
	}

	_, err = r.exec(ctx, tx, "INSERT INTO wallet_identification_history (wallet_id, status) VALUES ($1,$2)",
		walletId, identification)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err creating id history"))
	}
	return nil
}

func (r *WalletRepo) AddBalance(ctx context.Context, walletId string, balance, hold int, tx pgx.Tx) *errcode.ErrCode {
	_, err := r.exec(ctx, tx, "UPDATE wallets SET balance=balance+$1, hold=hold+$2 where id=$3",
		balance, hold, walletId)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprint("req:", walletId, balance, hold))
	}

	return nil
}

func (r *WalletRepo) Stat(ctx context.Context, req dto.WalletStatisticsReq) (dto.WalletStatisticsRes, *errcode.ErrCode) {

	rows, err := r.query(ctx, nil,
		`select type,sum(amount),count(*) from operations WHERE wallet_id = $1
    		AND status = $2
    		AND created_at BETWEEN $3::timestamp AND $4::timestamp
		group by type;`,
		req.WalletId, req.Status, req.DateStart, req.DateEnd)

	if err != nil {
		return dto.WalletStatisticsRes{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("req: %+v", req))
	}

	res := make(dto.WalletStatisticsRes)
	var (
		operationType string
		sum           int
		count         int
	)
	for rows.Next() {

		err = rows.Scan(&operationType, &sum, &count)
		if err != nil {
			return dto.WalletStatisticsRes{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("req: %+v", req))
		}
		res[operationType] = dto.Stats{
			Amount: sum,
			Count:  count,
		}
	}

	return res, nil
}

func (r *WalletRepo) Find(ctx context.Context, id string, tx pgx.Tx) (entity.Wallet, *errcode.ErrCode) {
	var model walletModel
	row, err := r.query(ctx, tx, "SELECT * FROM wallets where id=$1", id)
	if err != nil {
		return entity.Wallet{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("id: %s", id))
	}

	err = pgxscan.ScanOne(&model, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Wallet{}, errcode.New("wallet_not_found").WithErr(err).WithMsg(fmt.Sprintf("id: %s", id))
		}
		return entity.Wallet{}, errcode.New(errcode.DefaultCode).WithErr(err)
	}
	return entity.Wallet{
		ID:             model.ID.String(),
		CurrencyID:     model.CurrencyID.String(),
		CountryID:      model.CountryID.String(),
		Phone:          model.Phone.String,
		Balance:        model.Balance,
		Hold:           model.Hold,
		Status:         model.Status,
		Identification: model.Identification,
		CreatedAt:      model.CreatedAt,
	}, nil
}

func (r *WalletRepo) FindByPhone(ctx context.Context, phone string, tx pgx.Tx) (entity.Wallet, *errcode.ErrCode) {
	var model walletModel
	row, err := r.query(ctx, tx, "SELECT * FROM wallets where phone=$1", phone)
	if err != nil {
		return entity.Wallet{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("phone: %s", phone))
	}

	err = pgxscan.ScanOne(&model, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Wallet{}, errcode.New("wallet_not_found").WithErr(err).WithMsg(fmt.Sprintf("phone: %s", phone))
		}
		return entity.Wallet{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("db_out: %s", row.RawValues()))
	}
	return entity.Wallet{
		ID:             model.ID.String(),
		CurrencyID:     model.CurrencyID.String(),
		CountryID:      model.CountryID.String(),
		Balance:        model.Balance,
		Hold:           model.Hold,
		Status:         model.Status,
		Identification: model.Identification,
		CreatedAt:      model.CreatedAt,
	}, nil
}

func (r *WalletRepo) History(ctx context.Context, history dto.WalletHistoryReq, tx pgx.Tx) ([]entity.WalletHistory, *errcode.ErrCode) {

	// alt+enter -> convert to raw string
	q := `
SELECT
    t.id,
    t.consumer_id,
    t.type,
    t.description,
    t.order_id,
    t.service_provider_id,
    t.created_at,
    (
        SELECT json_agg(
            json_build_object(
                'id', o.id,
                'type', o.type,
                'wallet_id', o.wallet_id,
                'amount', o.amount,
                'status', o.status,
                'error_code', o.error_code,
                'balance_before', o.balance_before,
                'balance_after', o.balance_after,
                'hold_balance_before', o.hold_balance_before,
                'hold_balance_after', o.hold_balance_after,
                'created_at', o.created_at
            )
        )
        FROM operations o
        WHERE o.transaction_id = t.id
    ) AS operations
FROM transactions t
INNER JOIN operations op
    ON t.id = op.transaction_id
WHERE op.wallet_id = $1
    AND t.created_at >= $2
    AND t.created_at < $3
	%s
GROUP BY t.id, t.type, t.description, t.order_id,t.created_at
ORDER BY t.created_at DESC
LIMIT $4
OFFSET $5;
`

	params := []interface{}{history.WalletId, history.DateStart, history.DateEnd, history.PerPage, (history.Page - 1) * history.PerPage}
	var whereQ string
	if history.ProviderId != nil && *history.ProviderId != "" {
		whereQ = "AND t.service_provider_id = $6"
		params = append(params, history.ProviderId)
	}
	q = fmt.Sprintf(q, whereQ)

	rows, err := r.query(ctx, tx, q, params...)
	if err != nil {
		return []entity.WalletHistory{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err query from repo. request: %+v", history))
	}

	var txns []WalletHistoryTxModel
	for rows.Next() {
		var txn WalletHistoryTxModel
		var ops []WalletHistoryOperationModel

		err = rows.Scan(
			&txn.Id,
			&txn.ConsumerId,
			&txn.Type,
			&txn.Description,
			&txn.OrderId,
			&txn.ServiceProvideId,
			&txn.CreatedAt,
			&ops,
		)
		if err != nil {
			return []entity.WalletHistory{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err query from repo. request: %+v", history))
		}

		txn.Operations = ops
		txns = append(txns, txn)
	}
	return transformToEntitySlice(txns), nil
}

func (r *WalletRepo) exec(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) (pgconn.CommandTag, error) {
	if conn == nil {
		return r.db.Exec(ctx, q, args...)
	} else {
		return conn.Exec(ctx, q, args...)
	}
}

func (r *WalletRepo) query(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) (pgx.Rows, error) {
	if conn == nil {
		return r.db.Query(ctx, q, args...)
	} else {
		return conn.Query(ctx, q, args...)
	}
}

func (r *WalletRepo) queryRow(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) pgx.Row {
	if conn == nil {
		return r.db.QueryRow(ctx, q, args...)

	} else {
		return conn.QueryRow(ctx, q, args...)
	}
}
