package file

import (
	"cloud_go/Disk/define"
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
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var files []models.File
	err := l.svcCtx.Engine.Cols("id").In("id", req).And("user_id=?", userId).And("del_flag=?", define.StatusFileDeleted).Find(&files)
	if err != nil {
		return err
	}
	length := len(req.Ids)
	if length != len(files) {
		return errors.New("发生错误")
	}
	affected, err := l.svcCtx.Engine.In("id", req.Ids).Delete(&models.File{})
	if err != nil || affected != int64(length) {
		return errors.New("发生错误" + err.Error())
	}
	return nil
}
