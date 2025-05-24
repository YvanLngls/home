package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisUserDb *redis.Client
var RedisDataDb *redis.Client

func InitRedis() {
	RedisUserDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisUserDb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	RedisDataDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	_, err = RedisDataDb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
