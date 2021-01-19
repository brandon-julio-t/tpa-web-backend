package factories

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

type ApqCache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

func NewApqCache() *ApqCache {
	option, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(option)

	if err := client.Ping().Err(); err != nil {
		log.Fatal(err)
	}

	return &ApqCache{client: client, ttl: 24 * time.Hour}
}

func (c *ApqCache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *ApqCache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
