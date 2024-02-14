package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFoldersLogic {
	return &UpdateFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFoldersLogic) UpdateFolders(req *types.UpdateFoldersReq) error {
	// todo: add your logic here and delete this line

	return nil
}
