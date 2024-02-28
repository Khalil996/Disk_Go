package user

import (
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDetailLogic {
	return &UpdateDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDetailLogic) UpdateDetail(req *types.UpdateUserDetailReq) (resp *types.UpdateUserDetailRes, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserBasic)
	sign, err := l.svcCtx.Engine.Where("id=?", req.UserId).Get(user)
	if err != nil {
		return nil, err
	}
	if !sign {
		return nil, errors.New("user not found")
	}
	user.Name = req.Name
	user.Gender = req.Gender
	_, err = l.svcCtx.Engine.ID(req.UserId).Update(user)
	if err != nil {
		return nil, err
	}
	resp.Name = req.Name
	resp.Gender = req.Gender
	return
}
