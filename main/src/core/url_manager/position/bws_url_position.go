package register

import (
	inspect2 "bws_microservice_url/main/src/core/search_console"
	"bws_microservice_url/main/src/core/urls"
	"bws_microservice_url/main/src/dto"
	"bws_microservice_url/main/src/entity"
	"math"
	"strings"
	"time"

	"github.com/bindways/bw_microservice_share/bw_helper/bw_time_helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/searchconsole/v1"
)

type BwsUrlPositionService struct {
	searchConsoleServiceDep *inspect2.BwsSearchConsoleService
	urlServiceDep           *urls.BwsUrlTrackService
}

func (t *BwsUrlPositionService) Constructor() *BwsUrlPositionService {
	t.searchConsoleServiceDep = new(inspect2.BwsSearchConsoleService).Constructor()
	t.urlServiceDep = new(urls.BwsUrlTrackService).Constructor1()
	return t
}

/**
 * Insert url track
 */
func (t *BwsUrlPositionService) GetPosition(idCustomer primitive.ObjectID,
	requestData dto.BwsRequestData) (keyPerformance []entity.BwsKeywordPerformance, err error) {

	searchConsoleService, err := t.searchConsoleServiceDep.GetSearchConsoleService(idCustomer)
	if err != nil {
		return
	}
	query := &searchconsole.SearchAnalyticsQueryRequest{

		StartDate:  bw_time_helper.ConvertDateToString(bw_time_helper.TimePlusXMonths(-1)),
		EndDate:    bw_time_helper.ConvertDateToString(time.Now()),
		Dimensions: []string{"query"},
		DimensionFilterGroups: []*searchconsole.ApiDimensionFilterGroup{
			{
				Filters: []*searchconsole.ApiDimensionFilter{
					{
						Dimension:  "page",
						Expression: requestData.InspectionURL,
					},
				},
			},
		},
		RowLimit: 2500,
	}
	resp, err := searchConsoleService.Searchanalytics.Query(requestData.SiteURL, query).Do()
	if err != nil {
		return
	}
	keyPerformance = t.createKeywordPerformance(resp)
	return
}

/**
 * Set data about google console
 */
func (t *BwsUrlPositionService) createKeywordPerformance(
	searchResponse *searchconsole.SearchAnalyticsQueryResponse) (keyPerformance []entity.BwsKeywordPerformance) {

	if len(searchResponse.Rows) == 0 {
		return
	}
	keyPerformance = make([]entity.BwsKeywordPerformance, len(searchResponse.Rows))
	for i, dataRow := range searchResponse.Rows {
		keyPerformance[i] = entity.BwsKeywordPerformance{
			Keyword:     strings.Join(dataRow.Keys, " "),
			Clicks:      dataRow.Clicks,
			Impressions: dataRow.Impressions,
			Ctr:         math.Round(dataRow.Ctr),
			Position:    math.Round(dataRow.Position),
		}
	}
	return
}
