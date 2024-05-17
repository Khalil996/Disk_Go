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

type CopyFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyFilesLogic {
	return &CopyFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyFilesLogic) CopyFiles(req *types.CopyFilesReq) error {
	// todo: add your logic here and delete this line
	UserId := l.ctx.Value(define.UserIdKey).(int64)
	folderId := req.FolderId
	var err error
	defer mqs.LogSend(l.ctx, err, "CopyFiles", req.FileIds)

	var files []*models.File
	if folderId != 0 {
		has, err := l.svcCtx.Engine.ID(folderId).And("user_id = ?", UserId).Get(&models.Folder{})
		if err != nil {
			return errors.New("发生错误" + err.Error())
		} else if !has {
			return errors.New("文件夹不存在")
		}
	}
	err = l.svcCtx.Engine.In("id", req.FileIds).Find(&files)
	if err != nil {
		return errors.New("发生错误" + err.Error())
	}
	now := time.Now()
	for _, file := range files {
		file.Id = idgen.NextId()
		file.Name = file.Name + "_" + time.Now().Format(define.TimeFormat2) + "复制" + file.Ext
		file.Updated = now
		file.Created = now
		file.FolderId = req.FolderId
	}
	affected, err := l.svcCtx.Engine.Insert(files)
	if err != nil || affected != int64(len(files)) {
		return errors.New("发生错误" + err.Error())
	}
	return nil
}
