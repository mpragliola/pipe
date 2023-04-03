package pipe

func Of[In any](items ...In) <-chan In {
	out := make(chan In)

	go func() {
		defer close(out)

		for _, item := range items {
			out <- item
		}
	}()

	return out
}
