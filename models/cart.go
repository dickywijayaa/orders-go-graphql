package models

import (
	"time"
)

type Cart struct {
	Id			string		`json:"id"`
	BuyerId		string		`json:"buyer_id"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"-" pg:",soft_delete"`
}