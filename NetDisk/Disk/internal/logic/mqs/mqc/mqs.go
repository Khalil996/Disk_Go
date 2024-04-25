package mqc

import (
	"cloud_go/Disk/internal/config"
	"cloud_go/Disk/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//kq.MustNewQueue(c.KqConsumerConf, NewLogConsumer(ctx, svcContext)),
	}
}
