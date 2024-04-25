package admin

import (
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUsersLogic) ListUsers(req *types.PageReq) ([]*models.UserBasic, error) {
	// todo: add your logic here and delete this line
	var users []*models.UserBasic

	offset := int((req.Page - 1) * req.Size)
	if err := l.svcCtx.Engine.Limit(int(req.Size), offset).Find(&users); err != nil {
		logx.Errorf("ListUsersLogic.ListUsers err:%v", err)
		return nil, err
	}

	for _, user := range users {
		url, err := l.svcCtx.Minio.NewService().GenUrl(user.Avatar, "", false)
		if err != nil {
			continue
		}
		user.Avatar = url
	}

	return users, nil
}
