package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"cloud_go/Disk/internal/logic/file"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListFileMovableFolderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ParentFolderIdReq
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewListFileMovableFolderLogic(r.Context(), svcCtx)
		if resp, err := l.ListFileMovableFolder(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
