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

type GetUserInfoByShareIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByShareIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByShareIdLogic {
	return &GetUserInfoByShareIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByShareIdLogic) GetUserInfoByShareId(req *types.IdPathReq) (map[string]interface{}, error) {
	// todo: add your logic here and delete this line
	var user models.UserBasic
	var share models.Share
	//查找文件
	if has, err := l.svcCtx.Engine.ID(req.Id).Get(&share); err != nil {
		return nil, err
	} else if !has {
		return map[string]interface{}{
			"shareStatus": define.StatusShareNotExistOrDeleted,
		}, nil
	}
	//查找用户
	if has, err := l.svcCtx.Engine.ID(share.UserId).Get(&user); err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("用户不存在")
	}
	url, err := l.svcCtx.Minio.NewService().GenUrl(user.Avatar, "", false)
	if err != nil {
		return nil, errors.New("头像获取失败")
	}
	resp := map[string]interface{}{
		"userId":      user.Id,
		"name":        user.Name,
		"avatar":      url,
		"signature":   user.Signature,
		"userStatus":  user.Status,
		"shareStatus": share.Status,
	}

	return resp, nil
}
