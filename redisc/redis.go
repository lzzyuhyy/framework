package redisc

import (
	"encoding/json"
	"framework/consul"
	"github.com/go-redis/redis"
	"time"
)

type redisConf struct {
	Addr     string
	Password string
	DB       int
}

func WithRedisClient(handler func(cli *redis.Client) error) error {
	info, err := consul.GetKeyInfo("redis")
	if err != nil {
		return err
	}
	var r redisConf
	err = json.Unmarshal(info, &r)
	if err != nil {
		return err
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})

	defer cli.Close()

	return handler(cli)
}

func RedisLock(key string, val any, dur time.Duration) error {
	return WithRedisClient(func(cli *redis.Client) error {
		for {
			ok, err := cli.SetNX(key, val, dur).Result()
			if err != nil {
				return err
			}

			if ok {
				return nil
			}
		}
	})
}

func DelRedisLock(key string) error {
	return WithRedisClient(func(cli *redis.Client) error {
		_, err := cli.Del(key).Result()
		return err
	})
}
