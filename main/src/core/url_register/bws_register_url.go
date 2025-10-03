package url_register

import (
	"bws_microservice_url/main/src/core/inspect"
	"bws_microservice_url/main/src/core/urls"
	"bws_microservice_url/main/src/dto"
	"bws_microservice_url/main/src/entity"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlRegisterService struct {
	urlService     *urls.BwsUrlTrackService
	inspectService *inspect.BwsInspectService
}

func (t *BwsUrlRegisterService) Constructor() *BwsUrlRegisterService {
	t.urlService = new(urls.BwsUrlTrackService).Constructor1()
	t.inspectService = new(inspect.BwsInspectService).Constructor()
	return t
}

/**
 * Insert url track
 */
func (t *BwsUrlRegisterService) RegisterUrl(idCustomer primitive.ObjectID, trackCreator dto.BwsUrlTrackCreator) (urlTrack entity.BwsUrlTrack, err error) {
	if err = trackCreator.Valid(); err != nil {
		return
	}
	exist, _, err := t.urlService.GetUrlTrack(trackCreator.SiteURL)
	if err != nil {
		return
	}
	if exist {
		return urlTrack, errors.New("url already registered")
	}
	requestData, err := dto.NewBwsRequestData(trackCreator.SiteURL)
	if err != nil {
		return
	}
	engineTrack, err := t.inspectService.InspectURL(idCustomer, requestData)
	if err != nil {
		return
	}
	urlTrack = entity.NewBwsUrlTrack(idCustomer, trackCreator.SiteURL, engineTrack)
	err = t.urlService.InsertUrl(urlTrack)
	return
}
