package models

import (
	"time"
)

type Order struct {
	ID 					string		`json:"id"`
	BuyerId			 	string		`json:"buyer_id"`
	TotalPrice 			float64		`json:"total_price"`
	TotalShippingCost	float64		`json:"total_shipping_cost"`
	CreatedAt			time.Time 	`json:"created_at"`
	UpdatedAt			time.Time 	`json:"updated_at"`
	DeletedAt			*time.Time	`json:"-" pg:",soft_delete"`
}