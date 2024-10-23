package dept

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xlife/apps/merchant/api/internal/logic/boss/system/dept"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
)

func DeptUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeptUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewDeptUpdateLogic(r.Context(), svcCtx)
		resp, err := l.DeptUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
