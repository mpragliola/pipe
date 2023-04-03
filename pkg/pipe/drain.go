package pipe

func Drain[In any](in <-chan In) <-chan []In {
	appender := func(s []In, i In) []In {
		return append(s, i)
	}

	return Reduce(appender, make([]In, 0), in)
}
