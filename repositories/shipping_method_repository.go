package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type ShippingMethodRepository struct {
	DB *pg.DB
}

func (sm *ShippingMethodRepository) GetShippingMethods() ([]*models.ShippingMethod, error) {
	var shipping_methods []*models.ShippingMethod
	err := sm.DB.Model(&shipping_methods).Select()
	if err != nil {
		return nil, err
	}

	return shipping_methods, nil
}

func (sm *ShippingMethodRepository) GetShippingMethodByIds(ids []string) ([]*models.ShippingMethod, error) {
	var shipping_methods []*models.ShippingMethod
	err := sm.DB.Model(&shipping_methods).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	return shipping_methods, nil
}

func (sm *ShippingMethodRepository) GetShippingMethodById(id string) (*models.ShippingMethod, error) {
	var shipping_method models.ShippingMethod
	err := sm.DB.Model(&shipping_method).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &shipping_method, nil
}