package main

import (
	"bytes"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/brandon-julio-t/tpa-web-backend/bootstrap"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/factories"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/graph/resolvers"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	r := gin.Default()

	r.Use(
		cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			AllowOrigins:     []string{"http://localhost:4200", "https://tpa-web-br20-2.netlify.app"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		new(middlewares.GinProviderMiddleware).Create(),
		new(middlewares.AuthProviderMiddleware).Create(),
	)

	r.GET("/", func(context *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(context.Writer, context.Request)
	})

	r.GET("/graphql/assets/:id", func(context *gin.Context) {
		sid := context.Param("id")

		id, err := strconv.ParseInt(sid, 10, 64)
		if err != nil {
			log.Print(err)
			context.Status(http.StatusBadRequest)
		}

		var file models.AssetFile
		if err := facades.UseDB().First(&file, id).Error; err != nil {
			log.Print(err)
			context.Status(http.StatusNotFound)
		}

		reader := bytes.NewReader(file.File)

		context.DataFromReader(http.StatusOK, int64(reader.Len()), file.ContentType, reader, map[string]string{})
	})

	gql := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolvers.NewResolver(),
	}))

	gql.Use(extension.AutomaticPersistedQuery{
		Cache: factories.NewApqCache(),
	})

	gql.AddTransport(transport.POST{})
	gql.AddTransport(transport.MultipartForm{})
	gql.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	gql.Use(extension.Introspection{})
	gql.Use(apollotracing.Tracer{})

	r.Any("/graphql", func(context *gin.Context) {
		gql.ServeHTTP(context.Writer, context.Request)
	})

	log.Fatal(r.Run())
}
