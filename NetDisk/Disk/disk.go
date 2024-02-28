package main

import (
	"cloud_go/Disk/internal/config"
	"cloud_go/Disk/internal/handler"
	"cloud_go/Disk/internal/middleware"
	"cloud_go/Disk/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/disk-api.yaml", "the config file")

func main() {

	//日志
	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"

	logc.MustSetup(cfg)
	defer logc.Close()
	logc.Info(context.Background(), "hello world")

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.HandleCors)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
