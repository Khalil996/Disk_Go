package file

import (
	"net/http"

	"cloud_go/Disk/internal/logic/file"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateFoldersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateFoldersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewCreateFoldersLogic(r.Context(), svcCtx)
		err := l.CreateFolders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}