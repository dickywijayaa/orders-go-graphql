package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type CartRepository struct {
	DB *pg.DB
}

func (c *CartRepository) GetCartById(id string) (*models.Cart, error) {
	var cart models.Cart
	err := c.DB.Model(&cart).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *CartRepository) GetCartByBuyerId(buyer_id string) (*models.Cart, error) {
	var cart models.Cart
	err := c.DB.Model(&cart).Where("buyer_id = ?", buyer_id).First()
	if err != nil {
		return nil, err
	}

	return &cart, nil
}