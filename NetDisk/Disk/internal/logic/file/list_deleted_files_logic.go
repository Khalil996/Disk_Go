package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"log"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ListDeletedFilesLogic 用于处理删除文件列表的逻辑
type ListDeletedFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewListDeletedFilesLogic 创建ListDeletedFilesLogic实例
func NewListDeletedFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeletedFilesLogic {
	return &ListDeletedFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListDeletedFiles 列出删除的文件
func (l *ListDeletedFilesLogic) ListDeletedFiles() (resp []*types.DeletedFilesResp, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var files []*models.File
	var folders []*models.Folder
	//查找文件
	if err = l.svcCtx.Engine.Where("user_id = ?", userId).
		And("del_flag = ?", define.StatusFileDeleted).Asc("del_time").
		Find(&files); err != nil {
		return nil, err
	}

	log.Println(len(files), userId)

	var folderIds []int64
	m := make(map[int64]string)
	for _, file := range files {
		if _, ok := m[file.FolderId]; !ok {
			folderIds = append(folderIds, file.FolderId)
		} else {
			m[file.FolderId] = ""
		}
	}

	if err = l.svcCtx.Engine.Select("id, name").In("id", folderIds).
		//And("del_flag = ?", define.StatusFolderUndeleted).
		Find(&folders); err != nil {
		return nil, err
	}

	m = map[int64]string{}
	for _, folder := range folders {
		m[folder.Id] = folder.Name
	}

	for _, file := range files {
		resp = append(resp, &types.DeletedFilesResp{
			Id:         file.Id,
			Name:       file.Name,
			Status:     file.Status,
			Size:       file.Size,
			FolderId:   file.FolderId,
			FolderName: m[file.FolderId],
			DelTime:    file.DelTime,
		})
	}
	return
}
