package inspect

import (
	"bws_microservice_url/main/src/external"

	signup "github.com/bindways/bws_microservice_share/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type BwAuthService struct {
	token                *oauth2.Token
	microserviceCustomer *external.BwsMicroserviceCustomer
}

func (t *BwAuthService) Constructor() *BwAuthService {
	t.microserviceCustomer = new(external.BwsMicroserviceCustomer)
	return t
}

func (t *BwAuthService) GetToken(idCustomer primitive.ObjectID) (token *oauth2.Token, err error) {
	if t.token != nil && t.token.Valid() {
		return t.token, err
	}
	bwToken, err := t.microserviceCustomer.GetTokenFromCustomer(idCustomer)
	if err != nil {
		return
	}
	t.token = signup.ToToken(bwToken)
	return t.token, err
}
