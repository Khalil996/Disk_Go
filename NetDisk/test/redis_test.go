package test

import (
	"context"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestRedis(t *testing.T) {
	err := rdb.Set("name", "zhangsan", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
}

func TestGetKeys(t *testing.T) {
	err := rdb.Get("name").Err()
	if err != nil {
		panic(err)
	}
}
