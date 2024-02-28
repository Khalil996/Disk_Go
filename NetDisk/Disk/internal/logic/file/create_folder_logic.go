package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"log"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFolderLogic {
	return &CreateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFolderLogic) CreateFolder(req *types.CreateFolderReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	log.Println("UserId from context:", userId)
	folder := &models.Folder{}
	if req.ParentFolderId != 0 {
		if exist, err := l.svcCtx.Engine.ID(req.ParentFolderId).
			And("user_id = ?", userId).
			Exist(folder); err != nil {
			logx.Errorf("创建文件夹查询父文件夹是否存在出错，err: %v", err)
			return errors.New("发生错误，" + err.Error())
		} else if !exist {
			return errors.New("发生错误！")
		}
	}

	if exist, err := l.svcCtx.Engine.Where("name = ?", req.Name).
		And("parent_id = ?", req.ParentFolderId).Exist(folder); err != nil {
		logx.Errorf("创建文件夹查询同名文件夹是否存在出错，err: %v", err)
		return errors.New("发生错误，" + err.Error())
	} else if exist {
		return errors.New("文件夹名称已存在！")
	}

	newFolder := &models.Folder{
		ParentId: req.ParentFolderId,
		Name:     req.Name,
		UserId:   userId,
		DelFlag:  define.StatusFolderUndeleted,
	}
	if _, err := l.svcCtx.Engine.Insert(newFolder); err != nil {
		return errors.New("创建失败了，" + err.Error())
	}

	return nil
}
