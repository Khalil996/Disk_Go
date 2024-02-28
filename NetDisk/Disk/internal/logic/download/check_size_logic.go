package download

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSizeLogic {
	return &CheckSizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSizeLogic) CheckSize(req *types.CheckSizeReq) (resp *types.CheckSizeResp, err error) {
	// todo: add your logic here and delete this line
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		file   models.File
		fileFs models.FileFs
	)
	has, err := l.svcCtx.Engine.ID(req.FileId).And("user_id=?", userId).Get(&file)
	if err != nil || !has {
		return nil, err
	}
	if file.IsBig == define.SmallFileFlag {
		resp.IsBig = file.IsBig
		return resp, nil
	}

	has, err = l.svcCtx.Engine.ID(file.FsId).Get(&fileFs)
	if err != nil || !has {
		return nil, err
	}
	resp.IsBig = file.IsBig
	resp.ChunkNum = fileFs.ChunkNum
	return
}
