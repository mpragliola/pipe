package main

import (
	"context"
	"fmt"

	"github.com/mpragliola/pipe/pkg/pipe"
)

func main() {

	i := 0
	ctx, cancel := context.WithCancel(context.Background())

	emit := pipe.OfFunc(
		ctx,
		func() int {
			i++
			if i >= 5 {
				cancel()
			}

			return i
		},
	)

	for r := range emit {
		fmt.Println(r)
	}
}
