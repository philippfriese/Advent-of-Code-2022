package _5

import (
	"AoC2022/helper"
	"fmt"
	"strconv"
)

type MultiStack struct {
	stack [][]rune
}

func createMultiStack(size uint8) MultiStack {
	var arr [][]rune
	for i := uint8(0); i < size; i++ {
		arr = append(arr, []rune{})
	}
	return MultiStack{arr}
}

func fillMultiStack(ms MultiStack, lines []string) MultiStack {
	size := len(lines)
	for i := size - 2; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < len(ms.stack); j++ {
			r := rune(line[1+j*4])
			if r != ' ' {
				ms.push(uint8(j), r)
			}
		}

	}
	return ms
}

func (ms MultiStack) pop(i uint8) rune {
	stack := ms.stack[i]
	l := len(stack)
	s, e := stack[:l-1], stack[l-1]
	ms.stack[i] = s
	return e
}

func (ms MultiStack) push(i uint8, elem rune) {
	ms.stack[i] = append(ms.stack[i], elem)
}

func (ms MultiStack) print() {
	for i := 0; i < len(ms.stack); i++ {
		fmt.Printf("%d: ", i+1)
		for j := 0; j < len(ms.stack[i]); j++ {
			fmt.Printf("[%c]", ms.stack[i][j])
		}
		fmt.Println()
	}
}

func (ms MultiStack) move(amount uint8, from uint8, to uint8) {
	for i := uint8(0); i < amount; i++ {
		e := ms.pop(from - 1)
		ms.push(to-1, e)
	}
}

func (ms MultiStack) move9001(amount uint8, from uint8, to uint8) {
	s := ms.stack[from-1]
	l := uint8(len(s))
	moved := s[l-amount:]
	ms.stack[from-1] = s[:l-amount]
	ms.stack[to-1] = append(ms.stack[to-1], moved...)
}

func parseMove(move string) (uint8, uint8, uint8) {
	parts := helper.Split(move, " ")
	// there are no errors in my world!
	amount, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])
	return uint8(amount), uint8(from), uint8(to)
}

func t01() {
	lines := helper.Split(helper.ReadLines("aoc/05/05.inp"), "\n\n")
	stack := lines[0]
	moves := helper.Split(lines[1], "\n")

	stack_lines := helper.Split(stack, "\n")
	num_stacks := (len(stack_lines[0]) + 1) / 4
	ms := createMultiStack(uint8(num_stacks))
	fillMultiStack(ms, stack_lines)
	for _, s := range moves {
		ms.move(parseMove(s))
	}
	fmt.Println(helper.FoldI(ms.stack, "", func(stack []rune, acc string) string {
		return acc + string(stack[len(stack)-1])
	}))
}

func t02() {
	lines := helper.Split(helper.ReadLines("aoc/05/05.inp"), "\n\n")
	stack := lines[0]
	moves := helper.Split(lines[1], "\n")

	stack_lines := helper.Split(stack, "\n")
	num_stacks := (len(stack_lines[0]) + 1) / 4
	ms := createMultiStack(uint8(num_stacks))
	fillMultiStack(ms, stack_lines)
	for _, s := range moves {
		ms.move9001(parseMove(s))
	}
	fmt.Println(helper.FoldI(ms.stack, "", func(stack []rune, acc string) string {
		return acc + string(stack[len(stack)-1])
	}))
}
func Run() {
	t01()
	t02()
}
