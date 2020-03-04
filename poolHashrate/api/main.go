package main

import (
        "github.com/micro/go-micro/util/log"
	    "net/http"
        k8s "github.com/micro/examples/kubernetes/go/web"
        "github.com/micro/go-micro/web"
        "github.com/jiangjinyuan/go-micro-demo/poolHashrate/api/handler"
)

func main() {
	// create new web service
        service := k8s.NewService(
                web.Name("go.micro.web.greeter"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/greeter/call", handler.GreeterCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
