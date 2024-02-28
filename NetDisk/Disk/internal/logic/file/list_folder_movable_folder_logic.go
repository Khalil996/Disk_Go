package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"errors"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFolderMovableFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFolderMovableFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFolderMovableFolderLogic {
	return &ListFolderMovableFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFolderMovableFolderLogic) ListFolderMovableFolder(req *types.ListFolderMovableFolderReq) (resp []*types.ListFolderStruct, err error) {
	// todo: add your logic here and delete this line

	userId := l.ctx.Value(define.UserIdKey).(int64)
	parentFolderId := req.ParentFolderId
	var (
		folders     []*models.Folder
		selectedMap = make(map[int64]struct{})
	)

	if err := l.svcCtx.Engine.Cols("id", "name").
		Where("parent_id = ?", parentFolderId).
		And("user_id = ?", userId).Find(&folders); err != nil {
		return nil, errors.New("出错了" + err.Error())
	}

	for _, selectedId := range req.SelectedFolderIds {
		selectedMap[selectedId] = struct{}{}
	}

	for _, folder := range folders {
		if _, ok := selectedMap[folder.Id]; ok {
			continue
		}
		lfs := &types.ListFolderStruct{}
		lfs.Id = folder.Id
		lfs.Name = folder.Name
		resp = append(resp, lfs)
	}
	return
}
