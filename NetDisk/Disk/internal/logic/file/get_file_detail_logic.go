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

type GetFileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileDetailLogic {
	return &GetFileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileDetailLogic) GetFileDetail(req *types.IdPathReq) (resp *types.FileRes, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var file = &models.File{}
	has, err := l.svcCtx.Engine.ID(req.Id).And("user_id =?", userId).Get(file)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("文件不存在")
	}
	resp = &types.FileRes{}
	resp.Id = file.Id
	resp.Name = file.Name
	resp.Status = file.Status
	resp.Size = file.Size
	resp.FolderId = file.FolderId
	resp.Url = file.Url
	resp.Created = file.CreateAt.Format(define.TimeFormat1)
	resp.Updated = file.UpdateAt.Format(define.TimeFormat1)
	return
}
