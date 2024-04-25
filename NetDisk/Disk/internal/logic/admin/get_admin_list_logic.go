package admin

import (
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminListLogic {
	return &GetAdminListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminListLogic) GetAdminList(req *types.PageReq) (interface{}, error) {
	// todo: add your logic here and delete this line

	var admin []*models.Admin

	// 分页查询
	offset := int((req.Page - 1) * req.Size)
	if err := l.svcCtx.Engine.Limit(int(req.Size), offset).Find(&admin); err != nil {
		logx.Errorf("GetAdminListLogic.GetAdminList err:%v", err)
		return nil, err
	}

	return admin, nil
}
