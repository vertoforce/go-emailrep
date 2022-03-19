package emailrep

import "net/http"

// Client is an emailrep client that stores an api key
type Client struct {
	// It is not required to use an API key but emailrep.io will ask for one if you make too many requests
	APIKey string
	// Optional user agent.  Emailrepo.io asks for each app to set their user agent to the name of your app
	// If you leave it blank, the default go user agent will be used
	UserAgent  string
	HTTPClient *http.Client
}

func NewClient(APIKey string) *Client {
	return &Client{APIKey: APIKey, HTTPClient: http.DefaultClient}
}

func (c *Client) request(r *http.Request) (*http.Response, error) {
	if c.UserAgent != "" {
		r.Header.Add("User-Agent", c.UserAgent)
	}
	if c.APIKey != "" {
		r.Header.Add("Key", c.APIKey)
	}

	return c.HTTPClient.Do(r)
}
