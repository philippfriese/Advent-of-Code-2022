package _12

import (
	"AoC2022/helper"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"image"
)

func id(i int, j int, len int) int {
	return i * len + j
}

func arc(i int, j int, ioff int, joff int, graph *dijkstra.Graph, field *[][]rune) {
	if (*field)[i+ioff][j+joff] - (*field)[i][j] <= 1 {
        _ = graph.AddArc( // what are errors?
            id(i, j, len((*field)[0])),
            id(i+ioff, j+joff, len((*field)[0])),
            1)
	}
}


func makeGraph() (*dijkstra.Graph, image.Point, image.Point, [][]rune) {
	graph := dijkstra.NewGraph() // I still refuse to write dijkstra
	field := make([][]rune, 0)
	var start image.Point
	var end image.Point

	for i, line := range helper.Split(helper.ReadLines("aoc/12/12.inp"), "\n") {
		field = append(field, make([]rune, len(line)))
		for j, v := range line {
			if v == 'S' {
				v = 'a'
				start = image.Point{X: i, Y: j}
			} else if v == 'E' {
				v = 'z'
				end = image.Point{X: i, Y: j}
			}
			field[i][j] = v
			graph.AddVertex(id(i, j, len(field[0])))
		}
	}

	for i := range field {
		for j := range field[i] {
			if i > 0 			   { arc(i, j, -1, 0, graph, &field) }
			if i < len(field)-1    { arc(i, j, +1, 0, graph, &field) }
			if j > 0 			   { arc(i, j, 0, -1, graph, &field) }
			if j < len(field[i])-1 { arc(i, j, 0, +1, graph, &field) }
		}
	}
	return graph, start, end, field
}

func t01() {
	graph, start, end, field := makeGraph()
	r, _ := graph.Shortest(id(start.X, start.Y, len(field[0])),
                           id(end.X,   end.Y,   len(field[0])))
	fmt.Println(r.Distance)
}

func t02() {
	graph, _, end, field := makeGraph()
	shortest := int64(999999)
	for i := range field {
		for j := range field[i] {
			if field[i][j] == 'a' {
				r, _ := graph.Shortest(id(i, j, len(field[0])),
					                   id(end.X, end.Y, len(field[0])))
				if r.Distance > 0 && r.Distance < shortest {
					shortest = r.Distance
				}
			}
		}
	}
	fmt.Println(shortest)
}

func Run() {
	t01()
	t02()
}
