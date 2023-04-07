package pipe

// Scoop will drain a channel and return the values in a slice.
func Scoop[In any](in <-chan In) []In {
	out := make([]In, 0)

	for r := range in {
		out = append(out, r)
	}

	return out
}
