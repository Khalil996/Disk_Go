package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"cloud_go/Disk/internal/logic/file"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CopyFoldersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CopyFoldersReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
			return
		}

		l := file.NewCopyFoldersLogic(r.Context(), svcCtx)
		if err := l.CopyFolders(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
