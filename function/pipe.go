package function

func Pipe2[A, B, O any](f1 func(A) B, f2 func(B) O) func(A) O {
	return func(a A) O {
		return f2(f1(a))
	}
}

func Pipe3[A, B, C, D any](
	f1 func(A) B, f2 func(B) C, f3 func(C) D,
) func(A) D {
	return func(a A) D {
		return f3(Pipe2(f1, f2)(a))
	}
}

func Pipe4[A, B, C, D, E any](
	f1 func(A) B, f2 func(B) C, f3 func(C) D, f4 func(D) E,
) func(A) E {
	return func(a A) E {
		return f4(Pipe3(f1, f2, f3)(a))
	}
}

func Pipe5[A, B, C, D, E, F any](
	f1 func(A) B, f2 func(B) C, f3 func(C) D, f4 func(D) E, f5 func(E) F,
) func(A) F {
	return func(a A) F {
		return f5(Pipe4(f1, f2, f3, f4)(a))
	}
}
