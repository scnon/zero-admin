package main

import (
	"flag"
	"fmt"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/rpc"

	"xlife/apps/auth/rpc/internal/config"
	menuServer "xlife/apps/auth/rpc/internal/server/menu"
	roleServer "xlife/apps/auth/rpc/internal/server/role"
	userServer "xlife/apps/auth/rpc/internal/server/user"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		auth.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		auth.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		auth.RegisterMenuServer(grpcServer, menuServer.NewMenuServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpc.LogInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
