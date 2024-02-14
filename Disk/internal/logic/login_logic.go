package logic

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserBasic)
	sign, err := l.svcCtx.Engine.Where("username =? and password=?", req.Name, common.MD5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !sign {
		err = errors.New("账号密码错误")
		return nil, err
	}
	token, err := common.GenerateToken(user.Id, user.UserName, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	refreshToken, err := common.GenerateToken(user.Id, user.UserName, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginRes)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
