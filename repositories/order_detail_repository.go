package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type OrderDetailRepository struct {
	DB *pg.DB
}

func (od *OrderDetailRepository) GetOrderDetailsByOrderId(order_id string) ([]*models.OrderDetail, error) {
	var order_details []*models.OrderDetail 

	err := od.DB.Model(&order_details).Where("order_id = ?", order_id).Select()
	if err != nil {
		return nil, err
	}

	return order_details, nil
}

func (od *OrderDetailRepository) GetOrderDetails(order_ids []string) ([]*models.OrderDetail, error) {
	var order_details []*models.OrderDetail

	err := od.DB.Model(&order_details).Where("order_id in (?)", pg.In(order_ids)).Select()
	if err != nil {
		return nil, err
	}

	return order_details, nil
}