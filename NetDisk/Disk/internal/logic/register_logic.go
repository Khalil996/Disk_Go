package logic

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	// todo: add your logic here and delete this line
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if req.Name == "" || req.Email == "" || req.Password == "" || req.Code == "" {
		err = errors.New("参数不能为空")
		return nil, err
	}
	Name := "user_" + strconv.FormatInt(int64(rand.Int31()), 10)
	if err != nil {
		return nil, errors.New("验证码已过期")
	}
	if code != req.Code {
		err = errors.New("验证码不一致")
		return nil, err
	}
	cnt, err := l.svcCtx.Engine.Where("username = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("账号已存在")
		return nil, err
	}
	user := &models.UserBasic{
		UserName: req.Name,
		Password: common.MD5(req.Password),
		Name:     Name,
		Email:    req.Email,
		Status:   0,
		IsAdmin:  0,
		Used:     0,
		Capacity: define.DefaultCapacity,
	}
	n, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println(n)
	return resp, nil
}
