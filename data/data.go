package data

import (
	"github.com/Chlor87/go-data-manipulation/function"
)

type (
	M = map[string]any
	A = []any
)

func fail[T any]() (T, bool) {
	var t T
	return t, false
}

func pass[T any](v any) (T, bool) {
	res, ok := v.(T)
	if !ok {
		return fail[T]()
	}
	return res, true
}

func Get[T any](key, data any) (T, bool) {
	switch key := key.(type) {
	case string:
		asM, ok := data.(M)
		if !ok {
			return fail[T]()
		}
		v, ok := asM[key]
		if !ok {
			return fail[T]()
		}
		return pass[T](v)
	case int:
		s, ok := data.(A)
		if !ok {
			return fail[T]()
		}
		if key >= len(s) {
			return fail[T]()
		}
		return pass[T](s[key])
	default:
		return fail[T]()
	}
}

func GetPath[T any](path A, data any) (T, bool) {
	switch len(path) {
	case 0:
		return fail[T]()
	case 1:
		return Get[T](path[0], data)
	}
	next, ok := Get[any](path[0], data)
	if !ok {
		return fail[T]()
	}
	return GetPath[T](path[1:], next)
}

func copySetSlice[I, O any](k int, fn func(I) O, d any) A {
	asS, ok := d.(A)
	if !ok {
		asS = nil
	}
	cp := make(A, max(len(asS), k+1))
	copy(cp, asS)
	asI, ok := cp[k].(I)
	if ok {
		cp[k] = fn(asI)
	} else {
		var i I
		cp[k] = fn(i)
	}
	return cp
}

func copySetMap[I, O any](k string, fn func(I) O, d any) M {
	asM, ok := d.(M)
	if !ok {
		asM = M{}
	}
	cp := make(M, len(asM))
	for k, v := range asM {
		cp[k] = v
	}
	asI, ok := cp[k].(I)
	if ok {
		cp[k] = fn(asI)
	} else {
		var i I
		cp[k] = fn(i)
	}
	return cp
}

// Modify could use some generic magic
func Modify[I, O any](key any, fn func(I) O, data any) any {
	switch k := key.(type) {
	case string:
		return copySetMap(k, fn, data)
	case int:
		return copySetSlice(k, fn, data)
	default:
		return nil
	}
}

func ModifyPath[I, O any](path A, fn func(I) O, data any) any {
	switch len(path) {
	case 0:
		return nil
	case 1:
		return Modify(path[0], fn, data)
	}

	d, ok := Get[any](path[0], data)
	if !ok {
		var i I
		return setStructure(path, fn(i))
	}
	return Set(path[0], ModifyPath(path[1:], fn, d), data)
}

func Set[T any](key, value T, data any) any {
	return Modify(key, function.Const(value), data)
}

func SetPath[T any](path A, value T, data any) any {
	return ModifyPath(path, function.Const(value), data)
}

// setStructure is a fast path for creating a nested structure from path with
// one leaf (value)
func setStructure(path A, value any) any {
	if len(path) == 0 {
		return value
	}
	switch k := path[0].(type) {
	case string:
		return M{k: setStructure(path[1:], value)}
	case int:
		a := make(A, k+1)
		a[k] = setStructure(path[1:], value)
		return a
	default:
		return nil
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
