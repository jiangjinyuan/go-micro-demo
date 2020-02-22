package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangjinyuan/go-micro-demo/poolHashrate-web/handler"
)

func InitRoute() *gin.Engine{
	//create gin route
	s:=new(handler.Say)
	route:=gin.Default()
	route.POST("/poolHashrate",s.GetPoolHashrate)
	return route
}