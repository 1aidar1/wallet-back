package consumer_repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"git.example.kz/wallet/wallet-back/internal/dto"
	"git.example.kz/wallet/wallet-back/internal/entity"
	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConsumerRepo struct {
	db *pgxpool.Pool
}

func NewConsumerRepo(db *pgxpool.Pool) *ConsumerRepo {
	return &ConsumerRepo{
		db: db,
	}
}

func (r *ConsumerRepo) Create(ctx context.Context, req dto.ConsumerCreateReq, tx pgx.Tx) (entity.Consumer, *errcode.ErrCode) {
	model := consumerModel{
		Code:             req.Code,
		Slug:             req.Slug,
		Secret:           req.Secret,
		WhiteListMethods: req.WhiteListMethods,
	}
	rows, err := r.query(ctx, tx, "INSERT INTO consumers (code, slug, secret, white_list_methods) VALUES ($1,$2,$3,$4) RETURNING *",
		model.Code, model.Slug, model.Secret, model.WhiteListMethods)
	if err != nil {
		return entity.Consumer{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", req))
	}
	var created consumerModel
	err = pgxscan.ScanOne(&created, rows)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation && e.ConstraintName == consumers_unique_code_constraint {
			return entity.Consumer{}, errcode.New("consumer_code_taken").WithErr(err)
		}
		return entity.Consumer{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err scan from repo. request: %+v", req))
	}

	return entity.Consumer{
		ID:               created.ID,
		Code:             created.Code,
		Slug:             created.Slug,
		Secret:           created.Secret,
		WhiteListMethods: created.WhiteListMethods,
		CreatedAt:        created.CreatedAt,
		UpdatedAt:        created.UpdatedAt,
		DeletedAt:        created.DeletedAt,
	}, nil
}

func (r *ConsumerRepo) Read(ctx context.Context, consumerId string, tx pgx.Tx) (entity.Consumer, *errcode.ErrCode) {
	var model consumerModel
	row, err := r.query(ctx, tx, "SELECT * FROM consumers where id=$1", consumerId)
	if err != nil {
		return entity.Consumer{}, errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("id: %s", consumerId))
	}

	err = pgxscan.ScanOne(&model, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Consumer{}, errcode.New("consumer_not_found").WithErr(err).WithMsg(fmt.Sprintf("id: %s", consumerId))
		}
		return entity.Consumer{}, errcode.New(errcode.DefaultCode).WithErr(err)
	}
	return entity.Consumer{
		ID:               model.ID,
		Code:             model.Code,
		Slug:             model.Slug,
		Secret:           model.Secret,
		WhiteListMethods: model.WhiteListMethods,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        model.DeletedAt,
	}, nil
}

func (r *ConsumerRepo) Update(ctx context.Context, id string, req dto.ConsumerUpdateReq, tx pgx.Tx) *errcode.ErrCode {

	q := "UPDATE consumers SET code=$1, slug=$2, secret=$3, white_list_methods = $4, updated_at = $5 where id=$6"
	_, err := r.exec(ctx, tx, q, req.Code, req.Slug, req.Secret, req.WhiteListMethods, time.Now(), id)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation && e.ConstraintName == consumers_unique_code_constraint {
			return errcode.New("consumer_code_taken").WithErr(err)
		}
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", req))
	}

	return nil
}

func (r *ConsumerRepo) Delete(ctx context.Context, id string, tx pgx.Tx) *errcode.ErrCode {

	q := "UPDATE consumers SET deleted_at=$1 where id=$2"
	_, err := r.exec(ctx, tx, q, time.Now(), id)
	if err != nil {
		return errcode.New(errcode.DefaultCode).WithErr(err).WithMsg(fmt.Sprintf("err from repo. request: %+v", id))
	}

	return nil
}

func (r *ConsumerRepo) List(ctx context.Context, tx pgx.Tx) ([]entity.Consumer, *errcode.ErrCode) {
	rows, err := r.query(ctx, tx, "SELECT * FROM consumers where deleted_at is null")
	if err != nil {
		return nil, errcode.New(errcode.DefaultCode).WithErr(err)
	}

	list := make([]entity.Consumer, 0)
	for rows.Next() {
		var current consumerModel
		err = pgxscan.ScanRow(&current, rows)
		if err != nil {
			return nil, errcode.New(errcode.DefaultCode).WithErr(err)
		}
		list = append(list, entity.Consumer{
			ID:               current.ID,
			Code:             current.Code,
			Slug:             current.Slug,
			Secret:           current.Secret,
			WhiteListMethods: current.WhiteListMethods,
			CreatedAt:        current.CreatedAt,
			UpdatedAt:        current.UpdatedAt,
			DeletedAt:        current.DeletedAt,
		})
	}
	return list, nil
}

func (r *ConsumerRepo) exec(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) (pgconn.CommandTag, error) {
	if conn == nil {
		return r.db.Exec(ctx, q, args...)
	} else {
		return conn.Exec(ctx, q, args...)
	}
}

func (r *ConsumerRepo) query(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) (pgx.Rows, error) {
	if conn == nil {
		return r.db.Query(ctx, q, args...)
	} else {
		return conn.Query(ctx, q, args...)
	}
}

func (r *ConsumerRepo) queryRow(ctx context.Context, conn pgx.Tx, q string, args ...interface{}) pgx.Row {
	if conn == nil {
		return r.db.QueryRow(ctx, q, args...)

	} else {
		return conn.QueryRow(ctx, q, args...)
	}
}
