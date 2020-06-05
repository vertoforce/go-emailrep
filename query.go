package emailrep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://emailrep.io"
)

// Query Makes a request to emailrep.io to get information on the email reputation
// returnSummary is if you want the request to return a human readable summary
func Query(ctx context.Context, email string, returnSummary bool) (*SearchResponse, error) {
	// Build URL
	URL := fmt.Sprintf("%s/%s", baseURL, email)
	if returnSummary {
		URL += "?summary=true"
	}

	// Prepare and make request
	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 400:
		return nil, fmt.Errorf("invalid email")
	case 401:
		return nil, fmt.Errorf("invalid api key")
	case 429:
		return nil, fmt.Errorf("too many requests. contact emailrep.io for an api key")
	}

	ret := &SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
