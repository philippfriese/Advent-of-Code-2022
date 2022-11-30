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
