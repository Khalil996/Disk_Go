package upload

import (
	"cloud_go/Disk/common"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"cloud_go/Disk/internal/logic/upload"
	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.ParseForm(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		fileParam, err := common.GetSingleFile(r)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		if resp, err := l.Upload(&req, fileParam); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
