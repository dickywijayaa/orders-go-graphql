package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type OrderRepository struct {
	DB *pg.DB
}

func (o *OrderRepository) GetOrder() ([]*models.Order, error) {
	var orders []*models.Order
	err := o.DB.Model(&orders).Select()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepository) GetOrderTotalPrice(id string) (float64, error) {
	var order models.Order
	err := o.DB.Model(&order).Where("id = ?", id).First()
	if err != nil {
		return 0, err
	}

	return order.TotalPrice, nil
}

func (o *OrderRepository) GetOrderById(id string) (*models.Order, error) {
	var order models.Order
	err := o.DB.Model(&order).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *OrderRepository) GetBuyerOrders(buyer_id string) ([]*models.Order, error) {
	var orders []*models.Order
	err := o.DB.Model(&orders).Where("buyer_id = ?", buyer_id).Select()
	if err != nil {
		return nil, err
	}

	return orders, nil
}