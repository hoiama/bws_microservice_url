package server

import (
	api2 "bws_microservice_url/main/src/api"

	"github.com/bindways/bw_microservice_share/bw_helper/bw_error_helper"
	"github.com/bindways/bw_microservice_share/bw_server"
	bw_handler_gin "github.com/bindways/bw_microservice_share2/bw_gin"
	"github.com/bindways/bws_microservice_share/router"
	"github.com/gin-gonic/gin"
)

func BwsServer() {
	engine := gin.Default()
	new(bw_handler_gin.BwServerGinMetric).
		Constructor1(engine).
		ConfigMiddleware().
		ConfigHandler(router.BwsMicroserviceUrl)
	new(api2.BwsUrlController).Constructor1().Controller(engine)
	bw_server.BwMicroservicePrintServer(&router.BwsMicroserviceUrl)
	err := engine.Run(router.BwsMicroserviceUrl.GetPortHttp())
	bw_error_helper.CheckPanic(err)
}
