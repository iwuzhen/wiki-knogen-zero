// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	fresh "wiki-knogen-zero/internal/handler/fresh"
	translate "wiki-knogen-zero/internal/handler/translate"
	"wiki-knogen-zero/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/translate/zh",
				Handler: translate.TranslateHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/fresh/:path/:key",
				Handler: fresh.FreshPutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/fresh/:path/:key",
				Handler: fresh.FreshPostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/fresh/:path/:key",
				Handler: fresh.FreshGetHandler(serverCtx),
			},
		},
	)
}
