package api

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationAreaResp := LocationAreaResp{}
		if err := json.Unmarshal(data, &locationAreaResp); err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil	
	}
		fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, nil
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, nil
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode) 
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	if err := json.Unmarshal(data, &locationAreaResp); err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil	
}
