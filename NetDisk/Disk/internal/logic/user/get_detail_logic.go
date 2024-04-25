package user

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/models"
	"context"
	"errors"
	redis2 "github.com/redis/go-redis/v9"
	"strconv"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailLogic) GetDetail(req *types.IdPathReq, loginUserId int64) (interface{}, error) {
	// todo: add your logic here and delete this line
	var (
		userIdStr = strconv.FormatInt(loginUserId, 10)
		engine    = l.svcCtx.Engine
		rdb       = l.svcCtx.RDB
		minioSvc  = l.svcCtx.Minio.NewService()
		user      models.UserBasic
	)

	targetUserId := req.Id
	if req.Id == 0 {
		targetUserId = loginUserId
	}

	key := redis.UserInfoKey + userIdStr
	m, err := rdb.HGetAll(l.ctx, key).Result()
	if err != nil && err != redis2.Nil {
		logx.Errorf("获取用户info，redis获取失败，ERR: [%v]", err)
	} else if id, ok := m["id"]; err == redis2.Nil || !ok || id == "" {
		cols := "id, name, username, avatar, email, signature, status, used, capacity"
		if has, err := engine.Select(cols).ID(targetUserId).Get(&user); err != nil {
			logx.Errorf("更新用户info，数据库info获取失败，ERR: [%v]", err)
			return nil, err
		} else if !has {
			return nil, errors.New("用户信息有误")
		}
		url, _ := minioSvc.GenUrl(user.Avatar, "", false)
		m2 := map[string]interface{}{
			"id":        user.Id,
			"name":      user.Name,
			"username":  user.UserName,
			"avatar":    url,
			"email":     user.Email,
			"signature": user.Signature,
			"capacity":  user.Capacity,
			"status":    user.Status,
			"used":      user.Used,
		}
		if err = rdb.HSet(l.ctx, key, m2).Err(); err != nil {
			logx.Errorf("更新用户info，info->redis失败，ERR: [%v]", err)
		}
		go func() {
			if err2 := rdb.Expire(l.ctx, key, redis.UserInfoExpire).Err(); err2 != nil {
				logx.Errorf("更新用户info，info->redis设置expire失败，ERR: [%v]", err2)
				return
			}
		}()
		return m2, nil
	}
	return m, nil
}
