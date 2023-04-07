package pipe

// ReduceFunc is a typical reduction function, where some operation is performed
// on the currenty carry result and the item of the stream and the result is stored
// into the carry again.
type ReduceFunc[In, Out any] func(carry Out, item In) Out

// Reduce will apply a typical reduce function and progressively emit the result
// of the reduction; `init` is the initial carry value.
func Reduce[In, Out any](fn ReduceFunc[In, Out], init Out, in <-chan In) <-chan Out {
	out := make(chan Out)

	go func() {
		defer close(out)

		acc := init
		for x := range in {
			acc = fn(acc, x)
			out <- acc
		}
	}()

	return out
}
