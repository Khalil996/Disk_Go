package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Redis *Client

type (
	Conf struct {
		Addr     string
		Password string
		Db       int
	}

	Client struct {
		*redis.Client
	}
)

func Init(conf *Conf) *Client {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
	})
	//logx.Infof("[REDIS CONNECTING] InitRedis client: %v\n", client)

	if err := client.Ping(ctx).Err(); err != nil {
		panic("[REDIS ERROR] 连接redis失败 " + err.Error())
	}
	Redis = &Client{client}
	return &Client{client}
}
