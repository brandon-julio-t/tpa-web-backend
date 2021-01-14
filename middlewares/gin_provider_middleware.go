package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ginProviderMiddleware uint

const ginProviderMiddlewareKey ginProviderMiddleware = iota

type GinProviderMiddleware struct{}

func (GinProviderMiddleware) Create() func(context *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginProviderMiddlewareKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func UseGin(ctx context.Context) *gin.Context {
	ginContext, _ := ctx.Value(ginProviderMiddlewareKey).(*gin.Context)
	return ginContext
}
