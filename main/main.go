package main

import (
	ses_database "bws_microservice_url/main/src/config/database"
	"bws_microservice_url/main/src/config/log"
	"bws_microservice_url/main/src/config/remote"
	"bws_microservice_url/main/src/config/server"
	"bws_microservice_url/main/src/core/runner"
	"time"

	"github.com/bindways/bw_microservice_share/bw_helper/bw_time_helper"
)

func main() {
	log.BwsLoggerClientConfig()
	remote.BwsConfigureFromRemoteServer()
	ses_database.BwsConnectDatabase()
	go bw_time_helper.ExecuteEachXPeriod(runner.ExecuteRunner, 24*time.Hour)
	server.BwsServer()
}
