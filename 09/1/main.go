package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	width := len(lines[0])
	height := len(lines)
	grid := NewGrid(width, height)
	for y, line := range lines {
		for x, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				panic("failed to parse int")
			}

			grid.Set(val, x, y)
		}
	}

	risk := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid.IsLow(x, y) {
				risk += 1 + grid.Get(x, y)
			}
		}
	}

	fmt.Println(risk)
}

type Grid struct {
	data   []int
	width  int
	height int
}

func NewGrid(x, y int) *Grid {
	return &Grid{
		data:   make([]int, x*y),
		width:  x,
		height: y,
	}
}

func (g *Grid) Set(val, x, y int) {
	g.data[g.idx(x, y)] = val
}

func (g *Grid) Get(x, y int) int {
	if x >= g.width || y >= g.height || x < 0 || y < 0 {
		return 10
	}
	return g.data[g.idx(x, y)]
}

func (g *Grid) IsLow(x, y int) bool {
	this := g.Get(x, y)
	return this < g.Get(x+1, y) &&
		this < g.Get(x-1, y) &&
		this < g.Get(x, y+1) &&
		this < g.Get(x, y-1)
}

func (g *Grid) idx(x, y int) int {
	return (y * g.width) + x
}
