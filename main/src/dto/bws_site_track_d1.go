package dto

import (
	"bws_microservice_url/main/src/entity"
)

type BwsSiteTracksD1 []BwsSiteTrackD1

type BwsSiteTrackD1 struct {
	Keyword         string
	FindSiteToTrack bool
	entity.BwsUrlTrack
}

/**
 * Set link at current object
 */
func (t *BwsSiteTrackD1) SetKeyword(keyword string) *BwsSiteTrackD1 {
	t.Keyword = keyword
	return t
}

func (t *BwsSiteTrackD1) ToBwsSiteTrack() entity.BwsUrlTrack {
	return entity.BwsUrlTrack{
		Url:          t.Url,
		CreateAtDate: t.CreateAtDate,
	}
}

func (t *BwsSiteTracksD1) ToBwsSiteTracks() (siteTracks entity.BwsUrlTracks) {
	for _, d1 := range *t {
		siteTracks = append(siteTracks, d1.ToBwsSiteTrack())
	}
	return
}

/**
 * update track date today
 */
func (t *BwsSiteTracksD1) IsEmpty() (has bool) {
	return len(*t) == 0
}
