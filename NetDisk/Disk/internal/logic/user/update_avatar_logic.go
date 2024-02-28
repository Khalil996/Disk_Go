package user

import (
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(req *types.UpdateAvatarReq) (resp *types.UpdateAvatarRes, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserBasic)
	ava, err := l.svcCtx.Engine.Where("id=?", req.UserId).Get(user)
	if err != nil {
		return nil, err
	}
	if !ava {
		return nil, errors.New("用户不存在")
	}
	user.Avatar = req.Avatar

	_, err = l.svcCtx.Engine.Where("id=?", req.UserId).Cols("avatar").Update(user)
	if err != nil {
		return nil, err
	}
	resp.Avatar = req.Avatar
	return
}
