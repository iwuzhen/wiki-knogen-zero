package translate

import (
	"context"

	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranslateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTranslateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranslateLogic {
	return &TranslateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TranslateLogic) Translate(req *types.TranslateQuery) (resp []string, err error) {
	// todo: add your logic here and delete this line
	resp = l.svcCtx.BaiduTranslate.TranslateSliceString(req.Query)
	return
}
