package fresh

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wiki-knogen-zero/internal/logic/fresh"
	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"
)

func FreshGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FreshRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := fresh.NewFreshGetLogic(r.Context(), svcCtx)
		resp, err := l.FreshGet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
