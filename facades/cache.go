package facades

import (
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var redisInstance redis.UniversalClient

func getCache() redis.UniversalClient {
	if redisInstance == nil {
		option, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			log.Fatal(err)
		}

		redisInstance = redis.NewClient(option)
	}
	return redisInstance
}

func UseCache() redis.UniversalClient {
	return getCache()
}
