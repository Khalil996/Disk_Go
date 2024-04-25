package admin

import (
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"fmt"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAdminStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAdminStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminStatusLogic {
	return &SetAdminStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAdminStatusLogic) SetAdminStatus(req *types.SetStatusReq) error {
	// todo: add your logic here and delete this line

	var err error

	defer mqs.LogSend(l.ctx, err, "SetAdminStatus", req.Id, req.Status)

	bean := &models.Admin{Status: req.Status}
	if _, err = l.svcCtx.Engine.ID(req.Id).Cols("status").Update(bean); err != nil {
		err = errors.New(fmt.Sprintf("SetAdminStatus，更新失败，ERR: [%v]", err.Error()))
		logx.Error(err)
		return err
	}

	return nil
}
