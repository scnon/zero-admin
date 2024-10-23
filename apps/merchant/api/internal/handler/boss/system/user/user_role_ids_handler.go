package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xlife/apps/merchant/api/internal/logic/boss/system/user"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
)

func UserRoleIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRoleIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserRoleIdsLogic(r.Context(), svcCtx)
		resp, err := l.UserRoleIds(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
