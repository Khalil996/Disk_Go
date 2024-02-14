package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileMovableFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileMovableFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileMovableFolderLogic {
	return &ListFileMovableFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileMovableFolderLogic) ListFileMovableFolder(req *types.ParentFolderIdReq) (resp *types.ListFolderRes, err error) {
	// todo: add your logic here and delete this line

	return
}
