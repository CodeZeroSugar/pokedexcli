package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokeInfo(name string) (PokeInfo, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)

	data, err := c.Fetch(url)
	if err != nil {
		return PokeInfo{}, err
	}

	var resp PokeInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return PokeInfo{}, err
	}
	return resp, err
}
