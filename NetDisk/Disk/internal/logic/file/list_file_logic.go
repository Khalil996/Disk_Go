package file

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileLogic {
	return &ListFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileLogic) ListFile(req *types.ParentFolderIdReq) ([]*types.FileResp, error) {
	// todo: add your logic here and delete this line
	var (
		userId   = l.ctx.Value(define.UserIdKey).(int64)
		files    []*models.File
		resp     []*types.FileResp
		engine   = l.svcCtx.Engine
		minioSvc = l.svcCtx.Minio.NewService()
		rdb      = l.svcCtx.RDB
		key      = fmt.Sprintf(redis.FileFolderDownloadUrlKey, userId, req.ParentFolderId)
	)

	if err := engine.Desc("created").
		Select("id, name, size, object_name, type, status, created, updated").
		Where("folder_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).
		And("del_flag = ?", define.StatusFileUndeleted).
		Find(&files); err != nil {
		return nil, err
	}

	zs, redisErr := rdb.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if redisErr != nil && redisErr != redis2.Nil {
		logx.Errorf("通过文件夹id获取文件列表，redis获取set失败，ERR: [%v]", redisErr)
	}

	var urls []redis2.Z
	for i, file := range files {
		var url string
		if len(zs) == len(files) && redisErr == nil {
			url = zs[i].Member.(string)
		} else {
			url2, err := minioSvc.GenUrl(file.ObjectName, true)
			if err != nil {
				logx.Errorf("通过文件夹id获取文件列表，[%d]获取url失败，ERR: [%v]", file.Id, err)
			} else {
				url = url2
				urls = append(urls, redis2.Z{Member: url, Score: float64(file.Created.Unix())})
				if i == len(files)-1 {
					if err = rdb.ZAdd(l.ctx, key, urls...).Err(); err != nil {
						logx.Errorf("通过文件夹id获取文件列表，redis缓存url失败，ERR: [%v]", err)
					}
					if err = rdb.Expire(l.ctx, key, 7*24*time.Hour).Err(); err != nil {
						logx.Errorf("通过文件夹id获取文件列表，设置缓存expire失败，ERR: [%v]", err)
					}
				}
			}
		}

		resp = append(resp, &types.FileResp{
			Id:      file.Id,
			Name:    file.Name,
			Url:     url,
			Type:    file.Type,
			Size:    file.Size,
			Status:  file.Status,
			Updated: file.Updated.Format(define.TimeFormat1),
		})
	}
	return resp, nil
}
