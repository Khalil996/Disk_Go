package svc

import (
	"cloud_go/Disk/common/es"
	"cloud_go/Disk/common/minio"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/internal/config"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/internal/middleware"
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Minio  *minio.Client
	Auth   rest.Middleware
	Es     *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	idgenops := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(idgenops)

	minioClient := minio.Init(&c.Minio)
	engine := xorm.InitEngine(&c.Xorm)
	RDB := redis.Init(&c.Redis)
	esClient := es.Init(c.Eshost)
	initRedisData(RDB, c.Capacity)

	mqs.NewLogPusher(c.KqPusherConfs)
	return &ServiceContext{
		Config: c,
		Engine: engine,
		RDB:    RDB,
		Minio:  minioClient,
		Auth:   middleware.NewAuthMiddleware().Handle,
		Es:     esClient,
	}
}
func initRedisData(client *redis.Client, c uint64) {
	ctx := context.Background()
	key := redis.NetdiskCapacity
	if exist, err := client.Exists(ctx, key).Result(); err != nil {
		log.Fatalf("initRedisData，1，%v", err)
	} else if exist > 0 {
		return
	}

	if err := client.Set(ctx, key, c, 0).Err(); err != nil {
		log.Fatalf("initRedisData，2，%v", err)
	}
}
