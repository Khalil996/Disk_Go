package config

import (
	"cloud_go/Disk/common/minio"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Minio    minio.Conf
	Xorm     xorm.DbConf
	Redis    redis.Conf
	Eshost   string
	Capacity uint64
	Idgen    struct {
		WorkerId uint16
	}
	KqPusherConfs []*KqPusherConf
}
type KqPusherConf struct {
	Type    string
	Brokers []string
	Topic   string
}
