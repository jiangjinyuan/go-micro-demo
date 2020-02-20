package handler

import (
	"context"
	pool "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/model/poolHashrate"
	poolHashrate "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
	"github.com/micro/go-micro/util/log"
)

type PoolHashrate struct{}

var (
	poolHashrateService pool.Service
)

func (e *PoolHashrate) GetPoolHashrate(ctx context.Context, req *poolHashrate.Request, rsp *poolHashrate.Response) error {
	Hashrate, err := poolHashrateService.GetPoolHashrate(req.PoolID)
	if err != nil {
		rsp.Success = false
		rsp.Error = &poolHashrate.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.PoolHashrate = Hashrate
	rsp.Success = true
	return nil
}

// Init 初始化handler
func Init() {
	var err error
	poolHashrateService, err = pool.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}
