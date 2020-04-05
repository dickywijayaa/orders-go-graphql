package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type ProductRepository struct {
	DB *pg.DB
}

func (p *ProductRepository) GetProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := p.DB.Model(&products).Select()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProductById(id string) (*models.Product, error) {
	var product *models.Product
	err := p.DB.Model(&product).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) GetProductByIds(ids []string) ([]*models.Product, error) {
	var products []*models.Product
	err := p.DB.Model(&products).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	return products, nil
}