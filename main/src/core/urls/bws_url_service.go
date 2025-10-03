package urls

import (
	"bws_microservice_url/main/src/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlTrackService struct {
	projectDep *BwsProjectDao
}

func (t *BwsUrlTrackService) Constructor1() *BwsUrlTrackService {
	t.projectDep = new(BwsProjectDao).Constructor1()
	return t
}

/**
 * Insert url track
 */
func (t *BwsUrlTrackService) InsertUrl(url entity.BwsUrlTrack) (err error) {
	return t.projectDep.insertUrl(url)
}

/**
 * Get url track by url
 */
func (t *BwsUrlTrackService) GetUrlTrack(url string) (exist bool, urlTrack entity.BwsUrlTrack, err error) {
	return t.projectDep.getUrlTrack(url)
}

/**
 * Get all url tracks
 */
func (t *BwsUrlTrackService) GetUrlTrackByCustomer(idCustomer primitive.ObjectID) (urlTrack []entity.BwsUrlTrack, err error) {
	return t.projectDep.getUrlTracks(idCustomer)
}
