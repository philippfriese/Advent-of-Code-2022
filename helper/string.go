package helper

import (
	"strings"
)

func Split(line string, separator string) []string {
	lines := strings.Split(
		line,
		separator)
	return lines
}

func SplitConvert[T any](line string, separator string, f func(string) T) []T {
	return Collect(Split(line, separator), f)
}
