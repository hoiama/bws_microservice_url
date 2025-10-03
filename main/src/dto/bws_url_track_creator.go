package dto

import (
	"errors"
	"net/url"
)

type BwsUrlTrackCreator struct {
	SiteURL string `json:"siteUrl"`
}

func (t *BwsUrlTrackCreator) Valid() (err error) {
	if t.SiteURL == "" {
		return errors.New("Invalid site")
	}
	_, err = url.Parse(t.SiteURL)
	return
}
