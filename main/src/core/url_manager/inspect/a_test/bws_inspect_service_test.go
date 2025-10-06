package a_test

import (
	"bws_microservice_url/main/src/core/url_manager/inspect"
	"bws_microservice_url/main/src/dto"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBwInspectService_InspectURL(t1 *testing.T) {
	requestData := dto.BwsRequestData{
		SiteURL:       "sc-domain:partner.bindways.com",
		InspectionURL: "https://partner.bindways.com/blog/web/article/name/buscar-socios-encontrar-socio",
	}
	inspectService := new(inspect.BwsInspectService).Constructor()
	idCustomer, err := primitive.ObjectIDFromHex("68e00a3c785125ab3e97a467")
	inspectURL, err := inspectService.InspectURL(idCustomer, requestData)
	assert.Nil(t1, err)
	fmt.Println(inspectURL)
}
