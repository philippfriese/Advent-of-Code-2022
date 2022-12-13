package _13

import (
	"AoC2022/helper"
	"fmt"
	"golang.org/x/exp/slices"
)

// uses comparison trick I've seen on reddit after being frustrated by GoLangs lack of file := json.load(input)
// Go is _not_ the best language for this particular task
func compare(lhs string, rhs string) int {
	// I hate my life
	if rhs == "" {
		return 1
	}

	if lhs[0] != '[' && rhs[0] != '['{
		return helper.ConvInt(lhs) - helper.ConvInt(rhs)
	}

	if lhs[0] != '[' && rhs[0] == '[' {
		lhs = fmt.Sprintf("[%s]",lhs)
	}
	if lhs[0] == '[' && rhs[0] != '[' {
		rhs = fmt.Sprintf("[%s]",rhs)
	}
	lhs = lhs[1:len(lhs)-1]
	rhs = rhs[1:len(rhs)-1]

	if lhs == "" && rhs == "" {
		return 0
	}
	if lhs == "" {
		return -1
	}

	height_lhs := make([]int,len(lhs))
	current := 0
	for i,c := range lhs {
		if c == '[' {current += 1}
		if c == ']' {current -= 1}
		height_lhs[i] = current
	}
	height_rhs := make([]int,len(rhs))
	current = 0
	for i,c := range rhs {
		if c == '[' {current += 1}
		if c == ']' {current -= 1}
		height_rhs[i] = current
	}

	lhs_entries := []string{}
	prev := 0
	for i,v := range lhs {
		if height_lhs[i] == 0 && v == ',' {
			lhs_entries = append(lhs_entries, lhs[prev:i])
			prev = i +1
		}
	}
	lhs_entries = append(lhs_entries, lhs[prev:])

	rhs_entries := []string{}
	prev = 0
	for i,v := range rhs {
		if height_rhs[i] == 0 && v == ',' {
			rhs_entries = append(rhs_entries, rhs[prev:i])
			prev = i +1
		}
	}
	rhs_entries = append(rhs_entries, rhs[prev:])

	for i,lhs_entry := range lhs_entries {
		if i >= len(rhs_entries) {
			break
		}
		c := compare(lhs_entry, rhs_entries[i])
		if c != 0 {
			return c
		}
	}
	return len(lhs_entries) - len(rhs_entries)
}

func t01() {
	groups := helper.Split(helper.ReadLines("aoc/13/13.inp"), "\n\n")
	pairs := helper.Collect(groups, func(s string) []string {
		return helper.Split(s, "\n")
	})

	packets := helper.Collect(pairs, func(strs[]string) bool {
		return compare(strs[0], strs[1]) <= 0

	})
	sum := 0
	for i,b := range packets {
		if b { sum += i+1}
	}
	fmt.Println(sum)
}

func t02() {
	groups := helper.Split(helper.ReadLines("aoc/13/13.inp"), "\n\n")
	pairs := helper.Collect(groups, func(s string) []string {
		return helper.Split(s, "\n")
	})

	packets := make([]string,0)

	helper.Apply(pairs, func(strs[]string) {
		packets = append(packets,strs[0])
		packets = append(packets,strs[1])
	})
	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")

	slices.SortFunc(packets, func(lhs string, rhs string) bool {return compare(lhs,rhs) <= 0})

	a1 := 0
	a2 := 0
	for i,p := range packets {
		if p == "[[2]]" {a1 = i+1}
		if p == "[[6]]" {a2 = i+1}
	}
	fmt.Println(a1*a2)

}

func Run() {
	t01()
	t02()
}
