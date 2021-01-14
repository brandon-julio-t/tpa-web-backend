package middlewares

import (
	"context"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
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
		var user models.User

		if err := facades.UseDB().Debug().First(&user, claims.UserID).Error; err != nil {
			log.Print(err)
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), authProviderMiddlewareKey, &user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func UseAuth(ctx context.Context) *models.User {
	user, _ := ctx.Value(authProviderMiddlewareKey).(*models.User)
	return user
}
