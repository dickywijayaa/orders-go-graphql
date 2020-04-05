package models

import (
	"time"
)

type CartDetail struct {
	Id			string		`json:"id"`
	ProductId	string		`json:"product_id"`
	Quantity	int 		`json:"quantity"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"-" pg:",softdelete"`
}