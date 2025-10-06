package entity

import (
	"bws_microservice_url/main/src/enum"
	"time"
)

type BwsSearchEngineTrack struct {

	/**
	 * Search engine used
	 */
	SearchEngine enum.BwSearchEngineEnum `bson:"searchEngine" json:"searchEngine"`

	/**
	 * Current last of site in google
	 */
	LastPosition int `bson:"lastPosition" json:"lastPosition"`

	/**
	 * Primary crawler that was used by Google to crawl your site.
	 */
	Crawler enum.BwsCrawlerEnum `bson:"crawler" json:"crawler"`

	/**
	 * Primary crawler that was used by Google to crawl your site.
	 */
	IndexState enum.BwsIndexStateEnum `bson:"indexState" json:"indexState"`

	/**
	 * verdict can be PASS, PARTIAL and FAIL
	 */
	Verdict enum.BwsVerdictEnum `bson:"verdict" json:"verdict"`

	/**
	 * Could Google find and index the page. More details about page
	 */
	VerdictDetail string `bson:"verdictDetail" json:"verdictDetail"`

	/**
	 * Whether or not Google could retrieve the page from your
	 */
	PageState enum.BwsPageStateEnum `bson:"pageState" json:"pageState"`

	/**
	 * Canonical page
	 */
	Canonical string `bson:"canonical" json:"canonical"`

	/**
	 * URLs that link to the inspected URL, directly and indirectly.
	 */
	ReferringUrls []string `json:"referringUrls,omitempty"`

	/**
	 * Date that was crawled
	 */
	LastCrawlDate time.Time `bson:"lastCrawlDate" json:"lastCrawlDate"`

	/**
	 * Date that was updated
	 */
	UpdateAtDate time.Time `bson:"updateAtDate" json:"updateAtDate"`

	/**
	 * Store keyword data about url
	 */
	KeywordsPerformance BwsKeywordsPerformance `bson:"keywordsPerformance" json:"keywordsPerformance"`
}

/**
 * Update data rul track
 */
func (t *BwsSearchEngineTrack) UpdateData(searchEngineTrack BwsSearchEngineTrack) *BwsSearchEngineTrack {
	t.SearchEngine = searchEngineTrack.SearchEngine
	t.Crawler = searchEngineTrack.Crawler
	t.IndexState = searchEngineTrack.IndexState
	t.Verdict = searchEngineTrack.Verdict
	t.VerdictDetail = searchEngineTrack.VerdictDetail
	t.PageState = searchEngineTrack.PageState
	t.Canonical = searchEngineTrack.Canonical
	t.ReferringUrls = searchEngineTrack.ReferringUrls
	t.LastCrawlDate = searchEngineTrack.LastCrawlDate
	t.UpdateAtDate = time.Now()
	return t
}

/**
 * Update data keywordsPerformance
 */
func (t *BwsSearchEngineTrack) updateKeywordPerformance(keywordsPerformance BwsKeywordsPerformance) *BwsSearchEngineTrack {
	t.KeywordsPerformance = keywordsPerformance
	return t
}

/**
 * Check if is owner or profile
 */
func (t *BwsSearchEngineTrack) IsTracked() bool {
	return t.Verdict == enum.BwsVerdictPass
}

/**
 * Check if is owner or profile
 */
func (t *BwsSearchEngineTrack) IsUntracked() bool {
	return t.Verdict != enum.BwsVerdictPass
}
