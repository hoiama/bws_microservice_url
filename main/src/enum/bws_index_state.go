package enum

import "github.com/bindways/bw_microservice_share/bw_helper/bw_enum_helper"

type BwsIndexStateEnum string

const (
	/*
	 * Unknown indexing status.
	 */
	BwsIndexStateUndefined BwsIndexStateEnum = "INDEXING_STATE_UNSPECIFIED"

	/*
	 * Indexing allowed.
	 */
	BwsIndexStateAllowed BwsIndexStateEnum = "INDEXING_ALLOWED"

	/*
	 * Indexing not allowed, 'noindex' detected in 'robots' meta tag.
	 */
	BwsIndexStateBlockMeta BwsIndexStateEnum = "BLOCKED_BY_META_TAG"

	/*
	 * Indexing not allowed, 'noindex' detected in 'X-Robots-Tag' http header.
	 */
	BwsIndexStateBlockHeader BwsIndexStateEnum = "BLOCKED_BY_HTTP_HEADER"

	/*
	 * Reserved, no longer in use.
	 */
	BwsIndexStateBlockTxt BwsIndexStateEnum = "BLOCKED_BY_ROBOTS_TXT"
)

// Convert a string to enum
func NewBwsIndexStateEnum(keyQueueName string) (BwsIndexStateEnum, error) {
	return bw_enum_helper.ToEnum[BwsIndexStateEnum](keyQueueName,
		BwsIndexStateUndefined, BwsIndexStateAllowed, BwsIndexStateBlockMeta, BwsIndexStateBlockHeader, BwsIndexStateBlockTxt)
}
