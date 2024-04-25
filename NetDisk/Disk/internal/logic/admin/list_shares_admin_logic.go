package admin

import (
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSharesAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSharesAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSharesAdminLogic {
	return &ListSharesAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSharesAdminLogic) ListSharesAdmin(req *types.PageReq) ([]*models.Share, error) {
	// todo: add your logic here and delete this line
	var share []*models.Share
	offset := int((req.Page - 1) * req.Size)
	if err := l.svcCtx.Engine.Limit(int(req.Size), offset).Find(&share); err != nil {
		logx.Errorf("获取分享列表，查询shares失败，ERR: [%v]", err)
		return nil, errors.New("获取分享列表失败")
	}
	return share, nil
}
