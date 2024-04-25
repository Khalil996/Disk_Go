package admin

import (
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlLogic {
	return &GetUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlLogic) GetUrl(req *types.GetUrlReq) (url string, err error) {
	// todo: add your logic here and delete this line
	var (
		file       models.File
		shareFile  models.ShareFile
		has        bool
		objectName string
	)
	if req.Type == 0 {
		if has, err = l.svcCtx.Engine.Table(&models.Share{}).Alias("a").Select("c.object_name").
			Join("LEFT", []string{shareFile.ShareFile(), "b"}, "a.id=b.share_id").
			Join("LEFT", []string{file.TableName(), "c"}, "c.id=b.file_id").Where("a.id=?", req.Id).Get(&objectName); err != nil {
			logx.Errorf("GetUrlLogic GetUrl err:%v", err)
			return "", err
		}
	} else {
		if has, err = l.svcCtx.Engine.Select("object_name").
			Table(&models.File{}).ID(req.Id).
			Get(&objectName); err != nil {
			logx.Errorf("GetUrlLogic GetUrl err:%v", err)
			return "", err
		}
	}
	if !has {
		return "", errors.New("没有找到文件链接")
	}
	url, err = l.svcCtx.Minio.NewService().GenUrl(objectName, file.Name, true)

	return url, nil
}
