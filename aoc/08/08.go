package _8

import (
	"AoC2022/helper"
	"fmt"
)

func walk(forest *[][]int, current_start *int, i int, current_max *int, visible_map [][]bool,
	direction int, max int, row bool) {
	var _i int
	var _j int
	for j := *current_start + direction; j*direction < max; j = j + direction {
		if row {
			_i = i
			_j = j
		} else {
			_i = j
			_j = i
		}
		if (*forest)[_i][_j] > *current_max {
			visible_map[_i][_j] = true
			*current_max = (*forest)[_i][_j]
			*current_start = j
			return
		}
	}
	*current_start = max
}

func t01(forest *[][]int) {
	/*
		idea
		from each of the four corners, start with the first entry I
		then walk as far towards the opposite side until you find an entry J s.t. I < J
		J is visible, the rest remains initialised to not visible
		set new start position of index(J) and start walking from there, set I to J
		repeat until walked out of row/col

		advantage: skip ahead many elements in the "set I to J" step
		disadvantage: mess to implement
	*/
	visible_map := make([][]bool, len(*forest))
	for i, r := range *forest {
		visible_map[i] = make([]bool, len(r))

	}
	for i := 0; i < len(visible_map); i++ {
		for j := 0; j < len(visible_map); j++ {
			visible_map[0][j] = true
			visible_map[len(visible_map)-1][j] = true
		}
		visible_map[i][0] = true
		visible_map[i][len(visible_map)-1] = true
	}

	for i := 0; i < len(*forest); i++ {
		// forward
		current_start := 0
		current_max := (*forest)[i][current_start]
		for current_start < len((*forest)[i]) {
			walk(forest, &current_start, i, &current_max, visible_map, +1, len((*forest)[i]), true)
		}

		// backward
		current_start = len((*forest)[i]) - 1
		current_max = (*forest)[i][current_start]
		for current_start > 0 {
			walk(forest, &current_start, i, &current_max, visible_map, -1, 0, true)
		}
	}

	for j := 0; j < len(*forest); j++ {
		// forward
		current_start := 0
		current_max := (*forest)[current_start][j]
		for current_start < len((*forest)[j]) {
			walk(forest, &current_start, j, &current_max, visible_map, +1, len((*forest)[j]), false)
		}

		// backward
		current_start = len((*forest)[j]) - 1
		current_max = (*forest)[current_start][j]
		for current_start > 0 {
			walk(forest, &current_start, j, &current_max, visible_map, -1, 0, false)
		}
	}

	sum := helper.FoldI(visible_map, 0, func(l []bool, i int) int {
		return i + helper.FoldI(l, 0, func(b bool, j int) int {
			if b {
				return j + 1
			} else {
				return j
			}
		}) // kennen Sie: int(bool), Mr./Ms. Go?
	})
	fmt.Println(sum)
}

func view(forest *[][]int, i int, j int, direction int, max int, row bool) int {
	var _i int
	var _j int

	if row {
		_i = i
		_j = j
	} else {
		_i = j
		_j = i
	}

	steps := 0
	var __i int
	var __j int
	for k := _j + direction; k*direction < max; k = k + direction {
		if row {
			__i = _i
			__j = k
		} else {
			__i = k
			__j = _i
		}
		steps += 1
		if (*forest)[__i][__j] >= (*forest)[i][j] {
			return steps
		}
	}
	return steps

}

func t02(forest *[][]int) {
	// welcome to off-by-one hell
	current_max := 0
	for i := 1; i < len(*forest)-1; i++ {
		for j := 1; j < len((*forest)[i])-1; j++ {
			right := view(forest, i, j, +1, len(*forest), true)
			left := view(forest, i, j, -1, 1, true)
			top := view(forest, i, j, +1, len(*forest), false)
			bottom := view(forest, i, j, -1, 1, false)

			score := right * left * bottom * top
			//fmt.Printf("[%d,%d] %d: right: %d, left: %d, top: %d, bottom: %d, score: %d\n", i, j, (*forest)[i][j],
			//	right, left, top, bottom, score)
			if current_max < score {
				current_max = score
			}

		}

	}
	fmt.Println(current_max)
}
func Run() {
	lines := helper.Split(helper.ReadLines("aoc/08/08.inp"), "\n")
	forest := make([][]int, len(lines))
	for i, r := range lines {
		forest[i] = make([]int, len(r))
		for j, c := range r {
			height := helper.ConvInt(string(c))
			forest[i][j] = height
		}
	}
	t01(&forest)
	t02(&forest)
}
