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

type ListDeleteFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeleteFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeleteFilesLogic {
	return &ListDeleteFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeleteFilesLogic) ListDeleteFiles() (resp []*types.DeletedFilesResp, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var files []*models.File
	var folders []*models.Folder
	if err = l.svcCtx.Engine.Select("id, name, url, delete_at").Where("user_id = ?", userId).
		And("del_flag = ?", define.StatusFileDeleted).Asc("delete_at").
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
		And("del_flag = ?", define.StatusFolderUndeleted).
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
			Url:        file.Url,
			Status:     file.Status,
			Size:       file.Size,
			FolderId:   file.FolderId,
			FolderName: m[file.FolderId],
			DelTime:    file.DeleteAt.Format(define.TimeFormat1),
		})
	}
	return
}
