package upload

import (
	"cloud_go/Disk/common/redis"
	"context"
	"fmt"
	"strconv"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckChunkLogic {
	return &CheckChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 检查分块是否上传过
func (l *CheckChunkLogic) CheckChunk(req *types.CheckChunkReq) (*types.CheckChunkResp, error) {
	// todo: add your logic here and delete this line
	var resp = &types.CheckChunkResp{Status: 0}
	fileIdStr := strconv.FormatInt(req.FileId, 10)
	_, err := l.svcCtx.RDB.Exists(l.ctx, redis.UploadCheckBigFileKey+fileIdStr).Result()
	if err != nil {
		return nil, err
	}
	objectName := l.svcCtx.Minio.GenChunkObjectName(req.Hash, req.ChunkSeq)
	exist, err := l.svcCtx.Minio.NewService().IfExist(objectName)
	if err != nil {
		return resp, err
	} else if exist {
		resp.Status = 1
		return resp, nil
	}
	if err = l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf(redis.UploadCheckChunkKeyF, req.FileId, req.ChunkSeq), objectName, redis.UploadCheckChunkExpire).Err(); err != nil {
		return nil, err
	}
	return resp, nil
}
