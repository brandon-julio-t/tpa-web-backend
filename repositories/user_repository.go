package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"math"
)

type UserRepository struct{}

func (UserRepository) GetAll(page int) (*models.UserPagination, error) {
	userPerPage := 5
	var users []*models.User
	var count int64
	if err := usePreloadedUser().
		Model(&models.User{}).
		Count(&count).
		Scopes(facades.UsePagination(page, userPerPage)).
		Find(&users, "account_name != ?", "Admin").Error; err != nil {
		return nil, err
	}
	return &models.UserPagination{
		Data:       users,
		TotalPages: int(math.Ceil(float64(count)/float64(userPerPage))),
	}, nil
}

func usePreloadedUser() *gorm.DB {
	return facades.UseDB().Preload("Country").Preload("ProfilePicture")
}

func (UserRepository) GetByID(id int64) (*models.User, error) {
	var user models.User
	if err := usePreloadedUser().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) GetByAccountName(accountName string) (*models.User, error) {
	var user models.User
	if err := usePreloadedUser().First(&user, "account_name = ?", accountName).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) GetByCustomURL(customUrl string) (*models.User, error) {
	var user models.User
	if err := usePreloadedUser().First(&user, "custom_url = ?", customUrl).Error; err != nil {
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
