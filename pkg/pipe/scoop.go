package pipe

func Scoop[In any](in <-chan In) []In {
	out := make([]In, 0)

	for r := range in {
		out = append(out, r)
	}

	return out
}
