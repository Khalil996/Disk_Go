package test

import (
	"cloud_go/Disk/models"
	"testing"
	"xorm.io/xorm"
)

func TestMysqlTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:Khalil996!@tcp(0.0.0.0:3306)/disk_go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	err = engine.CreateTables(models.UserBasic{})
	if err != nil {
		panic(err)
	}
	err = engine.CreateTables(models.File{})
	if err != nil {
		panic(err)
	}
}
