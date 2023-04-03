package pipe

// Chunk receives from a channel, groups the values in chunks of size `size`
// and emits the chunk in a channel of slices of the original datatype
func Chunk[In any](size int, in <-chan In) <-chan []In {
	out := make(chan []In)

	go func() {
		defer close(out)

		chunk := make([]In, 0, size)

		for v := range in {
			chunk = append(chunk, v)
			if len(chunk) == size {
				out <- chunk
				chunk = make([]In, 0, size)
			}
		}

		// flush remaining
		if len(chunk) > 0 {
			out <- chunk
		}
	}()

	return out
}
