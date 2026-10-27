package redis

import (
	"context"
	"log"

	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{Addr: addr})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("redis client created")

	return rdb
}
