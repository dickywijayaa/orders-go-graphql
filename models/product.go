package models

import (
	"time"
)

type Product struct {
	ID 			string		`json:"id"`
	SellerId	string 		`json:"seller_id"`
	Name 		string		`json:"name"`
	Price 		float64 	`json:"price"`
	Weight		float64		`json:"weight"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"-" pg:",softdelete"`
}