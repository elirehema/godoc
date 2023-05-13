package redis

import (
	"fmt"

	"log"

	"github.com/go-redis/redis"
)

var RedisClient = initRedis()

func initRedis() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       10,
	})

	// test connection

	pong, err := c.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
	return c
}
