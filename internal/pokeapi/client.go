package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/CodeZeroSugar/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func (c *Client) Fetch(url string) ([]byte, error) {
	if b, ok := c.cache.Get(url); ok {
		return b, nil
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, body)
	return body, nil
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
