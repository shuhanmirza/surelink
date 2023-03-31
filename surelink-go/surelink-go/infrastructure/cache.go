package infrastructure

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Cache struct {
	Client *redis.Client
}

func NewCache(redisUrl string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	ctx := context.TODO()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Println("failed to ping to redis")
		panic(err)
	}

	return &Cache{
		Client: client,
	}
}
