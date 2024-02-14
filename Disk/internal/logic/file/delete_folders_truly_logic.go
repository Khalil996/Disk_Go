package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFoldersTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoldersTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoldersTrulyLogic {
	return &DeleteFoldersTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoldersTrulyLogic) DeleteFoldersTruly(req *types.DeleteFoldersReq) error {
	// todo: add your logic here and delete this line

	return nil
}