package _9

import (
	"AoC2022/helper"
	"fmt"
	"github.com/deckarep/golang-set/v2"
	"math"
	"strconv"
)

type Move struct {
	direction string
	amount    int
}

type Knot struct {
	x int
	y int
}

func moveBulk(head *Knot, tail *[]Knot, m Move, tail_positions *mapset.Set[Knot]) {
	for i := 0; i < m.amount; i++ {
		switch m.direction {
		case "R":
			{
				head.x += 1
			}
		case "L":
			{
				head.x -= 1
			}
		case "U":
			{
				head.y += 1
			}
		case "D":
			{
				head.y -= 1
			}
		}

		(*tail)[0].moveTail(head)
		for j := 1; j < len(*tail); j++ {
			(*tail)[j].moveTail(&(*tail)[j-1])
		}
		(*tail_positions).Add((*tail)[len(*tail)-1])
	}
}

func (tail *Knot) moveTail(head *Knot) {
	if helper.Abs(tail.x-head.x) <= 1 &&
		helper.Abs(tail.y-head.y) <= 1 {
		return
	}

	dif := Knot{head.x - tail.x, head.y - tail.y}
	if dif.x == 0 {
		tail.y += helper.Sign(dif.y)
	} else if dif.y == 0 {
		tail.x += helper.Sign(dif.x)
	} else {
		tail.x += helper.Sign(dif.x)
		tail.y += helper.Sign(dif.y)
	}
}

func print_knots(head Knot, tail Knot) {
	max_x := math.Max(math.Abs(math.Max(float64(head.x), float64(tail.x))), 2)
	max_y := math.Max(math.Abs(math.Max(float64(head.y), float64(tail.y))), 2)

	for i := max_y; i >= -max_y; i-- {
		for j := -max_x; j <= max_x; j++ {
			if tail.x == int(j) && tail.y == int(i) {
				fmt.Print("T")
			} else if head.x == int(j) && head.y == int(i) {
				fmt.Print("H")
			} else if i == 0 && j == 0 {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println()
}

func print_trace(positions mapset.Set[Knot]) {
	max_x := 0
	max_y := 0
	for k := range positions.Iterator().C {
		max_x = helper.Max(k.x, max_x)
		max_y = helper.Max(k.y, max_y)
	}

	for i := max_y; i >= -max_y; i-- {
		for j := -max_x; j <= max_x; j++ {
			if positions.Contains(Knot{j, i}) {
				fmt.Print("#")
			} else if i == 0 && j == 0 {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println()
}

func t01(input *[]Move) {
	head := Knot{0, 0}
	tail := []Knot{Knot{0, 0}}
	positions := mapset.NewSet[Knot]()
	helper.Apply((*input), func(m Move) {
		moveBulk(&head, &tail, m, &positions)
	})
	fmt.Println(positions.Cardinality())
}

func t02(input *[]Move) {
	head := Knot{0, 0}
	tail := make([]Knot, 9)
	for i, _ := range tail {
		tail[i] = Knot{0, 0}
	}
	positions := mapset.NewSet[Knot]()
	helper.Apply((*input), func(m Move) {
		moveBulk(&head, &tail, m, &positions)
	})
	fmt.Println(positions.Cardinality())
}
func Run() {
	lines := helper.Split(helper.ReadLines("aoc/09/09.inp"), "\n")
	input := helper.Collect(lines, func(s string) Move {
		vs := helper.Split(s, " ")
		amt, _ := strconv.Atoi(vs[1])
		return Move{vs[0], amt}
	})
	t01(&input)
	t02(&input)
}
