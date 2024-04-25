package logic

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)

	defer mqs.LogSend(l.ctx, err, "Login", username)
	// todo: add your logic here and delete this line

	user := new(models.UserBasic)
	sign, err := l.svcCtx.Engine.Where("username =? and password=?", username, common.MD5(password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !sign {
		err = errors.New("账号密码错误")
		return nil, err
	}
	if user.Status != define.StatusUserOk {
		reason := fmt.Sprintf(define.BanStr, define.BanM[user.Status])
		return nil, errors.New(reason)
	}
	token, err := common.GenerateToken(user.Id, user.UserName)
	if err != nil {
		return nil, err
	}
	key := redis.UserLogin + strconv.FormatInt(user.Id, 10)
	if err = l.svcCtx.RDB.Set(l.ctx, key, token, 7*24*time.Hour).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] Login 保存用户token失败，userid：%v %v\n", user.Id, err)
		l.svcCtx.RDB.Set(l.ctx, key, token, 7*24*time.Hour) // 重试
	}
	var userInfo types.UserInfo
	userInfo.UserId = user.Id
	userInfo.Name = user.Name
	userInfo.Avatar = user.Avatar
	userInfo.Email = user.Email
	userInfo.Signature = user.Signature
	userInfo.Status = user.Status

	resp = new(types.LoginResp)
	resp.Token = token
	resp.UserInfo = userInfo
	return
}
