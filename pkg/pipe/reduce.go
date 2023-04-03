package pipe

type ReduceFunc[In, Out any] func(carry Out, item In) Out

func Reduce[In, Out any](fn ReduceFunc[In, Out], init Out, in <-chan In) <-chan Out {
	out := make(chan Out)

	go func() {
		defer close(out)

		acc := init
		for x := range in {
			acc = fn(acc, x)
		}
		out <- acc
	}()

	return out
}
