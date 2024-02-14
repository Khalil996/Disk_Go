package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFilesTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesTrulyLogic {
	return &DeleteFilesTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesTrulyLogic) DeleteFilesTruly(req *types.DeleteFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
