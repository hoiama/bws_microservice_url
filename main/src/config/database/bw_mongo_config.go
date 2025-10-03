package ses_database

import (
	"context"

	"github.com/bindways/bw_microservice_share/bw_database"
	"go.mongodb.org/mongo-driver/mongo"
)

var ContextTODO = context.TODO()
var BwsProjectDatabaseUrl = "mongodb://localhost:27017/"

func BwsProjectCollection() *mongo.Collection {
	return BwsConnectDatabase().Collection("url_track")
}

/**
 * Connect with database
 */
func BwsConnectDatabase() *mongo.Database {
	return BwsConnectDatabase2("bws_microservice_url")
}

func BwsConnectDatabase2(databaseName string) *mongo.Database {
	return bw_database.ConnectDatabaseMongo(BwsProjectDatabaseUrl, databaseName, ContextTODO)
}
