package dictionary

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Currencies struct {
	codeMap map[string]entity.Currency
	idMap   map[string]entity.Currency
}

func initCurrencies(conn *pgxpool.Conn) Currencies {
	currencies := Currencies{
		codeMap: make(map[string]entity.Currency),
		idMap:   make(map[string]entity.Currency),
	}
	rows, err := conn.Conn().Query(context.Background(), "SELECT id,code FROM currencies")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var curr entity.Currency
		err = rows.Scan(&curr.ID, &curr.Code)
		if err != nil {
			panic(err)
		}
		currencies.codeMap[curr.Code] = curr
		currencies.idMap[curr.ID] = curr
	}
	return currencies
}

func (c *Currencies) ByCode(code string) entity.Currency {
	return c.codeMap[code]
}
func (c *Currencies) ById(id string) entity.Currency {
	return c.idMap[id]

}
