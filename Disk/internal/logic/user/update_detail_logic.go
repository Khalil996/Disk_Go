package user

import (
	"context"

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

func (l *UpdateDetailLogic) UpdateDetail(req *types.UpdateUserDetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
