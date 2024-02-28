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

func (l *GetFileDetailLogic) GetFileDetail(req *types.IdPathReq) (*types.FileResp, error) {
	// todo: add your logic here and delete this line
	var (
		userId   = l.ctx.Value(define.UserIdKey).(int64)
		engine   = l.svcCtx.Engine
		minioSvc = l.svcCtx.Minio.NewService()
		file     = &models.File{}
	)

	has, err := engine.ID(req.Id).And("user_id = ?", userId).Get(file)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("未能找到该文件信息！😿")
	}

	url, err := minioSvc.GenUrl(file.ObjectName, false)
	if err != nil {
		return nil, err
	}

	resp := &types.FileResp{}
	resp.Id = file.Id
	resp.Name = file.Name
	resp.Url = url
	resp.Ext = file.Ext
	resp.Size = file.Size
	resp.Status = file.Status
	resp.FolderId = file.FolderId
	resp.Created = file.Created.Format(define.TimeFormat1)
	resp.Updated = file.Updated.Format(define.TimeFormat1)
	return resp, nil
}
