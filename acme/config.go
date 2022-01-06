package acme

import (
	"net/http"
	"time"
)

const DefaultHostURL string = "http://localhost:19090"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL: DefaultHostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// XXX auth with token

	return &c, nil
}

func (c *Client) CreateCollection() int {
	//  XXX request header and body
	// req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/collection", c.HostURL)))

	return -1
}

