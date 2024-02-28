package download

import (
	"net/http"

	"cloud_go/Disk/internal/logic/download"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChunkDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChunkDownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := download.NewChunkDownloadLogic(r.Context(), svcCtx)
		err := l.ChunkDownload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
