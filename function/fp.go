package function

func Const[T any](v T) func(any) T {
	return func(any) T {
		return v
	}
}
