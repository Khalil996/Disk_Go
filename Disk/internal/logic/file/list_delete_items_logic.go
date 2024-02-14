package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeleteItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeleteItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeleteItemsLogic {
	return &ListDeleteItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeleteItemsLogic) ListDeleteItems() (resp *types.ListDeleteItemsRes, err error) {
	// todo: add your logic here and delete this line

	return
}
