package main

import (
	"fmt"
	"time"

	"github.com/mpragliola/pipe/pkg/pipe"
)

func main() {
	start := pipe.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	work := func(i int) int {
		s := time.Duration(i*500) * time.Millisecond
		fmt.Println("... process item", i, "for", s, "seconds")
		time.Sleep(s)

		return i
	}
	parallel := pipe.Parallel(4, work, start)

	for r := range parallel {
		fmt.Println("->", r)
	}
}
