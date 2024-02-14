package user

import (
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailLogic) GetDetail(req *types.GetUserDetailReq) (resp *types.GetUserDetailRes, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GetUserDetailRes{}
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("id = ?", req.UserId).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	resp.Name = user.UserName
	resp.Email = user.Email
	resp.Phone = user.Phone
	resp.Avatar = user.Avatar
	resp.Gender = user.Gender

	return
}
