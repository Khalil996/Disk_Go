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

type UpdateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFolderLogic {
	return &UpdateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFolderLogic) UpdateFolder(req *types.UpdateNameReq) error {
	// todo: add your logic here and delete this line

	userId := l.ctx.Value(define.UserIdKey).(int64)
	var err error
	defer mqs.LogSend(l.ctx, err, "UpdateFolder", req.Id, req.Name)
	affect, err := l.svcCtx.Engine.ID(req.Id).And("user_id=?", userId).Update(&models.Folder{Name: req.Name})
	if err != nil {
		return err
	} else if affect != 1 {
		return errors.New("更新文件夹失败")
	}
	return nil
}
