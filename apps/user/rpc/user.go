package main

import (
	"flag"
	"fmt"
	"github.com/zhanghongliang12/lebron/pkg/interceptor/rpcserver"

	"github.com/zhanghongliang12/lebron/apps/user/rpc/internal/config"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/internal/server"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/internal/svc"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "apps/user/rpc/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserRpcServer(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserRpcServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 添加rpc日志
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
