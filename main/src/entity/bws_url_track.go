package entity

import (
	"bws_microservice_url/main/src/enum"
	"time"

	"github.com/bindways/bw_microservice_share/bw_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlTracks []BwsUrlTrack

// Check order of elements
func (t BwsUrlTracks) Less(i, j int) bool {
	return t[i].CreateAtDate.After(t[j].CreateAtDate)
}

// size of elements, interface sort.Interface to invites
func (t BwsUrlTracks) Len() int {
	return len(t)
}

// change elements of position, interface sort.Interface to alter position of invites
func (t BwsUrlTracks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type BwsUrlTrack struct {

	/**
	 * idCustomer
	 */
	Id primitive.ObjectID `bson:"_id" json:"id"`

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
	SearchEngineTrack BwsSearchEngineTrack `bson:"searchEngineTrack" json:"searchEngineTrack"`
}

func NewBwsUrlTrack(idCustomer primitive.ObjectID, url string, engineTrack BwsSearchEngineTrack) BwsUrlTrack {
	return BwsUrlTrack{
		Id:                primitive.NewObjectID(),
		IdCustomer:        idCustomer,
		Url:               url,
		Status:            enum.BwsUrlTrackStatusActive,
		CreateAtDate:      time.Now(),
		SearchEngineTrack: engineTrack,
	}
}

/**
 * Update data rul track
 */
func (t *BwsUrlTrack) UpdateData(searchEngineTrack BwsSearchEngineTrack) *BwsUrlTrack {
	t.SearchEngineTrack.UpdateData(searchEngineTrack)
	return t
}

/**
 * Update data keywordsPerformance
 */
func (t *BwsUrlTrack) UpdateKeywordPerformance(keywordsPerformance BwsKeywordsPerformance) *BwsUrlTrack {
	t.SearchEngineTrack.updateKeywordPerformance(keywordsPerformance)
	return t
}

/**
 * Check if is owner or profile
 */
func (t *BwsUrlTrack) ValidIsOwner(idCustomer primitive.ObjectID) error {
	if t.IdCustomer != idCustomer {
		return bw_error.BwErrorNotOwnerResource.Pt()
	}
	return nil
}
