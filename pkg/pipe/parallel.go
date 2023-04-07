package pipe

import (
	"sync"
)

func Parallel[In, Out any](n int, fn PipeFunc[In, Out], in <-chan In) <-chan Out {
	out := make(chan Out)

	var wg sync.WaitGroup
	wg.Add(n)

	go func() {

		for i := 0; i < n; i++ {
			go func(i int) {
				for v := range in {
					out <- fn(v)
				}
				wg.Done()
			}(i)
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
