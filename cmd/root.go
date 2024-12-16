package cmd

import (
	"fmt"
	"weather-cli/pkg/api"
	"weather-cli/pkg/config"

	"github.com/spf13/cobra"
)

var WeatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Fetch weather for a location for today",
	Long: "CLI application in Go that accepts a country and a city as input\n" +
		"and returns the weather information for the current day. The output\n" +
		"will display the results from the API that responded the fastest.",
	Example: "weather --country=US --city=New York",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		country := args[0]
		city := args[1]
		res, err := Execute(country, city)
		if err != nil {
			return err
		}
		fmt.Printf("Weather for %s (%s):\n%s\n", city, country, res)
		return nil
	},
}

func Execute(country string, city string) (string, error) {
	conf := config.LoadConfig()

	openWeatherClient := api.NewOpenWeatherClient(conf.OpenWeatherApiKey)
	weatherapiClient := api.NewWeatherapiClient(conf.WeatherapiApiKey)
	loc, err := openWeatherClient.GetLonLat(country, city)
	if err != nil {
		return "", err
	}
	return api.GetFastestResponse([]api.WeatherClient{openWeatherClient, weatherapiClient}, loc)
}
