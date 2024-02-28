package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"cloud_go/Disk/internal/logic/file"
	"cloud_go/Disk/internal/svc"
)

func ListDeletedFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewListDeletedFilesLogic(r.Context(), svcCtx)
		resp, err := l.ListDeletedFiles()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
