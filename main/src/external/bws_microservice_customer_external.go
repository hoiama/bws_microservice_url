package external

import (
	"encoding/json"
	"fmt"

	"github.com/bindways/bw_microservice_share/bw_helper/bw_feign_client_helper"
	"github.com/bindways/bw_microservice_share/bw_microservice/bw_microservice_oauth2/bw_entity"
	"github.com/bindways/bws_microservice_share/cons"
	"github.com/bindways/bws_microservice_share/router"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsMicroserviceCustomer struct {
}

/**
 * Get all articles from project and limited size
 */
func (t *BwsMicroserviceCustomer) GetTokenFromCustomer(idCustomer primitive.ObjectID) (bwToken bw_entity.BwTokenJWT, err error) {
	urlFill := fmt.Sprintf("%s/%s/google/token", router.BwsMicroserviceCustomer.GetLocalFullRouterHttp(), idCustomer.Hex())
	responseBytes, err := bw_feign_client_helper.BwGetWithParams(urlFill, nil, cons.BwsMasterToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBytes, &bwToken)
	return
}
