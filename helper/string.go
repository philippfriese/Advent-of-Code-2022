package helper

import (
	"golang.org/x/exp/constraints"
	"strings"
)

func Split[T constraints.Ordered](fn string, separator string, f func(string) T) []T {
	lines := strings.Split(
		ReadLines(fn),
		separator)
	return Collect(lines, f)
}

func Collect[T any, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}
