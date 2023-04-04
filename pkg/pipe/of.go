package pipe

import "context"

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
