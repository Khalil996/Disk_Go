package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListfileByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListfileByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListfileByTypeLogic {
	return &ListfileByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListfileByTypeLogic) ListfileByType(req *types.FileTypeReq) (resp *types.FileRes, err error) {
	// todo: add your logic here and delete this line

	return
}
