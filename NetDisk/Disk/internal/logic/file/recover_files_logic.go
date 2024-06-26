package file

import (
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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

func (l *RecoverFilesLogic) RecoverFiles(req *types.RecoverFilesReq) error {
	// todo: add your logic here and delete this line
	var (
		ctx    = l.ctx
		userId = ctx.Value(define.UserIdKey).(int64)
		engine = l.svcCtx.Engine
		err    error
	)
	defer mqs.LogSend(l.ctx, err, "RecoverFiles", req.FileIds)

	fileIds, folderIds := req.FileIds, req.FolderIds
	_, err = engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		engine.ShowSQL(true)
		// 1.恢复文件delFlag字段
		fileBean := &models.File{
			DelFlag:  define.StatusFileUndeleted,
			DelTime:  0,
			SyncFlag: define.FlagSyncWrite,
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
				DelTime: 0,
			}
			if _, err := session.Cols("del_flag", "del_time").In("id", folderIds).
				And("del_flag = ?", define.StatusFolderDeleted).
				Update(folderBean); err != nil {
				logx.Errorf("恢复文件夹信息出错，err: %v", err)
				return nil, err
			}
			// 3.查询父文件夹
			var ids []int64
			subQuery := "select distinct(parent_id) from folder where id in ("
			for i, folderId := range folderIds {
				subQuery += strconv.FormatInt(folderId, 10)
				if i != len(folderIds)-1 {
					subQuery += ","
				}
			}
			if err := session.Select("id").
				Table(&models.Folder{}).
				Where("id in("+subQuery+") )").
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
