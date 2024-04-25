package admin

import (
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetFileStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetFileStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetFileStatusLogic {
	return &SetFileStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetFileStatusLogic) SetFileStatus(req *types.SetFileStatusReq) error {
	// todo: add your logic here and delete this line
	var err error

	defer mqs.LogSend(l.ctx, err, "SetFileStatus", req.Ids, req.Status)

	bean := &models.File{Status: req.Status}
	if _, err = l.svcCtx.Engine.In("id", req.Ids).Cols("status").Update(bean); err != nil {
		logx.Errorf("SetFileStatus err:%v", err)
		return err
	}

	return nil
}
