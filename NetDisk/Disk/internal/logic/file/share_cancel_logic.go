package file

import (
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareCancelLogic {
	return &ShareCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareCancelLogic) ShareCancel(req *types.IdStrsReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var err error
	// 记录日志
	defer mqs.LogSend(l.ctx, err, "ShareCancel", req.Ids)
	// 删除分享
	_, err = l.svcCtx.Engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		// 先删除 Share
		if _, err := session.In("id", req.Ids).Delete(&models.Share{UserId: userId}); err != nil {
			logx.Errorf("删除分享失败：%v", err)
			return nil, err
		}
		// 再删除 ShareFile
		if _, err := session.In("share_id", req.Ids).Delete(&models.ShareFile{}); err != nil {
			logx.Errorf("删除分享记录失败：%v", err)
			return nil, err
		}
		return nil, nil

	})
	return err
}
