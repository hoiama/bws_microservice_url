package log

import (
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_logger/bw_logger"
	router2 "github.com/bindways/bws_microservice_share/router"
)

func BwsLoggerClientConfig() {
	new(bw_logger.LogServerClient).Constructor(router2.BwsMicroserviceUrl.NameMicroservice).Connect()
}
