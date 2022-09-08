package translate

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wiki-knogen-zero/internal/logic/translate"
	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"
)

func TranslateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TranslateQuery
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := translate.NewTranslateLogic(r.Context(), svcCtx)
		resp, err := l.Translate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
