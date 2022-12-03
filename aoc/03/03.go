package _3

import (
	"AoC2022/helper"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

func runeToNumber(c rune) int {
	if (97 <= c) && (c <= 122) { // a-z
		return int(c) - 97 + 1
	} else if (65 <= c) && (c <= 90) { // A-Z
		return int(c) - 65 + 27
	}
	return 255
}
func t01() {
	lines := helper.Split(helper.ReadLines("aoc/03/03"), "\n")
	lines2 := helper.Collect(lines, func(s string) int {
		midpoint := len(s) / 2
		lower := s[:midpoint]
		upper := s[midpoint:]

		f := func(c rune, set mapset.Set[int]) mapset.Set[int] {
			set.Add(runeToNumber(c))
			return set
		}
		lower_set := helper.FoldI([]rune(lower), mapset.NewSet[int](), f)
		upper_set := helper.FoldI([]rune(upper), mapset.NewSet[int](), f)
		intersection := lower_set.Intersect(upper_set)
		v, _ := intersection.Pop()
		return v
	})
	fmt.Println(helper.Sum(lines2))

}

func t02() {
	lines := helper.Split(helper.ReadLines("aoc/03/03"), "\n")
	out := 0
	for i := 0; i < int(len(lines)/3); i++ {
		group := lines[i*3 : (i+1)*3]
		f := func(c rune, set mapset.Set[int]) mapset.Set[int] {
			set.Add(runeToNumber(c))
			return set
		}
		set1 := helper.FoldI([]rune(group[0]), mapset.NewSet[int](), f)
		set2 := helper.FoldI([]rune(group[1]), mapset.NewSet[int](), f)
		set3 := helper.FoldI([]rune(group[2]), mapset.NewSet[int](), f)
		intersection := set1.Intersect(set2).Intersect(set3)
		v, _ := intersection.Pop()
		out += v
	}
	fmt.Print(out)
}

func Run() {
	t01()
	t02()
}
