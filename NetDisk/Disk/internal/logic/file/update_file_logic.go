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

type UpdateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFileLogic) UpdateFile(req *types.UpdateNameReq) error {
	// todo: add your logic here and delete this line

	userId := l.ctx.Value(define.UserIdKey).(int64)

	affected, err := l.svcCtx.Engine.ID(req.Id).And("user_id=?", userId).Update(&models.File{Name: req.Name})
	if err != nil {
		return err
	} else if affected != 1 {
		return errors.New("更新失败")
	}
	return nil
}
