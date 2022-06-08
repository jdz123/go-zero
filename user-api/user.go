package main

import (
	"flag"
	"fmt"
	"go-zore/common/middleware"
	"go-zore/user-api/internal/config"
	"go-zore/user-api/internal/handler"
	"go-zore/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//获取配置文件
var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) // 将yaml 文件解析到config结构体里

	ctx := svc.NewServiceContext(c) //返回包括各个model的service 结构体
	fmt.Println(ctx)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.NewCommonJwtAuthMiddleware().handle)

	handler.RegisterHandlers(server, ctx) //注册路由

	logx.DisableStat() //去掉监控日志记录
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
