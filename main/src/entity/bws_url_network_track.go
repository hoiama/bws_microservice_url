package entity

import (
	"bws_microservice_url/main/src/enum"
	"time"
)

type BwsSearchEngineTrack struct {

	/**
	 * Search engine used
	 */
	BwsSearchEngine enum.BwSearchEngineEnum `bson:"searchEngine" json:"searchEngine"`

	/**
	 * Current position of site in google
	 */
	CurrentPosition int `bson:"currentPosition" json:"currentPosition"`

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
}
