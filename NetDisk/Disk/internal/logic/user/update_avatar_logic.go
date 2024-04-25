package user

import (
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"cloud_go/Disk/models"
	"context"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(fileParam *types.FileParam) (interface{}, error) {
	var (
		loginUserId = l.ctx.Value(define.UserIdKey).(int64)
		userIdStr   = strconv.FormatInt(loginUserId, 10)
		engine      = l.svcCtx.Engine
		rdb         = l.svcCtx.RDB
		minioSvc    = l.svcCtx.Minio.NewService()
	)

	index := strings.LastIndex(fileParam.FileHeader.Filename, ",")
	ext := fileParam.FileHeader.Filename[index+1:]
	objectName := "/avatar/" + userIdStr + ext

	var (
		urlErr error
		url    string
	)
	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		user := &models.UserBasic{Avatar: objectName}
		if _, err := session.ID(loginUserId).
			Update(user); err != nil {
			logx.Errorf("更新头像，更新数据库失败，ERR: [%v]", err)
			return nil, err
		}

		if err := minioSvc.Upload(l.ctx, objectName, fileParam.File); err != nil {
			return nil, err
		}

		url, urlErr = minioSvc.GenUrl(objectName, "", false)

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	if urlErr != nil {
		logx.Errorf("更新头像，生成url失败，ERR: [%v]", urlErr)
		return nil, nil
	}

	userInfoKey := redis.UserInfoKey + userIdStr
	m := make(map[string]interface{})
	if _, err = engine.Select("id, name, username, email, signature, status, used, capacity").
		ID(loginUserId).Table(&models.UserBasic{}).Get(&m); err != nil {
		logx.Errorf("更新头像，获取用户信息失败，ERR: [%v]", err)
		return nil, nil
	}

	m["avatar"] = url
	pipeline := rdb.Pipeline()
	pipeline.HSet(l.ctx, userInfoKey, m)
	pipeline.Expire(l.ctx, userInfoKey, redis.UserInfoExpire)
	if _, err = pipeline.Exec(l.ctx); err != nil {
		logx.Errorf("更新头像，用户信息->redis失败，ERR: [%v]", err)
	}

	return nil, nil
}
