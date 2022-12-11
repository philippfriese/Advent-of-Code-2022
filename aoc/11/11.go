package _11

import (
	"AoC2022/helper"
	"fmt"
	"sort"
	"strings"
)

type Operand uint8
const (
	Add Operand = iota
	Mul
	Mod
)

type Input struct {
	is_variable bool
	constant int
}
type Operation struct {
	lhs Input
	rhs Input
	operand Operand
}

func (o Operation) apply(input int) int {
	lhs := o.lhs.constant
	if o.lhs.is_variable { lhs = input }

	rhs := o.rhs.constant
	if o.rhs.is_variable { rhs = input }

	switch o.operand {
	case Add: return lhs + rhs
	case Mul: return lhs * rhs
	case Mod: return lhs % rhs
	}
	return 0
}

func (o Operation) applyMod(input int, mod int) int {
	lhs := o.lhs.constant
	if o.lhs.is_variable { lhs = input }

	rhs := o.rhs.constant
	if o.rhs.is_variable { rhs = input }

	switch o.operand {
	case Add: return (lhs + (rhs%mod)) % mod
	case Mul: return (lhs * (rhs%mod)) % mod
	}
	return 0
}

type Monkey struct {
	id int
	items [][]int
	operation Operation
	test Operation
	if_true int
	if_false int
	inspected_items int
}

func parseOperation(s string) Operation {
	tokens := helper.Split(s, " ")
	operation := Operation{}
	if tokens[0] == "old" {
		operation.lhs = Input{constant: 0, is_variable: true}
	} else {
		operation.lhs = Input{constant: helper.ConvInt(tokens[0]), is_variable: false}
	}

	switch tokens[1] {
	case "+": operation.operand = Add
	case "*": operation.operand = Mul
	}

	if tokens[2] == "old" {
		operation.rhs = Input{constant: 0, is_variable: true}
	} else {
		operation.rhs = Input{constant: helper.ConvInt(tokens[2]), is_variable: false}
	}

	return operation
}

func parseTest(s string) Operation {
	tokens := helper.Split(s, " ")
	return Operation{
		lhs: Input{constant: 0, is_variable: true},
		rhs: Input{constant: helper.ConvInt(tokens[2]), is_variable: false},
		operand: Mod}
}

func parseMonkey(s string, num_monkeys int) Monkey {
	lines := helper.Split(s, "\n")

	items := helper.Collect(helper.Split(helper.Split(lines[1], ":")[1],","), func(s string) []int {
		v := helper.ConvInt(strings.Trim(s," "))
		item := make([]int,0)
		for i:=0; i < num_monkeys; i++ {item = append(item, v)}
		return item
	})
	monkey := Monkey{
		id: helper.ConvInt(strings.Replace(helper.Split(lines[0]," ")[1], ":","", 1)),
		items: items,
		operation: parseOperation(helper.Split(lines[2],"= ")[1]),
		test: parseTest(helper.Split(lines[3],": ")[1]),
		if_true: helper.ConvInt(helper.Split(helper.Split(lines[4],":")[1], " ")[4]),
		if_false: helper.ConvInt(helper.Split(helper.Split(lines[5],":")[1], " ")[4])}
	return monkey
}


func (m *Monkey) inspect(monkeys *[]Monkey, div bool) {
	// due to not wanting to have two functions, one for p1 and p2 respectively, merge them and control via div parameter
	// if div: use good old /3 worry-management, else: use fancy modulo shizzle
	for _,item := range m.items {
		(*m).inspected_items += 1

		// GoLang, y u no have inline if?!
		for m_i := range item {
			if div { item[m_i] = m.operation.apply(item[m_i]) / 3
			} else { item[m_i] = m.operation.applyMod(item[m_i], (*monkeys)[m_i].test.rhs.constant) }

		}

		// spliteroo for p1/p2
		var test bool
		if div { test = m.test.apply(item[0]) == 0
		} else { test = item[m.id] == 0 }

		// throw item
		if test { (*monkeys)[m.if_true].items = append((*monkeys)[m.if_true].items, item)
		} else { (*monkeys)[m.if_false].items = append((*monkeys)[m.if_false].items, item) }
	}
	(*m).items = make([][]int,0)
}


func t01() {
	lines := helper.Split(helper.ReadLines("aoc/11/11.inp"), "\n\n")
	monkeys := helper.Collect(lines, func (line string) Monkey {return parseMonkey(line, 1)})

	for i := 0; i < 20; i++ {
		for m_i := range monkeys {
			monkeys[m_i].inspect(&monkeys,true)
		}
	}

	inspected := make([]int,0)
	for m_i := range monkeys {
		inspected = append(inspected, monkeys[m_i].inspected_items)
	}

	fmt.Println(inspected)
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	fmt.Println(inspected[0] * inspected[1])
}

func t02() {
	/*
	idea: instead of storing a number of reach item (representing its worry level), store
	      the modulo remainder for each monkey
	      example: item 79 => [79 % 23 = 10,
	 						   79 % 19 = 3,
	 						   79 % 13 = 3,
	 						   79 % 17 = 11]
	     when calculating the new worry level at monkey i, say 0, we do:
	     y = x*19 % 17 = x%17 * 19%17 %17
		 but, for x=79, we already know x%17 = 79%17, since we stored it in item[0]!
		 repeat this process for all other monkey indices:
		 item 79*19 => [79*19 % 23 = 6,
	                    79*19 % 19 = 0,
	                    79*19 % 13 = 6,
	                    79*19 % 17 = 5]
 	     now, for the next throw, we do the same thing. This item goes to monkey 3 (as !(6|23))
	     at monkey 3, we need to test (79*19)+3 % 17 = (79*19)%17 + 3%17 %17,
	     but we again know the value (79*19)%17, as it is exactly item[3]! rinse and repeat
	 */
	lines := helper.Split(helper.ReadLines("aoc/11/11.inp"), "\n\n")
	monkeys := helper.Collect(lines, func (line string) Monkey {return parseMonkey(line, len(lines))})
	for m_i := range monkeys { // initialise items by mod-ing them for all monkeys with the respective mod
		for m_j := range monkeys[m_i].items {
			for m_k := range monkeys[m_i].items[m_j] {
				monkeys[m_i].items[m_j][m_k] =  monkeys[m_i].items[m_j][m_k] % monkeys[m_k].test.rhs.constant
			}
		}
	}
	for i := 0; i < 10000; i++ {
		for m_i := range monkeys {
			monkeys[m_i].inspect(&monkeys,false)
		}
	}

	inspected := make([]int,0)
	for m_i := range monkeys {
		inspected = append(inspected, monkeys[m_i].inspected_items)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	fmt.Println(inspected[0] * inspected[1])
}

func Run() {
	t01()
	t02()
}
