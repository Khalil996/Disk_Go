package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSharesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSharesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSharesLogic {
	return &ListSharesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSharesLogic) ListShares() (resp []*types.ListShareStruct, err error) {
	// todo: add your logic here and delete this line
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		shares []*models.Share
	)
	if err := l.svcCtx.Engine.Where("user_id = ?", userId).Find(&shares); err != nil {
		return nil, errors.New("获取分享列表失败")
	}
	// 检查分享是否过期
	var expiredShares []string
	for _, share := range shares {
		status := share.Status
		// 如果分享未过期，检查是否在有效期内
		if share.Status == define.StatusShareNotExpired && time.Now().Unix()+10 > share.Expired {
			expiredShares = append(expiredShares, share.Id)
			status = define.StatusShareExpired
		}
		resp = append(resp, &types.ListShareStruct{
			Id:          share.Id,
			Pwd:         share.Pwd,
			Name:        share.Name,
			Created:     share.Created.Format(define.TimeFormat1),
			Expired:     share.Expired,
			Status:      status,
			DownloadNum: share.DownloadNum,
			ClickNum:    share.ClickNum,
			Type:        share.Type,
			Url:         share.Url,
		})
	}
	// 更新分享状态
	if len(expiredShares) > 0 {
		if _, err = l.svcCtx.Engine.In("id", expiredShares).
			Update(&models.Share{Status: define.StatusShareExpired}); err != nil {
			logx.Errorf("更新分享状态失败:%v", err)
		}
	}

	return
}
