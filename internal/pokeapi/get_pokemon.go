package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	dat, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokemon, nil
}
