package models

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = Init()
var RDB = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:Khalil996!@tcp(127.0.0.1:3306)/disk_go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return engine
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
