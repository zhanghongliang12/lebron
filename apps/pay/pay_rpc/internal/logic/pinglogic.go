package logic

import (
	"context"

	"github.com/zhanghongliang12/lebron/apps/pay/pay_rpc/internal/svc"
	"github.com/zhanghongliang12/lebron/apps/pay/pay_rpc/pay_rpc"

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

func (l *PingLogic) Ping(in *pay_rpc.Request) (*pay_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &pay_rpc.Response{}, nil
}
