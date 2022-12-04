package _4

import (
	"AoC2022/helper"
	"fmt"
)

type Range struct {
	start uint8
	end   uint8
}

func createRange(s string) (Range, error) {
	vars := helper.SplitConvert(s, "-", helper.ConvInt)
	start := uint8(vars[0])
	end := uint8(vars[1])
	if start > end {
		return Range{}, fmt.Errorf("Start cannot be after end! start: %s, end: %s", start, end)
	}
	return Range{start, end}, nil
}

func (r Range) intersects(other Range) bool {
	return ((other.start <= r.start) && (r.start <= other.end)) ||
		((other.start <= r.end) && (r.end <= other.end)) ||
		((r.start <= other.start) && (other.start <= r.end)) ||
		((r.start <= other.end) && (other.end <= r.end))
}

func (r Range) contains(other Range) bool {
	return (r.start <= other.start) && (other.end <= r.end)
}

func t01() {

	lines := helper.Collect(
		helper.Split(helper.ReadLines("aoc/04/04"), "\n"),
		func(s string) []Range {
			return helper.SplitConvert(s, ",", func(s string) Range {
				r, err := createRange(s)
				if err != nil {
					fmt.Printf("Error on range creation: %s\n", err)
				}
				return r
			})
		})

	contains := helper.Collect(lines, func(l []Range) int {
		left := l[0]
		right := l[1]
		if left.contains(right) || right.contains(left) {
			return 1
		} // hear my eyes roll
		return 0
	})
	fmt.Println(helper.Sum(contains))

}

func t02() {
	lines := helper.Collect(
		helper.Split(helper.ReadLines("aoc/04/04"), "\n"),
		func(s string) []Range {
			return helper.SplitConvert(s, ",", func(s string) Range {
				r, err := createRange(s)
				if err != nil {
					fmt.Printf("Error on range creation: %s\n", err)
				}
				return r
			})
		})

	contains := helper.Collect(lines, func(l []Range) int {
		left := l[0]
		right := l[1]
		if left.intersects(right) || right.intersects(left) {
			return 1
		} // hear my eyes roll
		return 0
	})
	fmt.Println(helper.Sum(contains))
}

func Run() {
	t01()
	t02()
}
