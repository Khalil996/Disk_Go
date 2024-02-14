package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesLogic {
	return &DeleteFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesLogic) DeleteFiles(req *types.DeleteFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
