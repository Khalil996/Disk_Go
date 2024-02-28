package file

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesLogic {
	return &DeleteFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesLogic) DeleteFiles(req *types.DeleteFilesReq) error {
	// todo: add your logic here and delete this line
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
		rdb    = l.svcCtx.RDB
	)

	bean := &models.File{
		DelFlag: define.StatusFileDeleted,
		DelTime: time.Now().Local().Unix(),
	}
	if affected, err := engine.In("id", req.FileIds).
		And("user_id = ?", userId).Update(bean); err != nil {
		logx.Errorf("删除文件失败，ERR: [%v]", err)
		return errors.New("删除过程出错，" + err.Error())
	} else if affected != int64(len(req.FileIds)) {
		return errors.New("删除过程出错！")
	}

	key := fmt.Sprintf(redis.FileFolderDownloadUrlKey, userId, req.FolderId)
	if err := rdb.ZRem(l.ctx, key, req.FileIds).Err(); err != nil {
		if err != redis2.Nil {
			logx.Errorf("删除文件更新redis缓存失败，ERR: [%v]", err)
		}
	}

	return nil
}
