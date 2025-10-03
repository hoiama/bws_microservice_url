package dto

import (
	"errors"
	"fmt"
	"net/url"
)

type BwsRequestData struct {
	SiteURL       string `json:"siteUrl"`
	InspectionURL string `json:"inspectionUrl"`
}

func (t *BwsRequestData) Valid() (err error) {
	if t.SiteURL == "" {
		return errors.New("Invalid site")
	}
	if t.InspectionURL == "" {
		return errors.New("Invalid inspection url")
	}
	return
}

func NewBwsRequestData(urlSite string) (requestData BwsRequestData, err error) {
	urlParsed, err := url.Parse(urlSite)
	if err != nil {
		return
	}
	requestData = BwsRequestData{
		SiteURL:       fmt.Sprintf("sc-domain:%s", urlParsed.Hostname()),
		InspectionURL: urlSite,
	}
	return
}
