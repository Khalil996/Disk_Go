package logic

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareInfoLogic {
	return &GetShareInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareInfoLogic) GetShareInfo(req *types.GetShareInfoReq) (resp *types.GetShareInfoResp, err error) {
	// todo: add your logic here and delete this line
	var files []*models.File
	var share models.Share
	if has, err := l.svcCtx.Engine.Where("id = ?", req.Id).Get(&share); err != nil {
		logx.Errorf("查询分享信息失败:%v", err)
		return nil, err
	} else if !has {
		return nil, errors.New("分享信息不存在")
	}
	if share.Pwd != req.Pwd {
		return nil, errors.New("密码错误")
	}
	//获取响应内容
	resp = &types.GetShareInfoResp{}

	resp.Name = share.Name
	resp.Created = share.Created.Format(define.TimeFormat1)
	resp.Expired = share.Expired
	resp.Owner = share.UserId

	//查询分享文件
	if err := l.svcCtx.Engine.Where("id in (select file_id from share_file where share_id = ?)", req.Id).Find(&files); err != nil {
		logx.Errorf("查询分享文件失败:%v", err)
		return nil, err
	}

	//
	for _, file := range files {
		url, err := l.svcCtx.Minio.NewService().GenUrl(file.ObjectName, file.Name, true)
		if err != nil {
			logx.Errorf("生成分享文件下载链接失败:%v", err)
			url = ""
		}
		resp.Items = append(resp.Items, &types.ListShareItemStruct{
			Id:      file.Id,
			Name:    file.Name,
			Updated: file.Updated.Format(define.TimeFormat1),
			Size:    file.Size,
			Url:     url,
			Status:  file.Status,
		})
	}

	if _, err := l.svcCtx.Engine.Where("id = ?", req.Id).
		SetExpr("click_num", "click_num + 1").
		Update(&models.Share{}); err != nil {
		logx.Errorf("获取分享列表，更新点击次数失败，ERR: [%v]", err)
	}

	return resp, nil
}
