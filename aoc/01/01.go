package _1

import (
	"AoC2022/helper"
	"fmt"
	"sort"
)

func t01() {
	lines := helper.Split(helper.ReadLines("aoc/01/011"), "\n\n")
	lines2 := helper.Collect(lines, func(s string) []int {
		return helper.SplitConvert(s, "\n", helper.ConvInt)
	})
	sum := helper.Collect(lines2, helper.Sum[int])

	var max int = 0
	var idxmax int = 0
	for i, b := range sum {
		if max < b {
			max = b
			idxmax = i
		}
	}
	fmt.Println(idxmax)
	fmt.Println(max)
}

func t02() {
	lines := helper.Split(helper.ReadLines("aoc/01/011"), "\n\n")
	lines2 := helper.Collect(lines, func(s string) []int {
		return helper.SplitConvert(s, "\n", helper.ConvInt)
	})
	sum := helper.Collect(lines2, helper.Sum[int])

	sort.Sort(sort.Reverse(sort.IntSlice(sum)))

	fmt.Println(helper.Sum(sum[:3]))
}

func AoC01() {
	t01()
	t02()
}
