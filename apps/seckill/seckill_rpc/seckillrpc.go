package main

import (
	"flag"
	"fmt"

	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/internal/config"
	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/internal/server"
	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/internal/svc"
	"github.com/zhanghongliang12/lebron/apps/seckill/seckill_rpc/seckill_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/seckillrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		seckill_rpc.RegisterSeckillRpcServer(grpcServer, server.NewSeckillRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}