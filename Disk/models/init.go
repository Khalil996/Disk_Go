package models

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/disk_go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return engine
}
