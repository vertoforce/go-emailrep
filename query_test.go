package emailrep

import (
	"context"
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	key := os.Getenv("EMAILREP_API_KEY")
	if key == "" {
		t.Skip("EMAILREP_API_KEY not set; skipping live API test")
	}
	c := NewClient(key)
	res, err := c.Query(context.Background(), "test@test.com", true)
	if err != nil {
		t.Error(err)
		return
	}
	if res.Email != "test@test.com" {
		t.Errorf("did not get valid response data")
	}

	// check for summary
	if res.Summary == "" {
		t.Errorf("did not get summary response")
	}
}
