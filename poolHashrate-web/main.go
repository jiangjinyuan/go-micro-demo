package main

import (
	"github.com/jiangjinyuan/go-micro-demo/poolHashrate-web/route"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	"github.com/jiangjinyuan/go-micro-demo/poolHashrate-web/handler"
	"github.com/micro/go-micro/web"
	poolHashrate "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("btccom.explorer.web.poolHashrate"),
		web.Version("latest"),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	//连接srv服务，set up srv client
	handler.PoolHashrateClient = poolHashrate.NewPoolHashrateService("btccom.explorer.srv.poolHashrate", client.DefaultClient)

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//service.HandleFunc("/poolHashrate", handler.PoolHashrate)
	service.Handle("/",route.InitRoute())

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
