package middleware

import (
	"fmt"
	"testing"
	"time"

	"my-admin/pkg/redis"
)

func TestAuth(t *testing.T) {
	token, err := GenerateToken("test", "123456")
	fmt.Printf("token:%s, err:%+v\n", token, err)

	result, err := parseToken(token)
	fmt.Printf("claim:%+v, err:%+v\n", result, err)
}

func TestStoreToken(t *testing.T) {
	token, err := GenerateToken("test", "123456")
	fmt.Printf("token:%s, err:%+v\n", token, err)
	c, err := redis.InitRedis("redis.test.com:6379", "")
	if err != nil {
		fmt.Printf("failed to init redis client, err:%+v\n", err)
		return
	}
	err = c.Set(token, nil, time.Minute*1).Err()
	if err != nil {
		fmt.Printf("failed to set token, token info:%s, err:%+v\n", token, err)
		return
	}

	_, err = c.Get(token).Result()
	if err != nil {
		fmt.Printf("token is not exist, err:%+v\n", err)
	}

	time.Sleep(15 * time.Second)
	err = c.Del(token).Err()
	err = c.Set(token, nil, time.Second*120).Err()
	if err != nil {
		fmt.Printf("failed update token, err:%+v\n", err)
	}

	result, err := parseToken(token)
	fmt.Printf("claim:%+v, err:%+v\n", result, err)

	time.Sleep(121 * time.Second)

	_, err = c.Get(token).Result()
	if err != nil {
		fmt.Printf("token is not exist, err:%+v\n", err)
	}

}
