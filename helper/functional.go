package helper

import "golang.org/x/exp/constraints"

func Collect[T any, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}

func Fold[T any](list []T, f func(T,T) T) T {
	return FoldI(list[1:], list[0], f)
}

func FoldI[T any](list []T, init T, f func(T,T)T) T {
	acc := init
	for _,v := range list {
		acc = f(acc,v)
	}
	return acc
}

func Sum[T constraints.Ordered](list []T) T{
	return Fold(list, func(a T,b T)T {return a+b})
}
