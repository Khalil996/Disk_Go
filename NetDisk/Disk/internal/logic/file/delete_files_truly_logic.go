package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFilesTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesTrulyLogic {
	return &DeleteFilesTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesTrulyLogic) DeleteFilesTruly(req *types.IdsReq) error {
	// todo: add your logic here and delete this line
	var err error
	defer mqs.LogSend(l.ctx, err, "DeleteFilesTruly", req.Ids)
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var files []*models.File
	err = l.svcCtx.Engine.In("id", req.Ids).And("user_id=?", userId).
		And("del_flag=?", define.StatusFileDeleted).
		Find(&files)
	if err != nil {
		return err
	}

	length := len(req.Ids)
	if length != len(files) {
		return errors.New("发生错误：请求的 ID 数量与数据库中找到的不一致")
	}
	affected, err := l.svcCtx.Engine.In("id", req.Ids).Delete(&models.File{})
	if err != nil {
		err = errors.New("发生错误，" + err.Error())
		return err
	} else if affected != int64(length) {
		return errors.New("发生错误" + err.Error())
	}
	return nil
}
