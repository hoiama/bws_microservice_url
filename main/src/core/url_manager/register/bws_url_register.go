package register

import (
	"bws_microservice_url/main/src/core/url_manager/index"
	"bws_microservice_url/main/src/core/url_manager/inspect"
	"bws_microservice_url/main/src/core/urls"
	"bws_microservice_url/main/src/dto"
	"bws_microservice_url/main/src/entity"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlRegisterService struct {
	urlServiceDep      *urls.BwsUrlTrackService
	inspectServiceDep  *inspect.BwsInspectService
	urlIndexServiceDep *index.BwsUrlIndex
}

func (t *BwsUrlRegisterService) Constructor() *BwsUrlRegisterService {
	t.urlServiceDep = new(urls.BwsUrlTrackService).Constructor1()
	t.inspectServiceDep = new(inspect.BwsInspectService).Constructor()
	t.urlIndexServiceDep = new(index.BwsUrlIndex).Constructor()
	return t
}

/**
 * Insert url track
 */
func (t *BwsUrlRegisterService) RegisterUrl(idCustomer primitive.ObjectID, trackCreator dto.BwsUrlTrackCreator) (urlTrack entity.BwsUrlTrack, err error) {
	if err = trackCreator.Valid(); err != nil {
		return
	}
	exist, urlTrack, err := t.urlServiceDep.GetUrlTrack(trackCreator.SiteURL)
	if err != nil {
		return
	}
	if exist {
		if err = urlTrack.ValidIsOwner(idCustomer); err != nil {
			return
		}
		return urlTrack, errors.New("url already registered")
	}

	requestData, err := dto.NewBwsRequestData(trackCreator.SiteURL)
	if err != nil {
		return
	}
	engineTrack, err := t.inspectServiceDep.InspectURL(idCustomer, requestData)
	if err != nil {
		return
	}
	urlTrack = entity.NewBwsUrlTrack(idCustomer, trackCreator.SiteURL, engineTrack)
	go t.urlIndexServiceDep.RequestIndex(urlTrack)
	err = t.urlServiceDep.InsertUrl(urlTrack)
	return
}
