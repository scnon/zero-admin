package main

import (
	"flag"
	"fmt"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/config"
	menuServer "zero-admin/apps/admin/rpc/internal/server/menu"
	roleServer "zero-admin/apps/admin/rpc/internal/server/role"
	userServer "zero-admin/apps/admin/rpc/internal/server/user"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		admin.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		admin.RegisterMenuServer(grpcServer, menuServer.NewMenuServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
