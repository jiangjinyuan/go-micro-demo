package main

import (
        "github.com/micro/go-micro/util/log"
	"net/http"

        "github.com/micro/go-micro/web"
        "github.com/jiangjinyuan/go-micro-demo/poolHashrate-web/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("btccom.explorer.web.poolHashrate"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/poolHashrate/get", handler.PoolHashrate)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
