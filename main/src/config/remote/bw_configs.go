package remote

import (
	ses_database "bws_microservice_url/main/src/config/database"
	"github.com/bindways/bw_microservice_share/bw_flag"
	"github.com/bindways/bw_microservice_share/bw_helper/bw_error_helper"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_config/bw_external"
)

func BwsConfigureFromRemoteServer() {
	environment := bw_flag.GetArgEnvironment()
	viper, err := bw_external.ConfigureEnvironment("bw_microservice_config_share", environment)
	if err != nil {
		bw_error_helper.CheckPanic(err)
	}
	ses_database.BwsProjectDatabaseUrl = viper.GetString("database.uri")
}

/**
 * Call to config environment to make tests
 */
func BwsConfigureEnvironmentTest() {
	ses_database.BwsProjectDatabaseUrl = "mongodb://localhost:27017/"
}
