package entity

import (
	"bws_microservice_url/main/src/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlTracks []BwsUrlTrack
type BwsUrlTrack struct {

	/**
	 * idCustomer
	 */
	IdCustomer primitive.ObjectID `bson:"idCustomer" json:"idCustomer"`

	/**
	 * Url of site keyword
	 */
	Url string `bson:"url" json:"url"`

	/**
	 * keyword
	 */
	Keyword string `bson:"keyword" json:"keyword"`

	/**
	 * Volume search by month
	 */
	SearchVolume int `bson:"searchVolume" json:"searchVolume"`

	/**
	 * status of project
	 */
	Status enum.BwsUrlTrackStatusEnum `bson:"status" json:"status"`

	/**
	 * Date that was updated
	 */
	CreateAtDate time.Time `bson:"createAtDate" json:"createAtDate"`

	/**
	 * search engine data track
	 */
	BwsSearchEngineTrack BwsSearchEngineTrack `bson:"searchEngineTrack" json:"searchEngineTrack"`
}

func NewBwsUrlTrack(idCustomer primitive.ObjectID, url string, engineTrack BwsSearchEngineTrack) BwsUrlTrack {
	return BwsUrlTrack{
		IdCustomer:           idCustomer,
		Url:                  url,
		Status:               enum.BwsUrlTrackStatusActive,
		CreateAtDate:         time.Now(),
		BwsSearchEngineTrack: engineTrack,
	}
}

/**
 * Set date at current object
 */
func (t *BwsUrlTrack) UpdateDate() *BwsUrlTrack {
	t.CreateAtDate = time.Now()
	return t
}

/**
 * Build date at current object
 */
func (t *BwsUrlTrack) Build() BwsUrlTrack {
	return *t
}

/**
 * Check is url same
 */
func (t *BwsUrlTrack) IsUrl(url string) bool {
	return t.Url == url
}

/**
 * update track date today
 */
func (t *BwsUrlTracks) IsEmpty() (has bool) {
	return len(*t) == 0
}
