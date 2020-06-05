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
func Query(ctx context.Context, email string) (*SearchResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", baseURL, email), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	ret := &SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}