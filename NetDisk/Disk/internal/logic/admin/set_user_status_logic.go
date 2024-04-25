package admin

import (
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserStatusLogic {
	return &SetUserStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserStatusLogic) SetUserStatus(req *types.SetStatusReq) error {
	// todo: add your logic here and delete this line
	var (
		err error
	)
	defer mqs.LogSend(l.ctx, err, "SetUserStatus", req.Id, req.Status)

	bean := &models.UserBasic{Status: req.Status}
	if _, err = l.svcCtx.Engine.ID(req.Id).
		Cols("status").
		Update(bean); err != nil {
		logx.Errorf("SetStatus，更新失败，ERR: [%v]", err)
		return err
	}

	return nil
}
