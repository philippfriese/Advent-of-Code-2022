package _14

import (
	"AoC2022/helper"
	"fmt"
	"image"
)

type Wall struct {
	x image.Point
	y image.Point
}

func makePoint(tuple string) image.Point {
	parts := helper.Split(tuple, ",")
	return image.Point{X: helper.ConvInt(parts[0]), Y: helper.ConvInt(parts[1])}
}

func toIndex(point image.Point, minX int, minY int) image.Point{
	return image.Point{point.X - minX, point.Y - minY}
}

func printField(field *[][]int, min_x int) {
	for i := range (*field) {
		for j := range (*field)[i] {
			if i == 0 && j == 500-min_x {
				fmt.Printf("+")
				continue
			}
			switch (*field)[i][j] {
			case 0: fmt.Printf(".")
			case 1: fmt.Printf("o")
			case 9: fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func get(field *[][]int, i int, j int, floor bool) int {
	if floor && i == len(*field)-1{
		fmt.Println("hit rock bottom")
		return 9
	}
	return (*field)[i][j]
}

func fall(field *[][]int, currentPosition image.Point, floor bool) bool {
	if !floor && currentPosition.Y >= len(*field)-1 { return false }


	if get(field, currentPosition.Y+1, currentPosition.X, floor) == 0 {
		return fall(field, image.Point{X: currentPosition.X, Y: currentPosition.Y+1}, floor)
	}
	if get(field, currentPosition.Y+1, currentPosition.X-1, floor) == 0 {
		return fall(field, image.Point{X: currentPosition.X-1, Y: currentPosition.Y+1}, floor)
	}
	if get(field, currentPosition.Y+1, currentPosition.X+1, floor) == 0 {
		return fall(field, image.Point{X: currentPosition.X+1, Y: currentPosition.Y+1}, floor)
	}
	(*field)[currentPosition.Y][currentPosition.X] = 1
	return true
}

func t01(segments *[]Wall, min_x int, min_y int, max_x int, max_y int) {
	width := max_x - min_x+1
	height := max_y + 1

	fmt.Printf("Width: %d, height: %d\n",width,height)

	field := make([][]int, height)
	for i := range field {
		field[i] = make([]int, width+1)
	}

	for _, segment := range *segments {
		start := toIndex(segment.x, min_x, min_y)
		end := toIndex(segment.y, min_x, min_y)
		if segment.x.X == segment.y.X {
			for i := helper.Min(start.Y, end.Y); i <= helper.Max(start.Y,end.Y); i++ {
				field[i][start.X] = 9
			}
		} else {
			for i := helper.Min(start.X, end.X); i <= helper.Max(start.X,end.X); i++ {
				field[start.Y][i] = 9
			}
		}
	}
	printField(&field, min_x)
	units := 0
	for fall(&field, toIndex(image.Point{500, 0}, min_x, min_y), false) {
		units += 1
	}
	printField(&field, min_x)
	fmt.Println(units)
}

func t02(segments *[]Wall, min_x int, min_y int, max_x int, max_y int) {
	height := max_y+3

	// set width according to expected height
	width := 2*height
	min_x = 500 - height
	fmt.Printf("Width: %d, height: %d\n",width,height)
	fmt.Printf("x: %d-%d, y: %d-%d\n", min_x, max_x, min_y, max_y)
	field := make([][]int, height)
	for i := range field {
		field[i] = make([]int, width)
	}

	for _, segment := range *segments {
		start := toIndex(segment.x, min_x, min_y)
		end := toIndex(segment.y, min_x, min_y)
		if segment.x.X == segment.y.X {
			for i := helper.Min(start.Y, end.Y); i <= helper.Max(start.Y,end.Y); i++ {
				field[i][start.X] = 9
			}
		} else {
			for i := helper.Min(start.X, end.X); i <= helper.Max(start.X,end.X); i++ {
				field[start.Y][i] = 9
			}
		}
	}
	printField(&field, min_x)
	units := 0
	start := toIndex(image.Point{500,0}, min_x, min_y)
	for field[start.Y][start.X] == 0 {
		fall(&field, start, true)
		units += 1
	}
	printField(&field, min_x)
	fmt.Println(units)
}

func Run() {
	lines := helper.Split(helper.ReadLines("aoc/14/14.inp"),"\n")

	segments := make([]Wall, 0)
	helper.Apply(lines, func (line string) {
		tokens := helper.Split(line, " -> ")
		for i := range tokens[1:] {
			segments = append(segments, Wall {
				x: makePoint(tokens[i]),
				y: makePoint(tokens[i+1]), // this is a bit backwards; range tokens[1:] starts at i=0
			})
		}
	})

	min_x := helper.MinArray(helper.Collect(segments, func(w Wall) int { return helper.Min(w.x.X, w.y.X) }))-1
	min_y := 0
	max_x := helper.MaxArray(helper.Collect(segments, func(w Wall) int { return helper.Max(w.x.X, w.y.X) }))
	max_y := helper.MaxArray(helper.Collect(segments, func(w Wall) int { return helper.Max(w.x.Y, w.y.Y) }))

	//t01(&segments, min_x, min_y, max_x, max_y)
	t02(&segments, min_x, min_y, max_x, max_y)
}
