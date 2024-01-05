package svc

import (
	"cloud_go/Disk/internal/config"
	"cloud_go/Disk/internal/middleware"
	"cloud_go/Disk/models"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Engine,
		RDB:    models.RDB,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
