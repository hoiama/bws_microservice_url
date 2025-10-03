package enum

import "github.com/bindways/bw_microservice_share/bw_helper/bw_enum_helper"

// BwsPageStateEnum representa os possíveis estados de busca de página (fetch state)
type BwsPageStateEnum string

const (
	/*
	 * Unknown fetch state.
	 */
	BwsPageStateUndefined BwsPageStateEnum = "PAGE_FETCH_STATE_UNSPECIFIED"

	/*
	 * Successful fetch.
	 */
	BwsPageStateSuccess BwsPageStateEnum = "SUCCESSFUL"

	/*
	 * Soft 404.
	 */
	BwsPageStateSoft404 BwsPageStateEnum = "SOFT_404"

	/*
	 * Blocked by robots.txt.
	 */
	BwsPageStateBlockedRobotsTxt BwsPageStateEnum = "BLOCKED_ROBOTS_TXT"

	/*
	 * Not found (404).
	 */
	BwsPageStateNotFound BwsPageStateEnum = "NOT_FOUND"

	/*
	 * Blocked due to unauthorized request (401).
	 */
	BwsPageStateAccessDenied BwsPageStateEnum = "ACCESS_DENIED"

	/*
	 * Server error (5xx).
	 */
	BwsPageStateServerError BwsPageStateEnum = "SERVER_ERROR"

	/*
	 * Redirection error.
	 */
	BwsPageStateRedirectError BwsPageStateEnum = "REDIRECT_ERROR"

	/*
	 * Blocked due to access forbidden (403).
	 */
	BwsPageStateAccessForbidden BwsPageStateEnum = "ACCESS_FORBIDDEN"

	/*
	 * Blocked due to other 4xx issue (not 403, 404).
	 */
	BwsPageStateBlocked4xx BwsPageStateEnum = "BLOCKED_4XX"

	/*
	 * Internal error.
	 */
	BwsPageStateInternalCrawlError BwsPageStateEnum = "INTERNAL_CRAWL_ERROR"

	/*
	 * Invalid URL.
	 */
	BwsPageStateInvalidURL BwsPageStateEnum = "INVALID_URL"
)

// Convert a string to enum
func NewBwsPageStateEnum(keyQueueName string) (BwsPageStateEnum, error) {
	return bw_enum_helper.ToEnum[BwsPageStateEnum](keyQueueName,
		BwsPageStateUndefined,
		BwsPageStateSuccess,
		BwsPageStateSoft404,
		BwsPageStateBlockedRobotsTxt,
		BwsPageStateNotFound,
		BwsPageStateAccessDenied,
		BwsPageStateServerError,
		BwsPageStateRedirectError,
		BwsPageStateAccessForbidden,
		BwsPageStateBlocked4xx,
		BwsPageStateInternalCrawlError,
		BwsPageStateInvalidURL,
	)
}
