package update

import (
	"bws_microservice_url/main/src/core/url_manager/index"
	"bws_microservice_url/main/src/core/url_manager/inspect"
	register "bws_microservice_url/main/src/core/url_manager/position"
	"bws_microservice_url/main/src/core/urls"
	"bws_microservice_url/main/src/dto"
	"bws_microservice_url/main/src/entity"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlUpdate struct {
	urlServiceDep         *urls.BwsUrlTrackService
	inspectServiceDep     *inspect.BwsInspectService
	urlIndexDep           *index.BwsUrlIndex
	urlPositionServiceDep *register.BwsUrlPositionService
}

func (t *BwsUrlUpdate) Constructor() *BwsUrlUpdate {
	t.urlServiceDep = new(urls.BwsUrlTrackService).Constructor1()
	t.inspectServiceDep = new(inspect.BwsInspectService).Constructor()
	t.urlIndexDep = new(index.BwsUrlIndex).Constructor()
	t.urlPositionServiceDep = new(register.BwsUrlPositionService).Constructor()
	return t
}

/**
 * Insert url track
 */
func (t *BwsUrlUpdate) UpdateUrl(idUrlTrack, idCustomer primitive.ObjectID) (urlTrack entity.BwsUrlTrack, err error) {
	urlTrack, err = t.urlServiceDep.GetUrlByI(idUrlTrack)
	if err != nil {
		return
	}
	if err = urlTrack.ValidIsOwner(idCustomer); err != nil {
		return
	}
	requestData, err := dto.NewBwsRequestData(urlTrack.Url)
	if err != nil {
		return
	}
	if err = t.hydrationWithConsoleData(idCustomer, requestData, &urlTrack); err != nil {
		return
	}
	err = t.urlServiceDep.UpdateUrl(urlTrack)
	return
}

/**
 * Get inspection data with track information in google
 * Get position performance in google
 */
func (t *BwsUrlUpdate) hydrationWithConsoleData(idCustomer primitive.ObjectID, requestData dto.BwsRequestData, urlTrack *entity.BwsUrlTrack) (err error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		engineTrack, err := t.inspectServiceDep.InspectURL(idCustomer, requestData)
		if err != nil {
			errChan <- fmt.Errorf("erro InspectURL: %w", err)
			return
		}
		urlTrack.UpdateData(engineTrack)
	}()

	go func() {
		defer wg.Done()
		keywordsPerformance, err := t.urlPositionServiceDep.GetPosition(idCustomer, requestData)
		if err != nil {
			errChan <- fmt.Errorf("erro GetPosition: %w", err)
			return
		}
		urlTrack.UpdateKeywordPerformance(keywordsPerformance)
	}()

	wg.Wait()
	close(errChan)
	if len(errChan) == 2 {
		for e := range errChan {
			if e != nil {
				err = e
			}
		}
	}
	return
}
