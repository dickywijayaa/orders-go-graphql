package models

type OrderDetail struct {
	Id		 		string	`json:"id"`
	OrderId			string	`json:"order_id"`
	ItemName		string 	`json:"item_name"`
	ItemPrice 		int 	`json:"item_price"`
	ItemQuantity 	int		`json:"item_quantity"`
}