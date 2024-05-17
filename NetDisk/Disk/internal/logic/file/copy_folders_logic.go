package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyFoldersLogic {
	return &CopyFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyFoldersLogic) CopyFolders(req *types.CopyFoldersReq) error {
	// todo: add your logic here and delete this line
	var (
		userId  = l.ctx.Value(define.UserIdKey).(int64)
		folders []*models.Folder
		err     error
	)
	defer mqs.LogSend(l.ctx, err, "copyFolders", req.ParentFolderId)

	if req.ParentFolderId != 0 {
		folder := &models.Folder{}
		if _, err = l.svcCtx.Engine.ID(req.ParentFolderId).And("user_id=?", userId).Get(folder); err != nil {
			err = errors.New("发生错误！" + err.Error())
			return err
		}
		if folder.Id == 0 {
			return errors.New("该目录不存在")
		}
	}

	if err = l.svcCtx.Engine.In("id", req.FolderIds).Find(&folders); err != nil {
		err = errors.New("发生错误！" + err.Error())
		return err
	}

	for _, folder := range folders {
		folder.Id = idgen.NextId()
		folder.Name = folder.Name + "_" + time.Now().Format(define.TimeFormat2) + "复制"
		folder.ParentId = req.ParentFolderId
	}
	if _, err = l.svcCtx.Engine.Insert(&folders); err != nil {
		return errors.New("发生错误！..." + err.Error())
	}

	return nil
}
