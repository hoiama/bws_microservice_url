package a_test

import (
	"bws_microservice_url/main/src/external"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleSearchConsoleApi_GetMetrics(t1 *testing.T) {
	result, err := new(external.GoogleSearchConsoleApi).Constructor1().GetMetrics("https://partner.bindways.com")
	assert.Nil(t1, err)
	assert.NotEqual(t1, 0, len(result.GscRows))
}
