package facades

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var redisInstance redis.UniversalClient

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	option, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	redisInstance = redis.NewClient(option)
}

func UseCache() redis.UniversalClient {
	return redisInstance
}
