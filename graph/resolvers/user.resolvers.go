package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Register(ctx context.Context, accountName string, email string, password string, countryID int64) (*models.User, error) {
	var country models.Country
	if err := facades.UseDB().First(&country, countryID).Error; err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		BaseModel:   models.BaseModel{},
		AccountName: accountName,
		Email:       email,
		Password:    string(hash),
		CountryID:   country.ID,
		Country:     country,
	}

	if err := facades.UseDB().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
