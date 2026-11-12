package aur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func request(endpoint string) ([]Package, error) {
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var data Response

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Results, nil
}

func Search(name string) ([]Package, error) {
	endpoint := "https://aur.archlinux.org/rpc/v5/search/" +
		url.PathEscape(name)

	return request(endpoint)
}

func Info(name string) ([]Package, error) {
	endpoint := "https://aur.archlinux.org/rpc/v5/info?arg[]=" +
		url.QueryEscape(name)

	return request(endpoint)
}
