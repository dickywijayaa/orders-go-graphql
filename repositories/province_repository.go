package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type ProvinceRepository struct {
	DB *pg.DB
}

func (p *ProvinceRepository) GetProvinces(ids []string) ([]*models.Province, error) {
	var provinces []*models.Province
	err := p.DB.Model(&provinces).Select()
	if err != nil {
		return nil, err
	}

	return provinces, nil
}

func (p *ProvinceRepository) GetProvinceById(id string) (*models.Province, error) {
	var province *models.Province
	err := p.DB.Model(&province).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return province, nil
}

func (p *ProvinceRepository) GetProvinceByIds(ids []string) ([]*models.Province, error) {
	var provinces []*models.Province
	err := p.DB.Model(&provinces).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	return provinces, nil
}