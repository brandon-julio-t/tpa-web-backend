package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/resolvers"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
)

const (
	graphqlEndpoint = "/graphql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"http://localhost:4200", "https://tpa-web-br20-2.netlify.app"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(new(middlewares.GinProviderMiddleware).Create())
	r.Use(new(middlewares.AuthProviderMiddleware).Create())

	r.GET("/", func(context *gin.Context) {
		playground.Handler("GraphQL playground", graphqlEndpoint).ServeHTTP(context.Writer, context.Request)
	})

	r.POST(graphqlEndpoint, func(context *gin.Context) {
		resolver := &resolvers.Resolver{}
		handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})).ServeHTTP(context.Writer, context.Request)
	})

	log.Fatal(r.Run())
}
