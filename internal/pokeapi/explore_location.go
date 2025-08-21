package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(area string) (RespLocationInfo, error) {
	url := baseURL + "/location-area"
	if area != "" {
		url = url + "/" + area
	}

	dat, exists := c.cache.Get(url)

	// Make API request if cache does not exist
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationInfo{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationInfo{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocationInfo{}, err
		}
	}

	locationInfo := RespLocationInfo{}
	err := json.Unmarshal(dat, &locationInfo)
	if err != nil {
		return RespLocationInfo{}, err
	}

	return locationInfo, nil
}
