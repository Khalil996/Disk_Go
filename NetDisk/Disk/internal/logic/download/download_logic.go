package download

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadReq) (string, error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var file models.File
	var fileFs models.FileFs

	if has, err := l.svcCtx.Engine.ID(req.FileId).
		And("user_id = ?", userId).
		Get(&file); err != nil || !has {
		return "", err
	}

	if has, err := l.svcCtx.Engine.ID(file.FsId).
		Get(&fileFs); err != nil || !has {
		return "", err
	}

	return "", nil
}
