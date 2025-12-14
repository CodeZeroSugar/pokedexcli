package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocationArea(name string) (LocationResults, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, name)

	data, err := c.Fetch(url)
	if err != nil {
		return LocationResults{}, err
	}

	var resp LocationResults
	if err := json.Unmarshal(data, &resp); err != nil {
		return LocationResults{}, err
	}
	return resp, nil
}
