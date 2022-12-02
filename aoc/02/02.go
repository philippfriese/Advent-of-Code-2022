package _2

import (
	"AoC2022/helper"
	"fmt"
)

type RPS uint8

const (
	Rock     RPS = 1
	Paper    RPS = 2
	Scissors RPS = 3
)

type Outcome uint8

const (
	Win  Outcome = 6
	Draw Outcome = 3
	Loss Outcome = 0
)

func assign(s string) RPS {
	switch s {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors

	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors

	default:
		return Rock
	}
}

func assignOutcome(s string) Outcome {
	switch s {
	case "X":
		return Loss
	case "Y":
		return Draw
	case "Z":
		return Win

	default:
		return Loss
	}
}

func beats(l RPS, r RPS) bool {
	if (l == Rock && r == Scissors) ||
		(l == Scissors && r == Paper) ||
		(l == Paper && r == Rock) {
		return true
	}
	return false
}

func solve(l RPS, r RPS) Outcome {
	if beats(l, r) {
		return Win
	} else if beats(r, l) {
		return Loss
	}
	return Draw
}

func t01() {
	lines := helper.Split(helper.ReadLines("aoc/02/test"), "\n")
	lines2 := helper.Collect(lines, func(s string) int {
		inner := helper.Split(s, " ")
		other := assign(inner[0])
		mine := assign(inner[1])
		return int(solve(mine, other)) + int(mine)
	})
	fmt.Println(helper.Sum(lines2))
}

func t02() {
	type Tuple struct {
		l RPS
		r Outcome
	}

	lines := helper.Split(helper.ReadLines("aoc/02/021"), "\n")

	choices := []RPS{Rock, Paper, Scissors}
	mapping := make(map[Tuple]RPS)
	for _, l := range choices {
		for _, r := range choices {
			outcome := solve(l, r)
			mapping[Tuple{r, outcome}] = l
		}
	}

	lines2 := helper.Collect(lines, func(s string) int {
		inner := helper.Split(s, " ")
		other := assign(inner[0])
		outcome := assignOutcome(inner[1])
		mine := mapping[Tuple{other, outcome}]
		return int(outcome) + int(mine)
	})
	fmt.Println(helper.Sum(lines2))
}

func Run() {
	t01()
	t02()
}
