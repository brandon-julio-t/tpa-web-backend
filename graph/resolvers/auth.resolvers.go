package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*models.User, error) {
	var user models.User

	if err := facades.UseDB().First(&user, "email = ?", email).Error; err != nil {
		return nil, err
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

	middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", (&http.Cookie{
		Name:     "jwt",
		Value:    jwtToken,
		Expires:  time.Now().Add(jwtExpireDuration),
		HttpOnly: true,
	}).String())

	return &user, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*models.User, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return nil, errors.New("not signed in")
	}

	middlewares.UseGin(ctx).Writer.Header().Add("Set-Cookie", (&http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Time{},
		HttpOnly: true,
	}).String())

	return user, nil
}

func (r *queryResolver) Auth(ctx context.Context) (*models.User, error) {
	if user := middlewares.UseAuth(ctx); user != nil {
		return user, nil
	}
	return nil, errors.New("not signed in")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
