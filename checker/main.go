package main

import (
	"fmt"
	"net/http"
	"time"
)

type SiteConfig struct {
	URL           string
	AcceptedCodes []int
	Frequency     int
}

type Result struct {
	URL    string
	Up     bool
	Status int
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type DefaultClient struct{}

func (c *DefaultClient) Get(url string) (resp *http.Response, err error) {

	return http.Get(url)
}

func check(cfg SiteConfig, client HttpClient, results chan<- Result) {
	resp, err := client.Get(cfg.URL)
	result := Result{
		URL: cfg.URL,
	}

	if err != nil {
		result.Up = false
		results <- result
		return
	}

	defer resp.Body.Close()
	result.Up = false
	result.Status = resp.StatusCode
	for _, code := range cfg.AcceptedCodes {
		if resp.StatusCode == code {
			result.Up = true
			result.Status = resp.StatusCode
			break
		}
	}

	results <- result
}

func scheduleCheck(cfg SiteConfig, client HttpClient, results chan<- Result) {
	go func() {
		timer := time.NewTicker(time.Duration(cfg.Frequency) * time.Second)

		for {
			select {
			case <-timer.C:
				check(cfg, client, results)
			}
		}

	}()
}

func main() {

	sites := []SiteConfig{
		{
			URL:           "https://google.com",
			Frequency:     1,
			AcceptedCodes: []int{200},
		}, {
			URL:           "http://localhost.cs",
			Frequency:     3,
			AcceptedCodes: []int{200},
		},
		{
			URL:           "https://go.dev",
			Frequency:     1,
			AcceptedCodes: []int{200},
		},
		{
			URL:           "https://innorik.com",
			Frequency:     1,
			AcceptedCodes: []int{200},
		},
	}

	client := &DefaultClient{}

	results := make(chan Result)

	for _, site := range sites {
		scheduleCheck(site, client, results)
	}

	for result := range results {
		if result.Up {
			fmt.Printf("%s is up with status code of  %d\n", result.URL, result.Status)
		} else {
			fmt.Printf("%s is down with status code of  %d\n", result.URL, result.Status)

		}
	}
}
