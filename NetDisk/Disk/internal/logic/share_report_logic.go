package logic

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareReportLogic {
	return &ShareReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareReportLogic) ShareReport(req *types.ReportReq) error {
	// todo: add your logic here and delete this line
	var err error

	defer mqs.LogSend(l.ctx, err, "ShareReport", req.ShareId)

	bean := &models.Share{
		Status: define.StatusShareIllegal,
		Reason: req.Reason,
	}
	if _, err = l.svcCtx.Engine.ID(req.ShareId).Update(bean); err != nil {
		err = errors.New("举报失败")
		logx.Error(err)
		return err
	}

	return nil
}
