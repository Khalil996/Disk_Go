package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFolderLogic {
	return &ListFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFolderLogic) ListFolder(req *types.ParentFolderIdReq) (resp []*types.ListFolderStruct, err error) {
	// todo: add your logic here and delete this line

	userId := l.ctx.Value(define.UserIdKey).(int64)
	var folders []*models.Folder
	if err := l.svcCtx.Engine.Where("parent_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).And("del_flag = ?",
		define.StatusFolderUndeleted).Find(&folders); err != nil {
		return nil, err
	}

	for _, folder := range folders {
		resp = append(resp, &types.ListFolderStruct{
			Id:      folder.Id,
			Name:    folder.Name,
			Updated: folder.UpdateAt.Format(define.TimeFormat1),
		})
	}

	return
}
