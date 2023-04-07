package pipe

type PipeFunc[In, Out any] func(In) Out

func Pipe[In, Out any](fn PipeFunc[In, Out], in <-chan In) <-chan Out {
	out := make(chan Out)

	go func() {
		defer close(out)
		for in := range in {
			out <- fn(in)
		}
	}()

	return out
}

// P is an alias for `pipe.Pipe()`
func P[In, Out any](fn func(In) Out, in <-chan In) <-chan Out {
	return Pipe(fn, in)
}
