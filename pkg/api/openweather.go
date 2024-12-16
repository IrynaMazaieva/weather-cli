package api

import (
	"fmt"
	"strings"

	"github.com/gavincarr/countries"
)

type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	}
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  float64 `json:"humidity"`
	}
}

func (r WeatherResponse) String() string {
	info := []string{}
	info = append(info, fmt.Sprintf("Current Temp: %.f°C, Feels Like: %.f°C", r.Main.Temp, r.Main.FeelsLike))
	info = append(info, fmt.Sprintf("The high today will be %.f°C and the low today will be %.f°C", r.Main.TempMax, r.Main.TempMin))
	info = append(info, fmt.Sprintf("The humidity will be %.f%%", r.Main.Humidity))

	if len(r.Weather) != 0 {
		info = append(info, fmt.Sprintf("The weather can be described as \"%s\"", strings.TrimSpace(r.Weather[0].Description)))
	}
	return strings.Join(info, "\n")
}

type OpenWeatherClient struct {
	apiKey string
}

func NewOpenWeatherClient(apiKey string) *OpenWeatherClient {
	return &OpenWeatherClient{apiKey: apiKey}
}

func (c *OpenWeatherClient) Name() string {
	return "OpenWeather"
}

func (c *OpenWeatherClient) FetchWeather(loc Location) (string, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", loc.Lat, loc.Lon, c.apiKey)
	response, err := makeGETRequest(url, WeatherResponse{})
	if err != nil {
		return "", err
	}
	return response.String(), nil
}

func (c *OpenWeatherClient) GetLonLat(countryName string, city string) (Location, error) {
	country := countries.ByName(countryName)
	if country == countries.Unknown {
		return Location{}, fmt.Errorf("country %s not found", country)
	}
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s,%s&appid=%s", city, country.Alpha3(), c.apiKey)
	resp, err := makeGETRequest(url, []Location{})
	if err != nil {
		return Location{}, err
	}
	if len(resp) == 0 {
		return Location{}, fmt.Errorf("no location found")
	}
	return resp[0], nil
}
