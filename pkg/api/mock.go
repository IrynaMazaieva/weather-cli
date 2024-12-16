package api

import "time"

type MockWeatherClient struct {
	MockName string
	Delay    time.Duration
	Response string
}

func NewMockWeatherClient(name string, delay time.Duration, response string) *MockWeatherClient {
	return &MockWeatherClient{
		MockName: name,
		Delay:    delay,
		Response: response,
	}
}

func (c *MockWeatherClient) Name() string {
	return c.MockName
}

func (c *MockWeatherClient) FetchWeather(loc Location) (string, error) {
	time.Sleep(c.Delay)
	return c.Response, nil
}
