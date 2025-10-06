package inspect

import (
	inspect2 "bws_microservice_url/main/src/core/search_console"
	"bws_microservice_url/main/src/dto"
	"bws_microservice_url/main/src/entity"
	"bws_microservice_url/main/src/enum"
	"errors"
	"time"

	"github.com/bindways/bw_microservice_share/bw_helper/bw_time_helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/searchconsole/v1"
)

type BwsInspectService struct {
	searchConsoleService *inspect2.BwsSearchConsoleService
}

func (t *BwsInspectService) Constructor() *BwsInspectService {
	t.searchConsoleService = new(inspect2.BwsSearchConsoleService).Constructor()
	return t
}

func (t *BwsInspectService) InspectURL(idCustomer primitive.ObjectID,
	requestData dto.BwsRequestData) (engineTrack entity.BwsSearchEngineTrack, err error) {

	if err = requestData.Valid(); err != nil {
		return
	}
	searchConsoleService, err := t.searchConsoleService.GetSearchConsoleService(idCustomer)
	if err != nil {
		return engineTrack, errors.New("cant create google search service")
	}
	inspectUrlIndexRequest := dto.NewInspectUrlIndexRequest(requestData)
	resp, err := searchConsoleService.UrlInspection.Index.Inspect(inspectUrlIndexRequest).Do()
	return t.SetInspectURL(resp)
}

func (t *BwsInspectService) SetInspectURL(resp *searchconsole.InspectUrlIndexResponse) (data entity.BwsSearchEngineTrack, err error) {
	data = entity.BwsSearchEngineTrack{}
	if resp.InspectionResult.IndexStatusResult.CrawledAs != "" {
		if crawler, err := enum.NewBwsCrawlerEnum(resp.InspectionResult.IndexStatusResult.CrawledAs); err != nil {
			data.Crawler = crawler
		}
	}
	indexingState, err := enum.NewBwsIndexStateEnum(resp.InspectionResult.IndexStatusResult.IndexingState)
	if err != nil {
		return
	}
	verdict, err := enum.NewBwsVerdictEnum(resp.InspectionResult.IndexStatusResult.Verdict)
	if err != nil {
		return
	}
	pageFetchState, err := enum.NewBwsPageStateEnum(resp.InspectionResult.IndexStatusResult.PageFetchState)
	if err != nil {
		return
	}
	if resp.InspectionResult.IndexStatusResult.LastCrawlTime != "" {
		if crawlTime, err := bw_time_helper.ParseLastCrawlTime(resp.InspectionResult.IndexStatusResult.LastCrawlTime); err != nil {
			data.LastCrawlDate = crawlTime
		}
	}
	data.SearchEngine = enum.BwsSearchEngineGoogle
	data.IndexState = indexingState
	data.Verdict = verdict
	data.VerdictDetail = resp.InspectionResult.IndexStatusResult.CoverageState
	data.PageState = pageFetchState
	data.Canonical = resp.InspectionResult.IndexStatusResult.GoogleCanonical
	data.ReferringUrls = resp.InspectionResult.IndexStatusResult.ReferringUrls
	data.UpdateAtDate = time.Now()
	return
}
