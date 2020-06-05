package emailrep

import (
	"context"
	"testing"
)

func TestQuery(t *testing.T) {
	res, err := Query(context.Background(), "test@test.com")
	if err != nil {
		t.Error(err)
		return
	}
	if res.Email != "test@test.com" {
		t.Errorf("did not get valid response data")
	}
}
