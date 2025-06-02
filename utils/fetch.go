package utils

import (
	"net/http"
	"time"
)

// Fetches response from a url.
func Fetch(url string) (*http.Response, int, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, 0, err
	}

	return resp, resp.StatusCode, nil
}
