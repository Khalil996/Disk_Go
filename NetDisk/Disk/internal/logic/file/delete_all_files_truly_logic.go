package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAllFilesTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAllFilesTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAllFilesTrulyLogic {
	return &DeleteAllFilesTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAllFilesTrulyLogic) DeleteAllFilesTruly() error {
	// todo: add your logic here and delete this line
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
		files  []*models.File
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "DeleteAllFilesTruly")

	if err := engine.Where("user_id = ?", userId).
		And("del_flag = ?", define.StatusFileDeleted).
		Find(&files); err != nil {
		return errors.New("出错啦！，" + err.Error())
	}

	if affected, err := engine.Where("user_id = ?", userId).
		And("del_flag = ?", define.StatusFileDeleted).
		Delete(&models.File{}); err != nil {
		return errors.New("删除过程出错啦，" + err.Error())
	} else if affected != int64(len(files)) {
		return errors.New("删除过程出错啦！")
	}

	// TODO: MQ
	go l.fs(files)
	return nil
}
func (l *DeleteAllFilesTrulyLogic) fs(files []*models.File) {

}
