package file

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverFilesLogic {
	return &RecoverFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverFilesLogic) RecoverFiles(req *types.RecoverFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
