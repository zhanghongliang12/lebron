// Code generated by goctl. DO NOT EDIT.
// Source: seckill_rpc.proto

package seckillrpcclient

import (
	"context"

	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/seckill_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = seckill_rpc.Request
	Response = seckill_rpc.Response

	SeckillRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultSeckillRpc struct {
		cli zrpc.Client
	}
)

func NewSeckillRpc(cli zrpc.Client) SeckillRpc {
	return &defaultSeckillRpc{
		cli: cli,
	}
}

func (m *defaultSeckillRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := seckill_rpc.NewSeckillRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
