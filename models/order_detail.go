package models

import (
	"time"
)

type OrderDetail struct {
	Id		 			string	`json:"id"`
	OrderId				string	`json:"order_id"`
	SellerId			string	`json:"seller_id"`
	ItemId				string	`json:"item_id"`
	ItemName			string 	`json:"item_name"`
	ItemPrice 			float64 `json:"item_price"`
	ItemQuantity	 	int		`json:"item_quantity"`
	ItemWeight		 	float64	`json:"item_weight"`
	ShippingMethodId	string	`json:"shipping_method_id"`
	ShippingCost		float64	`json:"shipping_cost"`
	CreatedAt			time.Time 	`json:"created_at"`
	UpdatedAt			time.Time 	`json:"updated_at"`
	DeletedAt			*time.Time	`json:"-" pg:",softdelete"`
}