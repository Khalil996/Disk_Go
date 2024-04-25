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

type MoveFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFoldersLogic {
	return &MoveFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFoldersLogic) MoveFolders(req *types.MoveFoldersReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	parentFolderId := req.ParentFolderId
	var (
		err    error
		folder = &models.Folder{}
	)

	defer mqs.LogSend(l.ctx, err, "MoveFolders", req.FolderIds, req.ParentFolderId)
	if parentFolderId != 0 {
		has, err := l.svcCtx.Engine.ID(parentFolderId).And("user_id = ?", userId).Get(folder)
		if err != nil {
			return errors.New("发生错误！")
		} else if !has {
			return errors.New("该目录不存在")
		}
	}

	folder = &models.Folder{ParentId: parentFolderId}
	affected, err := l.svcCtx.Engine.In("id", req.FolderIds).Update(folder)
	if err != nil || affected != int64(len(req.FolderIds)) {
		return errors.New("" + err.Error() + "发生错误！")
	}
	return nil
}
