package inspect

import (
	inspect "bws_microservice_url/main/src/core/auth"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"google.golang.org/api/indexing/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/searchconsole/v1"
)

type BwsSearchConsoleService struct {
	authService       *inspect.BwAuthService
	googleOauthConfig *oauth2.Config
}

func (t *BwsSearchConsoleService) Constructor() *BwsSearchConsoleService {
	t.authService = new(inspect.BwAuthService).Constructor()
	return t
}

func (t *BwsSearchConsoleService) getClient(idCustomer primitive.ObjectID) (client *http.Client, err error) {
	token, err := t.authService.GetToken(idCustomer)
	if err != nil {
		return
	}
	client = t.googleOauthConfig.Client(context.Background(), token)
	return
}

/**
 * Create search console service
 */
func (t *BwsSearchConsoleService) GetSearchConsoleService(idCustomer primitive.ObjectID) (service *searchconsole.Service, err error) {
	client, err := t.getClient(idCustomer)
	if err != nil {
		return
	}
	return searchconsole.NewService(context.Background(), option.WithHTTPClient(client))
}

/*
 * Create index console service
 */
func (t *BwsSearchConsoleService) GetIndexSearchConsoleService(idCustomer primitive.ObjectID) (service *indexing.Service, err error) {
	client, err := t.getClient(idCustomer)
	if err != nil {
		return
	}
	return indexing.NewService(context.Background(), option.WithHTTPClient(client))
}
