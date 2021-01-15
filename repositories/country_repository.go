package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

type CountryRepository struct{}

func (CountryRepository) GetByID(id int64) (*models.Country, error) {
	var country models.Country
	if err := facades.UseDB().First(&country, id).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

func (CountryRepository) GetAll() ([]*models.Country, error) {
	var countries []*models.Country
	if err := facades.UseDB().Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}
