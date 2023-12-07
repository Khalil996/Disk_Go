package svc

import (
	"cloud_go/Disk/internal/config"
	"cloud_go/Disk/models"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Engine,
		RDB:    models.RDB,
	}
}
