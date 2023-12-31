// Code generated by goctl. DO NOT EDIT.
// Source: cart_rpc.proto

package server

import (
	"context"

	"github.com/zhanghongliang12/lebron/apps/cart/cart_rpc/cart_rpc"
	"github.com/zhanghongliang12/lebron/apps/cart/cart_rpc/internal/logic"
	"github.com/zhanghongliang12/lebron/apps/cart/cart_rpc/internal/svc"
)

type CartRpcServer struct {
	svcCtx *svc.ServiceContext
	cart_rpc.UnimplementedCartRpcServer
}

func NewCartRpcServer(svcCtx *svc.ServiceContext) *CartRpcServer {
	return &CartRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *CartRpcServer) Ping(ctx context.Context, in *cart_rpc.Request) (*cart_rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
