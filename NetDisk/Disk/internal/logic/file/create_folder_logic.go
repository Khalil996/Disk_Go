package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"fmt"
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
	var err error

	defer mqs.LogSend(l.ctx, err, "CreateFolder", req.Name)

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

	baseName := req.Name
	newName := baseName
	suffix := 0

	for {
		// 检查当前生成的名称是否存在
		exist, err := l.svcCtx.Engine.Where("name = ?", newName).
			And("parent_id = ?", req.ParentFolderId).Exist(new(models.Folder))
		if err != nil {
			logx.Errorf("检查文件夹名称是否存在时出错，err: %v", err)
			return errors.New("发生错误，" + err.Error())
		}

		// 如果存在，生成新的名称并重新检查
		if exist {
			suffix++                                        // 增加数字后缀
			newName = fmt.Sprintf("%s%d", baseName, suffix) // 生成新的名称
		} else {
			break // 找到了一个不存在的名称，跳出循环
		}
	}

	newFolder := &models.Folder{
		ParentId: req.ParentFolderId,
		Name:     newName,
		UserId:   userId,
		DelFlag:  define.StatusFolderUndeleted,
	}
	if _, err := l.svcCtx.Engine.Insert(newFolder); err != nil {
		return errors.New("创建失败了，" + err.Error())
	}

	return nil
}
