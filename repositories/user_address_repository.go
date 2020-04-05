package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/go-pg/pg/v9"
)

type UserAddressRepository struct {
	DB *pg.DB
}

func (ua *UserAddressRepository) GetUserAddress() ([]*models.UserAddress, error) {
	var user_addresses []*models.UserAddress
	err := ua.DB.Model(&user_addresses).Select()
	if err != nil {
		return nil, err
	}

	return user_addresses, nil
}