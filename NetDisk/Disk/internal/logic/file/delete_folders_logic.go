package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"time"
	"xorm.io/xorm"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoldersLogic {
	return &DeleteFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoldersLogic) DeleteFolders(req *types.IdsReq) error {
	// todo: add your logic here and delete this line
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
	)

	folderIds := req.Ids
	_, err := engine.Transaction(func(session *xorm.Session) (interface{}, error) {
		var (
			folders []*models.Folder
			err     error
			now     = time.Now().Local().Unix()
		)

		for len(folderIds) > 0 {
			// 1.删除当前文件夹下的文件
			fileBean := &models.File{
				DelFlag: define.StatusFileDeleted,
				DelTime: now,
			}
			if _, err = session.In("folder_id", folderIds).
				And("user_id = ?", userId).
				And("del_flag = ?", define.StatusFileUndeleted).
				Update(fileBean); err != nil {
				return nil, err
			}

			// 2.删除当前选中的文件夹
			folderBean := &models.Folder{
				DelFlag: define.StatusFolderDeleted,
				DelTime: now,
			}
			if _, err = session.In("id", folderIds).
				And("user_id = ?", userId).
				And("del_flag = ?", define.StatusFolderUndeleted).
				Update(folderBean); err != nil {
				return nil, err
			}

			// 3.搜索下一层文件夹
			if err = session.Select("id").In("parent_id", folderIds).
				And("user_id = ?", userId).
				And("del_flag = ?", define.StatusFolderUndeleted).
				Find(&folders); err != nil {
				return nil, err
			}

			folderIds = []int64{}
			for _, folder := range folders {
				folderIds = append(folderIds, folder.Id)
			}
			folders = []*models.Folder{}
		}
		return nil, nil
	})
	if err != nil {
		return errors.New("删除过程出错！" + err.Error())
	}

	return nil
}
