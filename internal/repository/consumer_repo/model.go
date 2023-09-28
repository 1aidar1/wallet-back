package consumer_repo

import (
	"time"
)

type consumerModel struct {
	ID               string     `db:"id"`
	Code             string     `db:"code"`
	Slug             string     `db:"slug"`
	Secret           string     `db:"secret"`
	WhiteListMethods []string   `db:"white_list_methods"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
	DeletedAt        *time.Time `db:"deleted_at"`
}
