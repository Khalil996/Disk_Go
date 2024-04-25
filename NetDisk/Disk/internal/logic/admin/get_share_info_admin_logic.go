package admin

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareInfoAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareInfoAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareInfoAdminLogic {
	return &GetShareInfoAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareInfoAdminLogic) GetShareInfoAdmin(req *types.IdStrReq) (*types.GetShareInfoResp, error) {
	// todo: add your logic here and delete this line
	var files []*models.File

	if err := l.svcCtx.Engine.Where(
		"id in (select file_id from share_file where share_id = ?)",
		req.Id).Find(&files); err != nil {
		logx.Errorf("GetShareInfoAdmin，查询失败，ERR: [%v]", err)
		return nil, err
	}

	resp := &types.GetShareInfoResp{}
	for _, file := range files {
		url, err := l.svcCtx.Minio.NewService().GenUrl(file.ObjectName, file.Name, true)
		if err != nil {
			logx.Errorf("GetShareInfoAdminLogic:GetShareInfoAdmin:GenUrl:err:%v", err)
			url = ""
		}
		resp.Items = append(resp.Items, &types.ListShareItemStruct{
			Id:      file.Id,
			Name:    file.Name,
			Url:     url,
			Size:    file.Size,
			Updated: file.Updated.Format(define.TimeFormat1),
			Status:  file.Status,
			Type:    file.Type,
			DelFlag: file.DelFlag,
		})
	}

	return resp, nil
}
