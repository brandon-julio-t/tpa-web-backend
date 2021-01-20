module github.com/brandon-julio-t/tpa-web-backend

// +heroku goVersion go1.15

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/bxcodec/faker/v3 v3.5.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.4.9
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/google/uuid v1.1.4
	github.com/gorilla/websocket v1.4.2
	github.com/joho/godotenv v1.3.0
	github.com/mailjet/mailjet-apiv3-go v0.0.0-20201009050126-c24bc15a9394
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.11
	syreclabs.com/go/faker v1.2.3 // indirect
)
