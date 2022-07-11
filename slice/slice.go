package slice

func Reduce[P, C any](fn func(P, C) P, acc P, input []C) P {
	for _, c := range input {
		acc = fn(acc, c)
	}
	return acc
}

func Map[I, O any](fn func(I) O, input []I) []O {
	return Reduce(func(p []O, c I) []O {
		return append(p, fn(c))
	}, []O{}, input)
}

func Filter[T any](fn func(T) bool, input []T) []T {
	return Reduce(func(p []T, c T) []T {
		if fn(c) {
			return append(p, c)
		}
		return p
	}, []T{}, input)
}
