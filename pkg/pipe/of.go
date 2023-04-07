package pipe

import (
	"context"
)

// Of will sequentially emit the `items` of type `T` in a stream (channel) of type `T`.
// The stream will be closed when the `items` are exhausted. To pass arrays and slices
// you can use the spread (`...`) operator.
func Of[In any](items ...In) <-chan In {
	out := make(chan In)

	go func() {
		for _, item := range items {
			out <- item
		}

		close(out)
	}()

	return out
}

type OfFuncGenerator[In any] func() In

// OfFunc will generate a stream (channel) of type `T` from a function that returns
// a value of type `T`. The stream will be closed when the provided context will be
// canceled.
func OfFunc[In any](ctx context.Context, generate OfFuncGenerator[In]) <-chan In {
	out := make(chan In)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				out <- generate()
			}
		}
	}()

	return out
}
