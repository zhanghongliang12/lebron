package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhanghongliang12/lebron/apps/user/user_rpc/internal/svc"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
