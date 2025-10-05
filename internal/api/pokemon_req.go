package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, nil
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)	
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}	
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}
	
	return  pokemon, nil
}
