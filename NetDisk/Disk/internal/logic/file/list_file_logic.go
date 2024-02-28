package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileLogic {
	return &ListFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileLogic) ListFile(req *types.ParentFolderIdReq) ([]*types.FileRes, error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var files []*models.File
	var resp []*types.FileRes
	if err := l.svcCtx.Engine.Where("folder_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).And("del_flag = ?",
		define.StatusFileUndeleted).Find(&files); err != nil {
		return nil, err
	}
	for _, file := range files {
		resp = append(resp, &types.FileRes{
			Id:      file.Id,
			Name:    file.Name,
			Url:     file.Url,
			Size:    file.Size,
			Status:  file.Status,
			Updated: file.UpdateAt.Format(define.TimeFormat1),
		})
	}
	return resp, nil
}
