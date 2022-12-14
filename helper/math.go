package helper

import "math"

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func Max(x int, y int) int {
	if x < y { return y }
	return x
}

func Min(x int, y int) int {
	if x < y { return x }
	return y
}

func MinArray(x []int) int{
	return Fold(x, func(x int, acc int) int {
		return Min(x,acc)
	})
}

func MaxArray(x []int) int{
	return Fold(x, func(x int, acc int) int {
		return Max(x,acc)
	})
}
