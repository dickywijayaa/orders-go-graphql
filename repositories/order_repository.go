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

func (o *OrderRepository) GetOrderByIds(ids []string) ([]*models.Order, error) {
	var orders []*models.Order
	err := o.DB.Model(&orders).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepository) GetBuyerOrders(buyer_id string) ([]*models.Order, error) {
	var orders []*models.Order
	err := o.DB.Model(&orders).Where("buyer_id = ?", buyer_id).Select()
	if err != nil {
		return nil, err
	}

	return orders, nil
}


func (o *OrderRepository) GetOrderByBuyerIds(buyer_ids []string) ([]*models.Order, []error) {
	var orders []*models.Order
	err := o.DB.Model(&orders).Where("buyer_id in (?)", pg.In(buyer_ids)).Select()
	if err != nil {
		return nil, []error{err}
	}

	return orders, nil
}

func (o *OrderRepository) CreateOrder(tx *pg.Tx, order *models.Order) (*models.Order, error) {
	_, err := tx.Model(order).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return order, err
}

func (o *OrderRepository) CreateOrderDetails(tx *pg.Tx, order_details []*models.OrderDetail) (error) {
	_, err := tx.Model(&order_details).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) DeleteActiveCart(tx *pg.Tx, cart *models.Cart, cart_detail []*models.CartDetail) (error) {
	_, err := tx.Model(cart).WherePK().Delete()
	if err != nil {
		return err
	}
	
	_, err = tx.Model(&cart_detail).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}