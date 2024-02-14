package config

import (
	"cloud_go/Disk/common/minio"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Minio minio.Conf
}
