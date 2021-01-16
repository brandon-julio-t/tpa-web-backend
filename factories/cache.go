package factories

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

func NewCache(redisUrl string, ttl time.Duration) *Cache {
	option, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(option)

	if err := client.Ping().Err(); err != nil {
		log.Fatal(err)
	}

	return &Cache{client: client, ttl: ttl}
}

func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
