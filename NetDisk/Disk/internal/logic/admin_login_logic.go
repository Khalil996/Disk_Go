package logic

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	//admin := &models.Admin{
	//	Username: username,
	//}
	var err error

	defer mqs.LogSend(l.ctx, err, "AdminLogin", username)
	admin := new(models.Admin)
	sign, err := l.svcCtx.Engine.Where("username =? and password=?", username, common.MD5(password)).Get(admin)
	if err != nil {
		return nil, err
	}
	if !sign {
		err = errors.New("账号密码错误")
		return nil, err
	}
	if admin.Status == define.StatusAdminBanned {
		err = errors.New("该管理员账号不可用")
		return nil, err
	}
	//if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
	//	return nil, errors.New("用户名或密码错误2")
	//}
	token, err := common.GenerateToken(admin.Id, admin.Username)
	if err != nil {
		return nil, errors.New("系统错误")
	}
	key := redis.UserLogin + strconv.FormatInt(admin.Id, 10)
	if err := l.svcCtx.RDB.Set(l.ctx, key, token, 7*24*time.Hour); err != nil {
		logx.Errorf("设置用户登录信息失败,err:%v", err)
		l.svcCtx.RDB.Set(l.ctx, key, token, 7*24*time.Hour) //重试
	}
	var userInfo types.UserInfo
	userInfo.UserId = admin.Id
	userInfo.Name = admin.Name
	userInfo.Status = admin.Status

	resp := &types.LoginResp{}
	resp.Token = token
	resp.UserInfo = userInfo
	return resp, nil
}
