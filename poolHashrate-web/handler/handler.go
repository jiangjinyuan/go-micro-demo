package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	poolHashrate "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
	"github.com/micro/go-micro/util/log"
)

var (
	PoolHashrateClient poolHashrate.PoolHashrateService
)

type Say struct {

}

/*func PoolHashrate(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Logf("request body: %v",request)

	poolId := request["poolId"].(string)

	// call the backend service
	//poolHashrateClient := poolHashrate.NewPoolHashrateService("btccom.explorer.srv.poolHashrate", client.DefaultClient)
	rsp, err := PoolHashrateClient.GetPoolHashrate(context.TODO(), &poolHashrate.Request{
		PoolID:   poolId,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Logf("error",err)
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
}*/

func (s *Say) GetPoolHashrate(ctx *gin.Context){
	poolId := ctx.PostForm("poolId")
	log.Logf("request poolId: %s",poolId)
	rsp, err := PoolHashrateClient.GetPoolHashrate(context.TODO(), &poolHashrate.Request{
		PoolID:   poolId,
	})
	if err!=nil{
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200,rsp)
}
