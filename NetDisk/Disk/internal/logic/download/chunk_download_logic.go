package download

import (
	"context"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChunkDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChunkDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChunkDownloadLogic {
	return &ChunkDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChunkDownloadLogic) ChunkDownload(req *types.ChunkDownloadReq) error {
	// todo: add your logic here and delete this line

	return nil
}
