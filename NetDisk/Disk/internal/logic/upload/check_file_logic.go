package upload

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"math"
	"strconv"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFileLogic {
	return &CheckFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFileLogic) CheckFile(req *types.CheckFileReq) (resp *types.CheckFileResp, err error) {
	// todo: add your logic here and delete this line
	var fileFs models.FileFs
	has, err := l.svcCtx.Engine.Where("hash=?", req.Hash).Get(&fileFs)
	if err != nil {
		return nil, err
	} else if !has {
		return l.doWhenNotExist(req)
	}

	return l.doWhenExist(req, &fileFs)
}

// 检查文件是否存在
func (l *CheckFileLogic) doWhenNotExist(req *types.CheckFileReq) (*types.CheckFileResp, error) {
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		resp   types.CheckFileResp
	)

	fileId := idgen.NextId()
	fileIdStr := strconv.FormatInt(fileId, 10)
	key := redis.UploadCheckKey + fileIdStr
	resp.Status = define.StatusFileUnuploaded
	resp.FileId = fileId
	fileInfo := map[string]interface{}{
		"fileId":   fileId,
		"folderId": req.FolderId,
		"hash":     req.Hash,
		"ext":      req.Ext,
		"name":     req.Name,
		"size":     req.Size,
		"userId":   userId,
	}

	// 大文件情况
	if req.Size > int64(define.ShardingFloor) {
		resp.ConfirmShard = define.ConfirmShard
		key = redis.UploadCheckBigFileKey + fileIdStr
		fileInfo["chunkNum"] = math.Ceil(float64(req.Size) / define.ShardingSize)
		fileInfo["chunkSum"] = 0
		if _, err := l.svcCtx.RDB.HSet(l.ctx, key, fileInfo).Result(); err != nil {
			return nil, err
		}
	}
	if _, err := l.svcCtx.RDB.HSet(l.ctx, key, fileInfo).Result(); err != nil {
		return nil, err
	}
	go l.svcCtx.RDB.Expire(l.ctx, key, redis.UploadCheckExpire)
	return &resp, nil
}

func (l *CheckFileLogic) doWhenExist(req *types.CheckFileReq, fileFs *models.FileFs) (*types.CheckFileResp, error) {
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
		file   models.File
		resp   types.CheckFileResp
	)

	// 先判断该用户在该目录下有无该文件
	has, err := engine.Where("fs_id = ?", fileFs.Id).
		And("folder_id = ?", req.FolderId).
		And("user_id = ?", userId).Get(&file)
	if err != nil {
		return nil, err
	} else if has {
		return nil, errors.New("当前文件夹已存在该文件")
	}

	// 该文件夹无该文件，信息落库
	isBigFlag := define.SmallFileFlag
	if fileFs.Size > int64(define.ShardingFloor) {
		isBigFlag = define.BigFileFlag
	}
	file.UserId = userId
	file.FsId = fileFs.Id
	file.Name = req.Name
	file.FolderId = req.FolderId
	file.Status = define.StatusFileUploaded
	file.Url = fileFs.Url
	file.IsBig = isBigFlag
	file.DoneAt = time.Now().Local()
	file.DelFlag = define.StatusFileUndeleted
	if fileFs.Size > int64(define.ShardingFloor) {
		file.IsBig = define.BigFileFlag
	}
	if _, err = engine.Insert(&file); err != nil {
		return nil, err
	}

	resp.Status = define.StatusFileUploaded
	return &resp, nil
}
