package entity

import (
	"encoding/json"
	"sort"
)

type BwsKeywordsPerformance []BwsKeywordPerformance

// Check order of elements
func (t BwsKeywordsPerformance) Less(i, j int) bool {
	return t[i].Position < (t[j].Position)
}

// size of elements, interface sort.Interface to invites
func (t BwsKeywordsPerformance) Len() int {
	return len(t)
}

// change elements of position, interface sort.Interface to alter position of invites
func (t BwsKeywordsPerformance) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type BwsKeywordPerformance struct {
	Keyword     string  `json:"keyword"`
	Clicks      float64 `json:"clicks"`
	Impressions float64 `json:"impressions"`
	Ctr         float64 `json:"ctr"`
	Position    float64 `json:"position"`
}

/**
 * Marshal indicationD1 with methods
 */
func (t *BwsKeywordsPerformance) MarshalJSON() ([]byte, error) {
	type aliasType BwsKeywordsPerformance
	sort.Sort(t)
	return json.Marshal(aliasType(*t))
}
