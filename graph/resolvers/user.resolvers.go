package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"io/ioutil"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
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

	return new(repositories.UserRepository).Create(&models.User{
		AccountName: accountName,
		Email:       email,
		Password:    string(hash),
		CountryID:   country.ID,
		Country:     country,
	})
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, input *models.UpdateUser) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	country, err := new(repositories.CountryRepository).GetByID(input.CountryID)
	if err != nil {
		return nil, err
	}

	if input.Avatar != nil {
		profilePicture, err := ioutil.ReadAll(input.Avatar.File)
		if err != nil {
			return nil, err
		}

		user.ProfilePicture.File = profilePicture
		user.ProfilePicture.ContentType = input.Avatar.ContentType
		if err := facades.UseDB().Save(&user.ProfilePicture).Error; err != nil {
			return nil, err
		}
	}

	user.DisplayName = input.DisplayName
	user.RealName = input.RealName
	user.CustomURL = input.CustomURL
	user.Summary = input.Summary
	user.ProfileTheme = input.ProfileTheme
	user.CountryID = input.CountryID
	user.Country = *country

	return new(repositories.UserRepository).Update(user)
}

func (r *mutationResolver) SuspendAccount(ctx context.Context, id int64) (*models.User, error) {
	repo := new(repositories.UserRepository)

	user, err := repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.SuspendedAt = time.Now()

	return repo.Update(user)
}

func (r *queryResolver) GetProfile(ctx context.Context, customURL string) (*models.User, error) {
	return new(repositories.UserRepository).GetByCustomURL(customURL)
}

func (r *queryResolver) Users(ctx context.Context, page int64) (*models.UserPagination, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	if user.AccountName != "Admin" {
		return nil, errors.New("not authorized")
	}

	return new(repositories.UserRepository).GetAll(int(page))
}

func (r *userResolver) Wishlist(ctx context.Context, obj *models.User) ([]*models.Game, error) {
	var games []*models.Game
	return games, facades.UseDB().Model(obj).Association("UserWishlist").Find(&games)
}

func (r *userResolver) WishlistCount(ctx context.Context, obj *models.User) (int64, error) {
	return facades.UseDB().Model(obj).Association("UserWishlist").Count(), nil
}

func (r *userResolver) Cart(ctx context.Context, obj *models.User) ([]*models.Game, error) {
	var games []*models.Game
	return games, facades.UseDB().Model(obj).Association("UserCart").Find(&games)
}

func (r *userResolver) CartCount(ctx context.Context, obj *models.User) (int64, error) {
	return facades.UseDB().Model(obj).Association("UserCart").Count(), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
