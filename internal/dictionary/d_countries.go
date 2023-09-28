package dictionary

import (
	"context"

	"git.example.kz/wallet/wallet-back/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Countries struct {
	codeMap map[string]entity.Country
	idMap   map[string]entity.Country
}

func initCountries(conn *pgxpool.Conn) Countries {
	countries := Countries{
		codeMap: make(map[string]entity.Country),
		idMap:   make(map[string]entity.Country),
	}
	rows, err := conn.Conn().Query(context.Background(), "SELECT id,code FROM countries")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var curr entity.Country
		err = rows.Scan(&curr.ID, &curr.Code)
		if err != nil {
			panic(err)
		}
		countries.codeMap[curr.Code] = curr
		countries.idMap[curr.ID] = curr

	}
	return countries
}

func (c *Countries) ByCode(code string) entity.Country {
	return c.codeMap[code]
}

func (c *Countries) ById(id string) entity.Country {
	return c.idMap[id]

}
