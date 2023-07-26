package logic

import (
	"context"

	"github.com/zhanghongliang12/lebron/apps/cart/cart_rpc/cart_rpc"
	"github.com/zhanghongliang12/lebron/apps/cart/cart_rpc/internal/svc"

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

func (l *PingLogic) Ping(in *cart_rpc.Request) (*cart_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &cart_rpc.Response{
		Pong: "ok",
	}, nil
}
