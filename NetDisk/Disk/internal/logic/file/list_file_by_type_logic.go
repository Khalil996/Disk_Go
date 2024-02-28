package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileByTypeLogic {
	return &ListFileByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileByTypeLogic) ListFileByType(req *types.FileTypeReq) (resp []*types.FileRes, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var files []*models.File

	if err := l.svcCtx.Engine.Where("type = ?", req.FileType).
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
	return
}
