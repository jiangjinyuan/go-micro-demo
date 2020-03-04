package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	greeter "github.com/jiangjinyuan/go-micro-demo/poolHashrate/srv/proto/greeter"
)

type Greeter struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Greeter) Call(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Log("Received Greeter.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

