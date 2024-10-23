package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xlife/apps/merchant/api/internal/logic/boss/system/role"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
)

func RoleMenuIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleMenuIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleMenuIdsLogic(r.Context(), svcCtx)
		resp, err := l.RoleMenuIds(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
