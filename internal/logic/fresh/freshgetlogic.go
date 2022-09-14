package fresh

import (
	"context"

	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreshGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type dataSchema struct {
	ID   uint        `json:"id"`
	Data interface{} `json:"data"`
}

func NewFreshGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreshGetLogic {
	return &FreshGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreshGetLogic) FreshGet(req *types.FreshRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	ret := l.svcCtx.FreshS.GetLastestRecord(req.Path, req.Key)
	resp = &types.Response{
		Code:    200,
		Message: "ok",
		Data:    dataSchema{ret.ID, ret.Data},
	}
	return
}
