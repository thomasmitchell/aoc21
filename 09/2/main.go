package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	basins := []int{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			basins = append(basins, grid.GetBasin(x, y))
		}
	}

	sort.Ints(basins)
	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
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
	if g.oob(x, y) {
		return 9
	}
	return g.data[g.idx(x, y)]
}

func (g *Grid) GetBasin(x, y int) int {
	if g.Get(x, y) == 9 {
		return 0
	}

	g.Set(9, x, y)
	return 1 + g.GetBasin(x+1, y) + g.GetBasin(x-1, y) + g.GetBasin(x, y+1) + g.GetBasin(x, y-1)
}

func (g *Grid) idx(x, y int) int {
	return (y * g.width) + x
}

func (g *Grid) oob(x, y int) bool {
	return x >= g.width || y >= g.height || x < 0 || y < 0
}
