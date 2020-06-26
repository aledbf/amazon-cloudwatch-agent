package statsd

import (
	"testing"

	"github.com/aws/amazon-cloudwatch-agent/tool/runtime"

	"github.com/stretchr/testify/assert"
)

func TestStatsD_ToMap(t *testing.T) {
	expectedKey := "statsd"
	expectedValue := map[string]interface{}{
		"service_address":              ":8125",
		"metrics_collection_interval":  10,
		"metrics_aggregation_interval": 60,
	}
	ctx := new(runtime.Context)
	conf := new(StatsD)
	conf.Enable()
	key, value := conf.ToMap(ctx)
	assert.Equal(t, expectedKey, key)
	assert.Equal(t, expectedValue, value)
}