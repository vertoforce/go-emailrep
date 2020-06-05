package emailrep

import (
	"context"
	"fmt"
)

func Example() {
	c := &Client{}
	res, err := c.Query(context.Background(), "test@test.com", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Email)

	// Output: test@test.com
}
