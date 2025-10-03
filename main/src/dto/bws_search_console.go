package dto

import (
	"google.golang.org/api/searchconsole/v1"
)

func NewInspectUrlIndexRequest(requestData BwsRequestData) *searchconsole.InspectUrlIndexRequest {
	return &searchconsole.InspectUrlIndexRequest{
		InspectionUrl: requestData.InspectionURL,
		SiteUrl:       requestData.SiteURL,
		LanguageCode:  "pt-BR",
	}
}
