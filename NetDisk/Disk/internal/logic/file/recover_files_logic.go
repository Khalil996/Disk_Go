package file

import (
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"time"

	"cloud_go/Disk/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverFilesLogic {
	return &RecoverFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverFilesLogic) RecoverFiles(fileIds, folderIds []int64) error {
	// todo: add your logic here and delete this line
	var (
		ctx    = l.ctx
		userId = ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
	)

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		// 1.恢复文件delFlag字段
		fileBean := &models.File{
			DelFlag: define.StatusFileUndeleted,
			Model:   models.Model{DeleteAt: time.Now().Local().UTC()},
		}
		if affected, err := session.In("id", fileIds).
			And("user_id = ?", userId).
			Update(fileBean); err != nil {
			logx.Errorf("恢复文件信息出错，err: %v", err)
			return nil, err
		} else if affected != int64(len(fileIds)) {
			return nil, errors.New("恢复文件信息出错！")
		}

		for len(folderIds) > 0 {
			// 2.恢复文件夹delFlag字段
			folderBean := &models.Folder{
				DelFlag: define.StatusFolderUndeleted,
				Model:   models.Model{DeleteAt: time.Now().Local().UTC()},
			}
			if affected, err := session.In("id", folderIds).
				And("del_flag = ?", define.StatusFolderDeleted).
				Update(folderBean); err != nil {
				logx.Errorf("恢复文件夹信息出错，err: %v", err)
				return nil, err
			} else if affected != int64(len(folderIds)) {
				return nil, errors.New("恢复文件夹信息出错！")
			}

			// 3.查询父文件夹
			var ids []int64
			if err := session.Select("id").
				Where("id in ( select distinct(parent_id) "+
					"where id in (?) )", folderIds).
				And("del_flag = ?", define.StatusFolderDeleted).
				Find(&ids); err != nil {
				logx.Errorf("查询父文件夹出错，err: %v", err)
				return nil, err
			}
			folderIds = ids
		}
		return nil, nil
	})

	return err
}
