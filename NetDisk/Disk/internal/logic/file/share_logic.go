package file

import (
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareLogic {
	return &ShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareLogic) Share(req *types.ShareReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var (
		shareType int8 = define.TypeShareMulti
		err       error
	)
	defer mqs.LogSend(l.ctx, err, "Share", req.FileIds)
	if req.Pwd == "" {
		return errors.New("密码不能为空")
	}

	var file models.File
	if has, err := l.svcCtx.Engine.Select("name,ext").
		ID(req.FileIds[0]).Get(&file); err != nil {
		logx.Errorf("查询文件失败:%v", err)
		return errors.New("出错了")
	} else if !has {
		return errors.New("文件不存在")
	}

	shareName := file.Name
	if len(req.FileIds) == 1 {
		shareType = define.GetTypeByBruteForce(file.Ext)
	} else {
		shareName += "等..."
	}

	id := strconv.FormatInt(idgen.NextId(), 10)
	url := req.Prefix + id
	if req.Auto == 1 {
		url += "?pwd=" + req.Pwd
	}
	_, err = l.svcCtx.Engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		var shareFile []*models.ShareFile
		for _, fileId := range req.FileIds {
			shareFile = append(shareFile, &models.ShareFile{
				ShareId: id,
				FileId:  fileId,
			})
		}
		if _, err := session.Insert(shareFile); err != nil {
			logx.Errorf("插入分享文件失败:%v", err)
			return nil, err
		}

		created := time.Now().Local()
		expired := created.Unix() + define.ShareExpireType[req.ExpireType]
		share := &models.Share{}
		share.Id = id
		share.UserId = userId
		share.Name = shareName
		share.Pwd = req.Pwd
		share.Created = created
		share.Expired = expired
		share.Url = url
		share.Type = shareType

		if req.ExpireType == define.ShareExpireForever {
			share.Status = define.StatusShareForever
		}
		if _, err := session.Insert(share); err != nil {
			logx.Errorf("插入分享失败:%v", err)
			return nil, err
		}
		return nil, nil
	})

	return err
}
