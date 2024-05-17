package admin

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	StatisticLogic struct {
		logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	TypeCnt struct {
		Type int8  `json:"type"`
		Cnt  int64 `json:"cnt"`
	}
)

func NewStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticLogic {
	return &StatisticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticLogic) Statistic() (map[string]interface{}, error) {
	// todo: add your logic here and delete this line
	var (
		resp   map[string]interface{}
		user   = &models.UserBasic{}
		share  = &models.Share{}
		file   = &models.File{}
		folder = &models.Folder{}
		fs     = &models.FileFs{}
	)
	userCnt, err := l.svcCtx.Engine.Count(user)
	if err != nil {
		return nil, err
	}
	used, err := l.svcCtx.Engine.SumInt(user, "used")
	if err != nil {
		return nil, err
	}
	capacity, err := l.svcCtx.RDB.Get(l.ctx, redis.NetdiskCapacity).Result()
	if err != nil {
		return nil, err
	}
	shareSums, err := l.svcCtx.Engine.SumsInt(share, "download_num", "click_num")
	if err != nil {
		return nil, err
	}
	var typeCnt []*TypeCnt
	if err = l.svcCtx.Engine.Select("type,count(*) as cnt").GroupBy("type").Table(file).Find(&typeCnt); err != nil {
		return nil, err
	}
	fileCnt, err := l.svcCtx.Engine.Count(file)
	if err != nil {
		return nil, err
	}
	folderCnt, err := l.svcCtx.Engine.Count(folder)
	if err != nil {
		return nil, err
	}
	fsCnt, err := l.svcCtx.Engine.Count(fs)
	if err != nil {
		return nil, err
	}

	resp = map[string]interface{}{
		"cnt": map[string]interface{}{
			"user":     userCnt,
			"download": shareSums[0],
			"click":    shareSums[1],
			"file": map[string]interface{}{
				"file":   fileCnt,
				"type":   typeCnt,
				"folder": folderCnt,
				"fs":     fsCnt,
			},
		},
		"cap": map[string]interface{}{
			"used":     used,
			"capacity": capacity,
		},
	}

	return resp, nil
}
