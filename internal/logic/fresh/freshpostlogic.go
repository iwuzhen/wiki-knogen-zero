package fresh

import (
	"context"

	"wiki-knogen-zero/internal/svc"
	"wiki-knogen-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreshPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreshPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreshPostLogic {
	return &FreshPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreshPostLogic) FreshPost(req *types.FreshRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// consecutive put

	ret, errE := l.svcCtx.FreshS.PutConsecutiveRecord(req.Path, req.Key, req.ID, req.Data)
	resp = &types.Response{
		Code:    200,
		Message: "ok",
		Data:    dataSchema{ID: ret.ID},
	}
	if errE != nil {
		resp.Code = 412
		resp.Message = errE.Error()
	}

	return
}
