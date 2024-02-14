package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFoldersLogic {
	return &CreateFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFoldersLogic) CreateFolders(req *types.CreateFoldersReq) error {
	// todo: add your logic here and delete this line

	return nil
}
