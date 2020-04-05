package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type CartDetailRepository struct {
	DB *pg.DB
}

func (c *CartDetailRepository) GetCartDetailByCartId(cart_id string) (*models.CartDetail, error) {
	var cart_detail models.CartDetail
	err := c.DB.Model(&cart_detail).Where("cart_id = ?", cart_id).First()
	if err != nil {
		return nil, err
	}

	return &cart_detail, nil
}