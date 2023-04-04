package redis

import (
	"github.com/go-redis/redis"
)

const (
	RedisAuth = iota
)

func InitRedis(addr, passwd string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     addr,
		Password: passwd,
		PoolSize: 100,
	})
	_, err := client.Ping().Result()
	if err == nil {
		//TODO: 清除所有的redis token
		//client.FlushDB()
	}
	return client, err
}
