package file

import (
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
	"time"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFolderLogic {
	return &ShareFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFolderLogic) ShareFolder(req *types.ShareFolderReq) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var (
		folders    []*models.Folder
		shareFiles []*models.ShareFile
		err        error
	)
	defer mqs.LogSend(l.ctx, err, "ShareFolder", req.FolderIds)

	if req.Pwd == "" {
		return errors.New("密码不能为空")
	}
	// 生成分享链接
	id := strconv.FormatInt(idgen.NextId(), 10)
	url := req.Prefix + id + "?pwd=" + req.Pwd
	// 查询文件夹
	if has, err := l.svcCtx.Engine.In("id", req.FolderIds).Get(&folders); err != nil {
		logx.Errorf("查询文件夹失败:%v", err)
		return errors.New("查询文件夹失败")
	} else if !has {
		return errors.New("文件夹不存在")
	}
	folderName := folders[0].Name
	for len(req.FolderIds) > 0 {
		// 查询文件夹下的文件
		var fileIds []int64
		if err := l.svcCtx.Engine.In("folder_id", req.FolderIds).
			And("user_id=?", userId).
			And("del_flag=?", define.StatusFileUndeleted).
			Find(&fileIds); err != nil {
			logx.Errorf("查询文件夹下的文件失败:%v", err)
			return errors.New("查询文件夹下的文件失败")
		}

		// 查询文件夹下的文件夹
		var folderIds []int64
		if err := l.svcCtx.Engine.Select("id").In("parent_id", req.FolderIds).
			And("user_id =?", userId).
			And("del_flag=?", define.StatusFileUndeleted).Find(&folderIds); err != nil {
			logx.Errorf("查询文件夹下的文件夹失败:%v", err)
			return errors.New("查询文件夹下的文件夹失败")
		}

		req.FolderIds = folderIds
		for _, fileId := range fileIds {
			shareFiles = append(shareFiles, &models.ShareFile{
				ShareId: id,
				FileId:  fileId,
			})
		}
	}
	_, err = l.svcCtx.Engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if _, err := session.Insert(&shareFiles); err != nil {
			logx.Errorf("插入分享文件失败:%v", err)
			return nil, err
		}
		// 更新文件夹的分享状态
		created := time.Now().Local()
		expired := created.Unix() + define.ShareExpireType[req.ExpireType]
		share := &models.Share{}
		share.Id = id
		share.Pwd = req.Pwd
		share.Name = folderName
		share.UserId = userId
		share.Created = created
		share.Expired = expired
		share.Type = define.TypeShareMulti
		share.Url = url
		if _, err := session.Insert(share); err != nil {
			logx.Errorf("插入分享失败:%v", err)
			return nil, err
		}
		return nil, nil
	})

	return err
}
