package user

import (
	"cloud_go/Disk/common"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"cloud_go/Disk/internal/logic/user"
	"cloud_go/Disk/internal/svc"
)

func UpdateAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileParam, err := common.GetSingleFile(r)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}

		l := user.NewUpdateAvatarLogic(r.Context(), svcCtx)
		if resp, err := l.UpdateAvatar(fileParam); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
