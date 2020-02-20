package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	poolHashrate "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
)

func PoolHashrate(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	poolHashrateClient := poolHashrate.NewPoolHashrateService("btccom.explorer.srv.poolHashrate", client.DefaultClient)
	rsp, err := poolHashrateClient.GetPoolHashrate(context.TODO(), &poolHashrate.Request{
		PoolID: request["poolId"].(int32),
		PoolName: request["poolName"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.PoolHashrate,
		"success": rsp.Success,
		"error": rsp.Error,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
