package api

import (
	"fmt"
	"strings"
)

const forecastDaysCount = 1

type ForecastResponse struct {
	Current struct {
		Temp      float64 `json:"temp_c"`
		FeelsLike float64 `json:"feelslike_c"`
	}
	Forecast struct {
		ForecastDay []struct {
			Day struct {
				TempMax   float64 `json:"maxtemp_c"`
				TempMin   float64 `json:"mintemp_c"`
				Humidity  int     `json:"avghumidity"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func (f ForecastResponse) String() string {
	info := fmt.Sprintf("Current Temp: %.f°C, Feels Like: %.f°C", f.Current.Temp, f.Current.FeelsLike)
	if len(f.Forecast.ForecastDay) == 0 {
		return info
	}
	forecast := []string{info}
	forecast = append(forecast, fmt.Sprintf("The high today will be %.f°C and the low today will be %.f°C", f.Forecast.ForecastDay[0].Day.TempMax, f.Forecast.ForecastDay[0].Day.TempMin))
	forecast = append(forecast, fmt.Sprintf("The humidity will be %d%%", f.Forecast.ForecastDay[0].Day.Humidity))
	forecast = append(forecast, fmt.Sprintf("The weather can be described as \"%s\"", strings.TrimSpace(f.Forecast.ForecastDay[0].Day.Condition.Text)))
	return strings.Join(forecast, "\n")
}

type WeatherapiClient struct {
	apiKey string
}

func NewWeatherapiClient(apiKey string) *WeatherapiClient {
	return &WeatherapiClient{apiKey: apiKey}
}

func (c *WeatherapiClient) Name() string {
	return "Weatherapi"
}

func (c *WeatherapiClient) FetchWeather(loc Location) (string, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%f,%f&days=%d", c.apiKey, loc.Lat, loc.Lon, forecastDaysCount)
	resp, err := makeGETRequest(url, ForecastResponse{})
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
