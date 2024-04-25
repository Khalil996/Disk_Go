package admin

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

type SetShareStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetShareStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetShareStatusLogic {
	return &SetShareStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetShareStatusLogic) SetShareStatus(req *types.SetShareStatusReq) error {
	// todo: add your logic here and delete this line

	var err error

	defer mqs.LogSend(l.ctx, err, "SetShareStatus", req.Id, req.Status)
	_, err = l.svcCtx.Engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if req.Type != define.TypeShareMulti {
			bean := &models.File{Status: define.StatusFileIllegal}
			if _, err2 := session.Where("id="+"(select file_id from share_file where share_id =?)", req.Id).Update(bean); err2 != nil {
				logx.Errorf("SetShareStatus err:%v", err2)
				err = err2
				return nil, err
			}
		}
		bean := &models.Share{Status: req.Status}
		if _, err2 := session.ID(req.Id).Cols("status").Update(bean); err2 != nil {
			logx.Errorf("SetShareStatus err:%v", err2)
			err = err2
			return nil, err

		}
		return nil, nil
	})

	return err
}
