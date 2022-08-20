package gedis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

//RedisStore is used for interacting with redisClient
type RedisStore struct {
	Client *redis.Client
}

//NewRedisStore creates a new RedisStore
func NewRedisStore(redisUrl string) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	ctx := context.TODO()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Println("failed to ping to redis")
		panic(err)
	}

	return &RedisStore{
		Client: client,
	}
}
