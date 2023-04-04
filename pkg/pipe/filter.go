package pipe

func Filter[In any](f func(In) bool, in <-chan In) <-chan In {
	out := make(chan In)
	go func() {
		defer close(out)

		for v := range in {
			if f(v) {
				out <- v
			}
		}
	}()

	return out
}
