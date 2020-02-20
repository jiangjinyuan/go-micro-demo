package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"time"

	poolHashrate "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
	"github.com/micro/go-micro/client"
)

func PoolHashrate(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Logf("request body: %v",request)
	// call the backend service
	poolHashrateClient := poolHashrate.NewPoolHashrateService("btccom.explorer.srv.poolHashrate", client.DefaultClient)
	rsp, err := poolHashrateClient.GetPoolHashrate(context.TODO(), &poolHashrate.Request{
		PoolID:   request["poolId"].(int32),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg":     rsp.PoolHashrate,
		"success": rsp.Success,
		"error":   rsp.Error,
		"ref":     time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
