package emailrep

import (
	"context"
	"testing"
)

func TestQuery(t *testing.T) {
	c := &Client{}
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
