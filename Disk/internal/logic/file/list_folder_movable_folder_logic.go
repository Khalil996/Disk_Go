package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFolderMovableFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFolderMovableFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFolderMovableFolderLogic {
	return &ListFolderMovableFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFolderMovableFolderLogic) ListFolderMovableFolder(req *types.ListFolderMovableFolderReq) (resp *types.ListFolderRes, err error) {
	// todo: add your logic here and delete this line

	return
}
