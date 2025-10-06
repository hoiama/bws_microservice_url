package index

import (
	inspect2 "bws_microservice_url/main/src/core/search_console"
	"bws_microservice_url/main/src/entity"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/indexing/v3"
)

type BwsUrlIndex struct {
	searchConsoleService *inspect2.BwsSearchConsoleService
}

func (t *BwsUrlIndex) Constructor() *BwsUrlIndex {
	t.searchConsoleService = new(inspect2.BwsSearchConsoleService).Constructor()
	return t
}

/**
 * Insert url track
 */
func (t *BwsUrlIndex) requestIndexUrl(idCustomer primitive.ObjectID, urlToNotify string) (err error) {
	searchConsoleService, err := t.searchConsoleService.GetIndexSearchConsoleService(idCustomer)
	if err != nil {
		return
	}
	notification := &indexing.UrlNotification{
		Url:  urlToNotify,
		Type: "URL_UPDATED",
	}
	_, err = searchConsoleService.UrlNotifications.Publish(notification).Do()
	return
}

func (t *BwsUrlIndex) RequestIndex(urlTrack entity.BwsUrlTrack) {
	if urlTrack.SearchEngineTrack.IsUntracked() {
		if err := t.requestIndexUrl(urlTrack.IdCustomer, urlTrack.Url); err != nil {
			log.Err(err).Msg("error to index url")
		}
	}
}
