package user

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"cloud_go/Disk/models"
	"context"
	"errors"
	redis2 "github.com/redis/go-redis/v9"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDetailLogic {
	return &UpdateDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDetailLogic) UpdateDetail(req *types.UpdateUserDetailReq) error {
	// todo: add your logic here and delete this line
	var (
		loginUserId = l.ctx.Value(define.UserIdKey).(int64)
		userIdStr   = strconv.FormatInt(loginUserId, 10)
		engine      = l.svcCtx.Engine
		rdb         = l.svcCtx.RDB
		user        models.UserBasic
	)

	key := redis.UserInfoKey + userIdStr
	m, err := rdb.HGetAll(l.ctx, key).Result()
	if err != nil && err != redis2.Nil {
		logx.Errorf("更新用户info，redis获取失败，ERR: [%v]", err)
	} else if id, ok := m["id"]; err == redis2.Nil || !ok || id == "" {
		cols := "name, username, avatar, email, signature, status, used, capacity"
		if has, err := engine.Cols(cols).ID(loginUserId).Get(&user); err != nil {
			logx.Errorf("更新用户info，数据库info获取失败，ERR: [%v]", err)
			return err
		} else if !has {
			return errors.New("用户信息有误")
		}
	} else {
		status, _ := strconv.Atoi(m["status"])
		used, _ := strconv.ParseInt(m["used"], 10, 64)
		capacity, _ := strconv.ParseInt(m["capacity"], 10, 64)
		user.Id = loginUserId
		user.Name = req.Name
		user.UserName = m["username"]
		user.Avatar = m["avatar"]
		user.Email = req.Email
		user.Signature = req.Signature
		user.Status = int8(status)
		user.Used = used
		user.Capacity = capacity
	}

	if req.Email != user.Email {
		code, err := rdb.Get(l.ctx, redis.EmailValidCode+req.Email).Result()
		if err != nil && err != redis2.Nil {
			logx.Errorf("更新用户info，redis获取邮箱验证码失败，ERR: [%v]", err)
			return errors.New("出错啦")
		}
		if code != req.Code {
			return errors.New("验证码错误！请重新获取😭")
		}
	}

	user.Email = req.Email
	if affected, err := engine.Cols("name", "email", "signature").
		ID(loginUserId).Update(user); err != nil {
		logx.Errorf("更新用户info，更新数据库失败，ERR: [%v]", err)
		return errors.New("出错了，请稍后，" + err.Error())
	} else if affected != 1 {
		return errors.New("出错了，请稍后")
	}

	m2 := map[string]interface{}{
		"id":        user.Id,
		"name":      user.Name,
		"username":  user.UserName,
		"avatar":    user.Avatar,
		"email":     user.Email,
		"signature": user.Signature,
		"capacity":  user.Capacity,
		"status":    user.Status,
		"used":      user.Used,
	}
	if err = rdb.HSet(l.ctx, key, m2).Err(); err != nil {
		logx.Errorf("更新用户info，更新redis失败，ERR: [%v]", err)
	}

	return nil
}
