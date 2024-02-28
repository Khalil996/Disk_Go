package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesLogic {
	return &DeleteFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesLogic) DeleteFiles(req *types.IdsReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	cond := &models.File{
		DelFlag: define.StatusFileDeleted,
		Model:   models.Model{DeleteAt: time.Now().Local().UTC()},
	}
	affected, err := l.svcCtx.Engine.In("id", req.Ids).And("user_id = ?", userId).Update(cond)
	if err != nil || affected != int64(len(req.Ids)) {
		return errors.New("删除文件失败" + err.Error())
	}
	return nil
}
