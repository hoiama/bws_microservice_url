package a_test

import (
	"bws_microservice_url/main/src/core/url_manager/position"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBwsUrlRegisterService_GetPosition(t1 *testing.T) {
	urlPositionService := new(register.BwsUrlPositionService).Constructor()
	idCustomer, err := primitive.ObjectIDFromHex("68e00a3c785125ab3e97a467")
	assert.Nil(t1, err)
	idUrlTrack, err := primitive.ObjectIDFromHex("000000000000000000000000")
	assert.Nil(t1, err)
	_, err = urlPositionService.GetPosition(idCustomer, idUrlTrack)
	assert.Nil(t1, err)
}
