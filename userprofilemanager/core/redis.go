package core

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisSet(Key string, Value string) bool {
	redisHost := GetEnvrionment("REDIS_URI")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	err := client.Set(ctx, Key, Value, 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func RedisSetInterface(Key string, Value interface{}) bool {
	redisHost := GetEnvrionment("REDIS_URI")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	err := client.Set(ctx, Key, Value, 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func RedisSetWithExpire(Key string, Value string, Expire time.Duration) bool {
	redisHost := GetEnvrionment("REDIS_URI")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	err := client.Set(ctx, Key, Value, 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}

	err = client.Expire(ctx, Key, Expire).Err()
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func RedisAppend(Key string, Value string) bool {
	redisHost := GetEnvrionment("REDIS_URI")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	err := client.Append(ctx, Key, Value).Err()
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func RedisGet(Key string) string {
	redisHost := GetEnvrionment("REDIS_URI")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	val, err := client.Get(ctx, Key).Result()
	if err != nil {
		log.Println(err)
		return ""
	}

	return val

}
func RedisDel(Key string) bool {
	redisHost := GetEnvrionment("REDIS_URI")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	err := client.Del(ctx, Key).Err()
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func RedisExists(Key string) bool {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     GetEnvrionment("REDIS_URI"),
		Password: "",
		DB:       0,
	})

	_, err := client.Exists(ctx, Key).Result()
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
