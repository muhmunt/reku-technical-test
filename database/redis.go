package db

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fmt.Println("Redis Connected.")
	return redisCli
}
