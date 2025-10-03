package enum

import "github.com/bindways/bw_microservice_share/bw_helper/bw_enum_helper"

type BwsVerdictEnum string

const (
	/*
	 * Equivalent to "Valid" for the page or item in Search Console.
	 */
	BwsVerdictPass BwsVerdictEnum = "PASS"

	/*
	 * Reserved, no longer in use.
	 */
	BwsVerdictPartial BwsVerdictEnum = "PARTIAL"

	/*
	 * Equivalent to "Error" or "Invalid" for the page or item in Search Console
	 */
	BwsVerdictFail BwsVerdictEnum = "FAIL"

	/*
	 * Equivalent to "Excluded" for the page or item in Search Console
	 */
	BwsVerdictNeutral BwsVerdictEnum = "NEUTRAL"
)

// Convert a string to enum
func NewBwsVerdictEnum(keyQueueName string) (BwsVerdictEnum, error) {
	return bw_enum_helper.ToEnum[BwsVerdictEnum](keyQueueName,
		BwsVerdictPass, BwsVerdictPartial, BwsVerdictFail, BwsVerdictNeutral)
}
