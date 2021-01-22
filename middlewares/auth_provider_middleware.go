package middlewares

import (
	"context"
	"errors"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
)

type authProviderMiddleware uint

const authProviderMiddlewareKey authProviderMiddleware = iota

type AuthProviderMiddleware struct{}

func (AuthProviderMiddleware) Create() func(context *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("jwt")
		if err != nil || cookie == nil {
			log.Print(err)
			c.Next()
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &models.UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return facades.UseSecret(), nil
		})

		if err != nil || token == nil {
			log.Print(err)
			c.Next()
			return
		}

		claims, _ := token.Claims.(*models.UserJwtClaims)

		user, err := new(repositories.UserRepository).GetByID(claims.UserID)
		if err != nil {
			log.Print(err)
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), authProviderMiddlewareKey, user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func UseAuth(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(authProviderMiddlewareKey).(*models.User)
	if !ok || user == nil {
		return nil, errors.New("not authenticated")
	}
	return user, nil
}
