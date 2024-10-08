package store

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/apps/business/api/internal/logic/boss/store"
	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"
)

func DeleteStoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteStoreReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := store.NewDeleteStoreLogic(r.Context(), svcCtx)
		err := l.DeleteStore(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
