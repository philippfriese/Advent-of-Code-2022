package helper

import (
	"golang.org/x/exp/constraints"
	"strings"
)

func Split(line string, separator string) []string {
	lines := strings.Split(
		line,
		separator)
	return lines
}

func SplitConvert[T constraints.Ordered](line string, separator string, f func(string) T) []T {
	return Collect(Split(line, separator), f)
}
