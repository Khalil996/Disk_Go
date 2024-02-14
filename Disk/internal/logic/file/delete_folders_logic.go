package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoldersLogic {
	return &DeleteFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoldersLogic) DeleteFolders(req *types.DeleteFoldersReq) error {
	// todo: add your logic here and delete this line

	return nil
}
