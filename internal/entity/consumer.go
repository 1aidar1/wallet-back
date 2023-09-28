package entity

import (
	"time"
)

type Consumer struct {
	ID               string     `json:"id"`
	Code             string     `json:"code"`
	Slug             string     `json:"slug"`
	Secret           string     `json:"secret"`
	WhiteListMethods []string   `json:"white_list_methods"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}
