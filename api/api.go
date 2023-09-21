package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"gabrieleromanato/unsplash/utils"
	"io"
	"net/http"
	"net/url"
	"os"
)

func SearchImages(query string) ([]string, error) {
	var results map[string]interface{}
	var urls []string

	baseURL := "https://api.unsplash.com"
	resource := "/search/photos"

	params := url.Values{}
	params.Add("query", query)
	params.Add("client_id", os.Getenv("APIKEY"))

	uri, _ := url.ParseRequestURI(baseURL)
	uri.Path = resource
	uri.RawQuery = params.Encode()

	urlStr := fmt.Sprintf("%v", uri)

	resp, err := http.Get(urlStr)
	if err != nil {
		return urls, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return urls, errors.New("Request error")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return urls, err
	}

	jsonData := string(body)

	jsonErr := json.Unmarshal([]byte(jsonData), &results)

	if jsonErr != nil {
		return urls, jsonErr
	}

	urls, err = utils.GetURLsFromResponse(results)

	if err != nil {
		return urls, err
	}

	return urls, nil
}
