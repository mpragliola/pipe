package pipe

// Spread flattens a channel of slices of `[]T` into a channel of `T` type
func Spread[In any](in <-chan []In) <-chan In {
	out := make(chan In)

	go func() {
		defer close(out)
		for v := range in {
			for _, i := range v {
				out <- i
			}
		}
	}()

	return out
}
