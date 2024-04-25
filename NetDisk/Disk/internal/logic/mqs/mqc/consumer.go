package mqc

import (
	"cloud_go/Disk/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *LogConsumer {
	return &LogConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogConsumer) Consume(key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	return nil
}
