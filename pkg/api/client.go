package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const timeout = 5 * time.Second

type Location struct {
	City    string  `json:"name"`
	Country string  `json:"country"`
	Lon     float64 `json:"lon"`
	Lat     float64 `json:"lat"`
}

type WeatherClient interface {
	FetchWeather(Location) (string, error)
	Name() string
}

func GetFastestResponse(clients []WeatherClient, loc Location) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan string, len(clients))
	var wg sync.WaitGroup
	var errs []string

	for _, client := range clients {
		wg.Add(1)
		go func(c WeatherClient) {
			defer wg.Done()
			start := time.Now()
			data, err := c.FetchWeather(loc)
			elapsed := time.Since(start)
			if err == nil {
				fmt.Printf("%s client responded in %v\n", c.Name(), elapsed)
				select {
				case ch <- data:
				case <-ctx.Done():
				}
			} else {
				fmt.Printf("%s client failed with error: %s\n", c.Name(), err)
				errs = append(errs, fmt.Sprintf("%s: %s", c.Name(), err))
			}
		}(client)
	}

	// Close the channel when all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("timeout reached, errors: %v", errs)
	case resp := <-ch:
		return resp, nil
	}
}

func makeGETRequest[T any](url string, data T) (T, error) {
	var response T
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return response, err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
