package _10

import (
	"AoC2022/helper"
	"fmt"
)

type InstType uint8
const (
	noop InstType = iota
	addx
)

type Instruction struct {
	instType InstType
	value int
	cycles int
	clock int
}

func (i Instruction) String() string {
	return fmt.Sprintf("[typ:%v,val:%d,cyc:%d,clk: %d]", i.instType, i.value, i.cycles, i.clock)
}

func parseInstruction(s string) Instruction {
	parts := helper.Split(s, " ")

	switch parts[0] {
	case "noop": { return Instruction{noop, 0, 1, 0} }
	case "addx": { return Instruction{addx, helper.ConvInt(parts[1]), 2, 0} }
	}
	return Instruction{}
}

type CPU struct {
	x int
	clock int
}

func (cpu *CPU) tick(instr *Instruction) bool {
	(*instr).clock += 1
	(*cpu).clock += 1
	done := instr.clock == instr.cycles
	switch (*instr).instType {
	case noop: {}
	case addx: { if done { (*cpu).x += instr.value } }
	}
	return done
}

type CRT struct {
	screen [][]rune
	col int
	row int
}

func (crt *CRT) draw(cpu *CPU) {
	if len(crt.screen) <= crt.row {
		(*crt).screen = append((*crt).screen, make([]rune, 40))
	}

	if cpu.x - 1 <= crt.col && crt.col <= cpu.x+1 {
		(*crt).screen[crt.row][crt.col-1] = '#'
	} else {
		(*crt).screen[crt.row][crt.col-1] = '.'
	}

	if crt.col == 40 {
		(*crt).row += 1
		(*crt).col = 1
	} else {
		(*crt).col += 1
	}

}

func (crt CRT) print() {
	for _,row := range crt.screen {
		for _,r := range row {
			fmt.Printf("%s", string(r))
		}
		fmt.Print("\n")
	}
}

func t01() {
	lines := helper.Split(helper.ReadLines("aoc/10/10.inp"), "\n")
	instructions := helper.Collect(lines, func(s string) Instruction {return parseInstruction(s)})
	cpu := CPU{1,1}

	sum := 0
	current_instruction := 0
	for current_instruction < len(instructions) {
		if cpu.tick(&instructions[current_instruction]) { current_instruction += 1 }
		if (cpu.clock-20) % 40 == 0 { sum += cpu.x*cpu.clock }
	}
	fmt.Println(sum)
}

func t02() {
	lines := helper.Split(helper.ReadLines("aoc/10/10.inp"), "\n")
	instructions := helper.Collect(lines, func(s string) Instruction {return parseInstruction(s)})
	cpu := CPU{1,1}
	crt := CRT{screen: make([][]rune, 0), col:1, row: 0}
	crt.screen = append(crt.screen, make([]rune, 40))

	current_instruction := 0
	for current_instruction < len(instructions) {
		if cpu.tick(&instructions[current_instruction]) { current_instruction += 1 }
		crt.draw(&cpu)
	}
	crt.print()
}

func Run() {
	t01()
	t02()
}