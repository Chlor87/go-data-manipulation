package function

func Curry2[A, B, C any](fn func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return fn(a, b)
		}
	}
}

func Curry3[A, B, C, D any](fn func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return Curry2(func(b B, c C) D {
			return fn(a, b, c)
		})
	}
}
