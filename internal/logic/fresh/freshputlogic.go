package fresh

import (
	"context"

	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreshPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreshPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreshPutLogic {
	return &FreshPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreshPutLogic) FreshPut(req *types.FreshRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.svcCtx.FreshS.PutRecord(req.Path, req.Key, req.Data)
	resp = &types.Response{
		Code: 200,
	}
	return
}
