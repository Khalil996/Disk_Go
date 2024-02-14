package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyFilesLogic {
	return &CopyFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyFilesLogic) CopyFiles(req *types.CopyFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
