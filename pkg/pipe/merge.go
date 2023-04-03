package pipe

import "sync"

// Merge merges the input streams into a single stream
func Merge[In any](ins ...<-chan In) <-chan In {
	var wg sync.WaitGroup
	out := make(chan In)

	wg.Add(len(ins))

	go func() {
		for _, in := range ins {
			// read each channel in goroutine, otherwise we will read
			// them sequentially
			go func(in <-chan In) {
				for v := range in {
					out <- v
				}

				wg.Done()
			}(in)
		}
	}()

	// Close when done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
