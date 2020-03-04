package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	k8s "github.com/micro/example/kubernetes/go/micro"
	"github.com/jiangjinyuan/go-micro-demo/poolHashrate/srv/handler"
	//"github.com/jiangjinyuan/go-micro-demo/poolHashrate/srv/subscriber"

	greeter "github.com/jiangjinyuan/go-micro-demo/poolHashrate/srv/proto/greeter"
)

func main() {
	// New Service
	service := k8s.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeter.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.greeter", service.Server(), new(subscriber.Greeter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
