package main

import (
	"github.com/jiangjinyuan/go-micro-demo/basic"
	"github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/handler"
	"github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/model"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	poolHashrate "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	//micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("btccom.explorer.srv.poolHashrate"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
		)

	// Register Handler
	poolHashrate.RegisterPoolHashrateHandler(service.Server(), new(handler.PoolHashrate))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("btccom.explorer.srv.poolHashrate", service.Server(), new(subscriber.PoolHashrate))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

/*func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}*/