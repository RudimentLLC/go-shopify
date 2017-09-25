package shopify

import (
	"fmt"

	"github.com/zpatrick/rclient"
)

type Client struct {
	Client *rclient.RestClient
}

func NewClient(domain, key, pass string, options ...rclient.ClientOption) *Client {
	baseURL := fmt.Sprintf("https://%s:%s@%s", key, pass, domain)

	return &Client{
		Client: rclient.NewRestClient(baseURL, options...),
	}
}
