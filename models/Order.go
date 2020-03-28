package models

type Order struct {
	ID 			string	`json:"id"`
	BuyerId 	string	`json:"buyer_id"`
	TotalPrice 	float64	`json:"total_price"`
}