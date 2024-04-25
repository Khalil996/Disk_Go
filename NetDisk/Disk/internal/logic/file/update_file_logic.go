package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"strings"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFileLogic) UpdateFile(req *types.UpdateNameReq) error {
	// todo: add your logic here and delete this line

	userId := l.ctx.Value(define.UserIdKey).(int64)
	var err error
	defer mqs.LogSend(l.ctx, err, "UpdateFile", req.Id, req.Name)

	ext := req.Name[strings.LastIndex(req.Name, "."):]
	fType := define.GetTypeByBruteForce(ext)
	file := &models.File{Name: req.Name, Type: fType}
	if _, err = l.svcCtx.Engine.ID(req.Id).
		And("user_id = ?", userId).
		Update(file); err != nil {
		return err
	}
	return nil
}
