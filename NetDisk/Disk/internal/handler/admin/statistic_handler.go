package admin

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"cloud_go/Disk/internal/logic/admin"
	"cloud_go/Disk/internal/svc"
)

func StatisticHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewStatisticLogic(r.Context(), svcCtx)

		if resp, err := l.Statistic(); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
