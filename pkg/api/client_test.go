package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetFastestResponse_significantDifference(t *testing.T) {
	slowMockClient := NewMockWeatherClient("Slow", 10*time.Second, "Slow response")
	fastMockClient := NewMockWeatherClient("Fast", 1*time.Second, "Fast response")

	fastest, err := GetFastestResponse([]WeatherClient{slowMockClient, fastMockClient}, Location{})
	assert.NoError(t, err)
	assert.Equal(t, "Fast response", fastest)
}

func Test_GetFastestResponse_sameSpeed(t *testing.T) {
	firstMockClient := NewMockWeatherClient("First", 1*time.Second, "First response")
	secondMockClient := NewMockWeatherClient("Second", 1*time.Second, "Second response")

	fastest, err := GetFastestResponse([]WeatherClient{firstMockClient, secondMockClient}, Location{})
	assert.NoError(t, err)
	assert.Equal(t, "First response", fastest)
}

func Test_GetFastestResponse_bothTooSlow(t *testing.T) {
	slowMockClient := NewMockWeatherClient("Slow", 10*time.Second, "Slow response")
	slowerMockClient := NewMockWeatherClient("Slower", 10*time.Second, "Slower response")

	_, err := GetFastestResponse([]WeatherClient{slowMockClient, slowerMockClient}, Location{})
	assert.Error(t, err)
}
