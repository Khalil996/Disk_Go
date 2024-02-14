package svc

import (
	"cloud_go/Disk/common/minio"
	redis2 "cloud_go/Disk/common/redis"
	xorm2 "cloud_go/Disk/common/xorm"
	"cloud_go/Disk/internal/config"
	"cloud_go/Disk/internal/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Minio  *minio.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient := minio.Init(&c.Minio)

	return &ServiceContext{
		Config: c,
		Engine: xorm2.Engine,
		RDB:    redis2.RDB,
		Minio:  minioClient,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
