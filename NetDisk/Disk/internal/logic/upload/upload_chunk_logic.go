package upload

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"mime/multipart"
	"strconv"
	"strings"
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

	chunkSum, _ := strconv.ParseInt(fileInfo["chunkSum"], 10, 64)
	chunkNum, _ := strconv.ParseInt(fileInfo["chunkNum"], 10, 64)
	//判断是否是最后一个分片
	if chunkSum+1 == chunkNum {
		_, err = l.svcCtx.Engine.DoTransaction(l.createSchedule(req, fileParam.File, objectName, chunkSum, fileInfo))
		l.svcCtx.RDB.Del(l.ctx, bigFileKey)
	} else {
		//分块上传，然后对分块进行+1
		if err = l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileParam.File); err != nil {
			return err
		}
		_, err = l.incr(bigFileKey, 1)
	}
	l.svcCtx.RDB.Del(l.ctx, checkKey)

	return nil
}

// 创建事务
func (l *UploadChunkLogic) createSchedule(req *types.UploadChunkReq, fileData multipart.File,
	objectName string, chunkNum int64, fileInfo map[string]string) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {

		objectName2 := objectName[:strings.LastIndex(objectName, "@")]
		size, _ := strconv.ParseInt(fileInfo["size"], 10, 64)
		fsId := idgen.NextId()
		fileFs := &models.FileFs{}
		fileFs.Id = fsId
		fileFs.Bucket = l.svcCtx.Minio.BucketName
		fileFs.Ext = fileInfo["ext"]
		fileFs.Name = fileInfo["name"]
		fileFs.Hash = fileInfo["hash"]
		fileFs.Size = size
		fileFs.ObjectName = objectName2
		fileFs.ChunkNum = chunkNum
		fileFs.Status = define.StatusFsFileNeedMerge
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
		file.Ext = fileInfo["ext"]
		file.ObjectName = objectName2
		file.Size = size
		file.Type = define.GetTypeByBruteForce(fileInfo["ext"])
		file.Status = define.StatusFileUploaded
		file.IsBig = define.BigFileFlag
		file.DoneAt = time.Now().Local()
		file.DelFlag = define.StatusFileUndeleted
		file.SyncFlag = define.FlagSyncWrite
		if _, err := session.Insert(file); err != nil {
			return nil, err
		}

		fileScheduleId := idgen.NextId()
		fileSchedule := &models.FileSchedule{}
		fileSchedule.Id = fileScheduleId
		fileSchedule.FileId = req.FileId
		fileSchedule.FsId = fsId
		fileSchedule.ChunkNum = chunkNum
		fileSchedule.Stage = define.StageMerging
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

		ms := &common.MergeStruct{}
		ms.ObjectName = objectName
		ms.Hash = fileInfo["hash"]
		ms.FsId = fsId
		ms.SId = fileScheduleId
		ms.ChunkNum = chunkNum
		ms.Size = size
		go common.Merge(ms, l.errCallBack)

		return nil, nil
	}
}

// redis指定key+1
func (l *UploadChunkLogic) incr(key string, value int64) (int64, error) {
	return l.svcCtx.RDB.HIncrBy(l.ctx, key, "chunkSum", value).Result()
}

func (l *UploadChunkLogic) errCallBack(sId int64) {
	fs := &models.FileSchedule{Stage: define.StageNeedMerge}
	if _, err := l.svcCtx.Engine.ID(sId).Update(fs); err != nil {
		logx.Errorf("createSchedule->errCallBack，退回schedule:%v 状态出错，ERR：[%v]", sId, err)
		return
	}
}
