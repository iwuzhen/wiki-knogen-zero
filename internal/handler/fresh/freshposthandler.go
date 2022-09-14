package fresh

import (
	"net/http"

	"wiki-knogen-zero/internal/logic/fresh"
	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FreshPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FreshRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := fresh.NewFreshPostLogic(r.Context(), svcCtx)
		resp, err := l.FreshPost(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
