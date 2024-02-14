package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFilesLogic {
	return &UpdateFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFilesLogic) UpdateFiles(req *types.UpdateFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
