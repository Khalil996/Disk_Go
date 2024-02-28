package config

import (
	"cloud_go/Disk/common/minio"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Minio minio.Conf
	Redis redis.Conf
	Xorm  xorm.DbConf
	Idgen struct {
		WorkerId uint16
	}
}
