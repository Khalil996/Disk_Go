package logic

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailSendLogic {
	return &EmailSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailSendLogic) EmailSend(req *types.EmailSendReq) (resp *types.EmailSendRes, err error) {
	// todo: add your logic here and delete this line
	if req.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	cnt, err := l.svcCtx.Engine.Where("email=?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("邮箱已存在")
		return nil, err
	}

	code := common.RandCode()
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	err = common.SendCode(req.Email, code)
	log.Println(err)
	return resp, nil
}
