package helper

import "golang.org/x/exp/constraints"

func Collect[T any, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}

func Apply[T any](list []T, f func(T)) {
	for _, v := range list {
		f(v)
	}
}

func Fold[T any](list []T, f func(T, T) T) T {
	return FoldI(list[1:], list[0], f)
}

func FoldI[T any, U any](list []T, init U, f func(T, U) U) U {
	acc := init
	for _, v := range list {
		acc = f(v, acc)
	}
	return acc
}

func Filter[T any](list []T, f func(T) bool) []T {
	output := make([]T, 0)
	for _, t := range list {
		if f(t) {
			output = append(output, t)
		}
	}
	return output
}

func Sum[T constraints.Ordered](list []T) T {
	return Fold(list, func(a T, b T) T { return a + b })
}
