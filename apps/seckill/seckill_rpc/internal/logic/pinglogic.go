package logic

import (
	"context"

	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/internal/svc"
	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/seckill_rpc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *PingLogic) Ping(in *seckill_rpc.Request) (*seckill_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &seckill_rpc.Response{}, nil
}
