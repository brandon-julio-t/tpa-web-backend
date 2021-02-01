package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Login(ctx context.Context, accountName string, password string) (*models.User, error) {
	user, err := new(repositories.UserRepository).GetByAccountName(accountName)
	if err != nil {
		return nil, err
	}

	if !user.SuspendedAt.IsZero() {
		return nil, errors.New("account suspended")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	jwtExpireDuration, err := time.ParseDuration("15m")
	if err != nil {
		return nil, err
	}

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, models.UserJwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpireDuration).Unix(),
		},
		UserID: user.ID,
	}).SignedString(facades.UseSecret())

	if err != nil {
		return nil, err
	}

	if os.Getenv("ENV") == "production" {
		middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", (&http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			Expires:  time.Now().Add(jwtExpireDuration),
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}).String())
	} else {
		middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", (&http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			Expires:  time.Now().Add(jwtExpireDuration),
			HttpOnly: true,
		}).String())
	}

	user.Status = "online"
	return user, facades.UseDB().Save(user).Error
}

func (r *mutationResolver) Logout(ctx context.Context) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	if os.Getenv("ENV") == "production" {
		middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", (&http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Time{},
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			MaxAge:   0,
		}).String())
	} else {
		middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", (&http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Time{},
			HttpOnly: true,
			MaxAge:   0,
		}).String())
	}

	user.Status = "offline"
	return user, facades.UseDB().Save(user).Error
}

func (r *queryResolver) Auth(ctx context.Context) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if user != nil {
		return user, nil
	}
	return nil, err
}

func (r *queryResolver) RefreshToken(ctx context.Context) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if user == nil {
		return false, errors.New("not authenticated")
	}

	jwtExpireDuration, err := time.ParseDuration("15m")
	if err != nil {
		return false, err
	}

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, models.UserJwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpireDuration).Unix(),
		},
		UserID: user.ID,
	}).SignedString(facades.UseSecret())

	if err != nil {
		return false, err
	}

	var cookie *http.Cookie

	if os.Getenv("ENV") == "production" {
		cookie = &http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			Expires:  time.Now().Add(jwtExpireDuration),
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
	} else {
		cookie = &http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			Expires:  time.Now().Add(jwtExpireDuration),
			HttpOnly: true,
		}
	}

	middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", cookie.String())

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
