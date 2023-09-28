package dictionary

import (
	"context"

	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CodeErrors struct {
	m map[string]errcode.ErrCode
}

func (e *CodeErrors) Get(code string) *errcode.ErrCode {
	err, ok := e.m[code]
	if !ok {
		return errcode.NewDefaultErr()
	}
	return &err
}

func initErrors(conn *pgxpool.Conn) CodeErrors {
	errs := make(map[string]errcode.ErrCode)
	rows, err := conn.Conn().Query(context.Background(), "SELECT code,description,http_code FROM error_codes")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var curr errcode.ErrCode
		err = rows.Scan(&curr.Code, &curr.Desc, &curr.Http)
		if err != nil {
			panic(err)
		}
		errs[curr.Code] = curr
	}
	return CodeErrors{
		m: errs,
	}
}
