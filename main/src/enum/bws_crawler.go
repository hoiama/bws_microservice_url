package enum

import "github.com/bindways/bw_microservice_share/bw_helper/bw_enum_helper"

type BwsCrawlerEnum string

const (
	/*
	 * Unknown user agent
	 */
	BwsCrawlerUndefined BwsCrawlerEnum = "CRAWLING_USER_AGENT_UNSPECIFIED"

	/*
	 * Desktop user agent..
	 */
	BwsCrawlerDesktop BwsCrawlerEnum = "DESKTOP"

	/*
	 * Mobile user agent.
	 */
	BwsCrawlerMobile BwsCrawlerEnum = "MOBILE"
)

// Convert a string to enum
func NewBwsCrawlerEnum(keyQueueName string) (BwsCrawlerEnum, error) {
	return bw_enum_helper.ToEnum[BwsCrawlerEnum](keyQueueName,
		BwsCrawlerUndefined, BwsCrawlerDesktop, BwsCrawlerMobile)
}
