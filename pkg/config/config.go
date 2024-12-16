package config

import "os"

const (
	OPEN_WEATHER_API_KEY = "OPEN_WEATHER_API_KEY"
	WEATHERAPI_API_KEY   = "WEATHERAPI_API_KEY"
)

type Config struct {
	OpenWeatherApiKey string
	WeatherapiApiKey  string
}

func LoadConfig() Config {
	return Config{
		OpenWeatherApiKey: os.Getenv(OPEN_WEATHER_API_KEY),
		WeatherapiApiKey:  os.Getenv(WEATHERAPI_API_KEY),
	}
}
