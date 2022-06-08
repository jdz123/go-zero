package main

import (
	"flag"
	"fmt"

	"go-zore/user-rpc/internal/config"
	"go-zore/user-rpc/internal/server"
	"go-zore/user-rpc/internal/svc"
	"go-zore/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) //将配置文件映射给config
	ctx := svc.NewServiceContext(c)
	svr := server.NewUsercenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//拦截器
	s.AddUnaryInterceptors((fune(ctx context.Context，req interface(l，info *grpc. UnaryServerInfo， handler gpc. UnaryHandler)(resp interface {},err error){

	})

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}


func TestServerInterceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error) {
	fmt.Printf("TestServerInterceptor====>")
	return handler(ctx,req)
}
