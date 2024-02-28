package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"fmt"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFilesLogic {
	return &MoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFilesLogic) MoveFiles(req *types.MoveFilesReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	folderId := req.FolderId

	if folderId != 0 {
		has, err := l.svcCtx.Engine.ID(folderId).And("user_id = ?", userId).Get(&models.Folder{})
		if err != nil {
			return errors.New("发生错误！" + err.Error())
		} else if !has {
			return errors.New("该目录不存在")
		}
	}

	file := &models.File{FolderId: folderId}
	affected, err := l.svcCtx.Engine.In("id", req.FileIds).Update(file)
	fmt.Println(affected, err)
	if err != nil {
		return errors.New("移动出错！" + err.Error())
	} else if affected != int64(len(req.FileIds)) {
		return errors.New("移动出错！")
	}

	return nil
}
