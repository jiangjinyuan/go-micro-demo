package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangjinyuan/go-micro-demo/poolHashrate-web/handler"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/jiangjinyuan/go-micro-demo/poolHashrate-web/docs"
)

// @title Go-Micro-Demo API
// @version 1.0
// @description This is a explorer API
// @termsOfService https://btc.com
// @contract.name API Support
// @contract.url https://btc.com
// @contract.email support@btc.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8088
// @BasePath /
func InitRoute() *gin.Engine{
	//create gin route
	s:=new(handler.Say)
	route:=gin.Default()
	route.POST("/poolHashrate",s.GetPoolHashrate)
	//
	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json") // The url pointing to API definition
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,url))
	return route
}