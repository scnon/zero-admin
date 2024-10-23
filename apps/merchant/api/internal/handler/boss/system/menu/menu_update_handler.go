package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xlife/apps/merchant/api/internal/logic/boss/system/menu"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
)

func MenuUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewMenuUpdateLogic(r.Context(), svcCtx)
		resp, err := l.MenuUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
