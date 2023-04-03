package pipe

func Pipe[In, Out any](fn func(In) Out, in <-chan In) <-chan Out {
	out := make(chan Out)

	go func() {
		defer close(out)
		for in := range in {
			out <- fn(in)
		}
	}()

	return out
}
