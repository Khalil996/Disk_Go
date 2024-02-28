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
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)

	fileIdStr := strconv.FormatInt(req.FileId, 10)
	key := redis.UploadCheckKey + fileIdStr
	fileInfo, err := l.svcCtx.RDB.HGetAll(l.ctx, key).Result()
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
		if has, err := l.svcCtx.Engine.Where("id =?", folderId).Exist(&models.FileFs{}); err != nil {
			return nil, err
		} else if !has {
			return nil, errors.New("文件夹不存在")
		}
	}
	_, err = l.svcCtx.Engine.DoTransaction(l.saveAndUpload(fileInfo, fileParam.File))
	if err != nil {
		return nil, err
	}
	go l.svcCtx.RDB.Del(l.ctx, key)
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
		file.Size = size
		file.FsId = fsId
		file.FolderId = folderId
		file.UserId = userId
		file.IsBig = define.SmallFileFlag
		file.DoneAt = time.Now().Local()
		file.Status = define.StatusFileUploaded
		if _, err := session.Insert(file); err != nil {
			return nil, err
		}

		if err := l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileData); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
