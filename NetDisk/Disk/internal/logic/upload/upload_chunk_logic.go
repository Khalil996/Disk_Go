package upload

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"mime/multipart"
	"strconv"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadChunkLogic {
	return &UploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadChunkLogic) UploadChunk(req *types.UploadChunkReq, fileParam *types.FileParam) error {
	// todo: add your logic here and delete this line
	checkKey := fmt.Sprintf(redis.UploadCheckChunkKeyF, req.FileId, req.ChunkSeq)
	// 判断是否已经上传过
	objectName, err := l.svcCtx.RDB.Get(l.ctx, checkKey).Result()
	if err != nil {
		return err
	}
	// 判断是否是第一次上传
	bigFileKey := redis.UploadCheckBigFileKey + strconv.FormatInt(req.FileId, 10)
	fileInfo, err := l.svcCtx.RDB.HGetAll(l.ctx, bigFileKey).Result()

	if err != nil {
		return err
	}
	// 判断是否是最后一个分片
	chunkSum, _ := strconv.ParseInt(fileInfo["chunkSum"], 10, 64)
	chunkNum, _ := strconv.ParseInt(fileInfo["chunkNum"], 10, 64)

	if chunkSum+1 == chunkNum {
		_, err = l.svcCtx.Engine.DoTransaction(l.createSchedule(req, fileParam.File, objectName, chunkSum, fileInfo))
	} else {
		if err = l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileParam.File); err != nil {
			return err
		}
		_, err = l.incr(bigFileKey, 1)
	}
	_, _ = l.svcCtx.RDB.Del(l.ctx, checkKey).Result()

	return nil
}

// 创建事务
func (l *UploadChunkLogic) createSchedule(req *types.UploadChunkReq, fileData multipart.File,
	objectName string, chunkNum int64, fileInfo map[string]string) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {

		size, _ := strconv.ParseInt(fileInfo["size"], 10, 64)
		fsId := idgen.NextId()
		fileFs := &models.FileFs{}
		fileFs.Id = fsId
		fileFs.Bucket = l.svcCtx.Minio.BucketName
		fileFs.Ext = fileInfo["ext"]
		fileFs.Name = fileInfo["name"]
		fileFs.Hash = fileInfo["hash"]
		fileFs.Size = size
		fileFs.ObjectName = objectName
		fileFs.Url = ""
		fileFs.ChunkNum = chunkNum
		fileFs.Status = define.StatusFsBigFileNeedMerge
		if _, err := session.Insert(fileFs); err != nil {
			return nil, err
		}

		userId, _ := strconv.ParseInt(fileInfo["userId"], 10, 64)
		folderId, _ := strconv.ParseInt(fileInfo["folderId"], 10, 64)
		file := models.File{}
		file.Name = fileInfo["name"]
		file.UserId = userId
		file.FsId = fsId
		file.FolderId = folderId
		file.Url = ""
		file.Ext = fileInfo["ext"]
		file.ObjectName = objectName
		file.Size = size
		file.Type = define.GetTypeByBruteForce(fileInfo["ext"])
		file.Status = define.StatusFileUploaded
		file.IsBig = define.BigFileFlag
		file.DoneAt = time.Now().Local()
		if _, err := session.Insert(file); err != nil {
			return nil, err
		}

		fileSchedule := &models.FileSchedule{}
		fileSchedule.FileId = req.FileId
		fileSchedule.FsId = fsId
		fileSchedule.ChunkNum = chunkNum
		if _, err := session.Insert(fileSchedule); err != nil {
			return nil, err
		}

		key := redis.UploadCheckBigFileKey + strconv.FormatInt(req.FileId, 10)
		if _, err := l.incr(key, 1); err != nil {
			return nil, err
		}

		if err := l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileData); err != nil {
			l.incr(key, -1)
			return nil, err
		}

		return nil, nil
	}
}

// redis指定key+1
func (l *UploadChunkLogic) incr(key string, value int64) (int64, error) {
	return l.svcCtx.RDB.HIncrBy(l.ctx, key, "chunkSum", value).Result()
}
