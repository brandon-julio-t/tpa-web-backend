package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := usePreloaded().Find(&users, "account_name != ?", "Admin").Error; err != nil {
		return nil, err
	}
	return users, nil
}

func usePreloaded() *gorm.DB {
	return facades.UseDB().Preload("Country")
}

func (UserRepository) GetByID(id int64) (*models.User, error) {
	var user models.User
	if err := usePreloaded().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) GetByAccountName(accountName string) (*models.User, error) {
	var user models.User
	if err := usePreloaded().First(&user, "account_name = ?", accountName).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) GetByCustomURL(customUrl string) (*models.User, error) {
	var user models.User
	if err := usePreloaded().First(&user, "custom_url = ?", customUrl).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) Create(user *models.User) (*models.User, error) {
	return user, facades.UseDB().Create(user).Error
}

func (UserRepository) Update(user *models.User) (*models.User, error) {
	return user, facades.UseDB().Save(user).Error
}
