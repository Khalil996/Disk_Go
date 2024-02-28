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

type ListFileMovableFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileMovableFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileMovableFolderLogic {
	return &ListFileMovableFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileMovableFolderLogic) ListFileMovableFolder(req *types.ParentFolderIdReq) ([]*types.ListFolderStruct, error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value(define.UserIdKey).(int64)
	var folders []*models.Folder
	if err := l.svcCtx.Engine.Cols("id", "name").Where("parent_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).Find(&folders); err != nil {
		return nil, errors.New("出错了" + err.Error())
	}

	var resp []*types.ListFolderStruct
	for _, folder := range folders {
		lfs := &types.ListFolderStruct{}
		lfs.Id = folder.Id
		lfs.Name = folder.Name
		resp = append(resp, lfs)
	}
	return resp, nil
}
