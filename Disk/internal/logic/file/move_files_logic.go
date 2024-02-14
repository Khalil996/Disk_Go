package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFilesLogic {
	return &MoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFilesLogic) MoveFiles(req *types.MoveFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
