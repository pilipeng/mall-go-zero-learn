package logic

import (
	"context"

	"mall-go-zero-learn.com/apps/app/internal/svc"
	"mall-go-zero-learn.com/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommentLogic {
	return &RecommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommentLogic) Recomment(req *types.RecommendRequest) (resp *types.RecommendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
