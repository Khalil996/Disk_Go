package admin

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminInfoLogic {
	return &GetAdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminInfoLogic) GetAdminInfo() (interface{}, error) {
	// todo: add your logic here and delete this line

	adminId := l.ctx.Value(define.UserIdKey).(int64)
	var admin models.Admin
	if has, err := l.svcCtx.Engine.ID(adminId).Get(&admin); err != nil {
		logx.Errorf("GetAdminInfo,获取失败.Err: [%v]", err)
		return nil, err
	} else if !has {
		logx.Errorf("GetAdminInfo,获取失败.Err: [%v]", err)
		return nil, err
	}

	return admin, nil
}
