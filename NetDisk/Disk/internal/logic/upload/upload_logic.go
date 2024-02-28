package upload

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadReq, fileParam *types.FileParam) (interface{}, error) {
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
		rdb    = l.svcCtx.RDB
	)

	fileIdStr := strconv.FormatInt(req.FileId, 10)
	key := redis.UploadCheckKey + fileIdStr
	fileInfo, err := rdb.HGetAll(l.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if fileInfo["userId"] != strconv.FormatInt(userId, 10) {
		return nil, errors.New("信息有误1")
	}
	if fileInfo["name"] != fileParam.FileHeader.Filename {
		return nil, errors.New("信息有误2")
	}

	folderId := fileInfo["folderId"]
	if folderId != "0" {
		if has, err := engine.Where("id = ?", folderId).
			Exist(&models.FileFs{}); err != nil {
			return nil, err
		} else if !has {
			return nil, errors.New("信息有误3")
		}
	}

	_, err = engine.DoTransaction(l.saveAndUpload(fileInfo, fileParam.File))
	if err != nil {
		return nil, err
	}

	go rdb.Del(l.ctx, key)
	return nil, nil
}

func (l *UploadLogic) saveAndUpload(fileInfo map[string]string, fileData multipart.File) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {

		size, _ := strconv.ParseInt(fileInfo["size"], 10, 64)
		filename, objectName := l.svcCtx.Minio.GenObjectName(fileInfo["hash"], fileInfo["ext"])
		fsId := idgen.NextId()
		fileFs := &models.FileFs{}
		fileFs.Id = fsId
		fileFs.Bucket = l.svcCtx.Minio.BucketName
		fileFs.Name = filename
		fileFs.ObjectName = objectName
		fileFs.Ext = fileInfo["ext"]
		fileFs.Hash = fileInfo["hash"]
		fileFs.Size = size
		fileFs.Url = ""
		fileFs.Status = define.StatusFsUploaded
		if _, err := session.Insert(fileFs); err != nil {
			return nil, err
		}

		userId, _ := strconv.ParseInt(fileInfo["userId"], 10, 64)
		folderId, _ := strconv.ParseInt(fileInfo["folderId"], 10, 64)
		file := &models.File{}
		file.Name = fileInfo["name"]
		file.Url = ""
		file.ObjectName = objectName
		file.Size = size
		file.Ext = fileInfo["ext"]
		file.FsId = fsId
		file.FolderId = folderId
		file.UserId = userId
		file.Type = define.GetTypeByBruteForce(fileInfo["ext"])
		file.IsBig = define.SmallFileFlag
		file.DoneAt = time.Now().Local()
		file.Status = define.StatusFileUploaded
		file.DelFlag = define.StatusFileUndeleted
		if _, err := session.Insert(file); err != nil {
			return nil, err
		}

		if err := l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileData); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
