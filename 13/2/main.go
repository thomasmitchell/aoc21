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
	dots := []Point{}

	var dotLines, foldLines []string
	for i, line := range lines {
		if line == "" {
			dotLines = lines[:i]
			foldLines = lines[i+1:]
			break
		}
	}

	var maxX, maxY int
	for _, line := range dotLines {
		coordParts := strings.Split(line, ",")
		toAppend := Point{
			X: atoi(coordParts[0]),
			Y: atoi(coordParts[1]),
		}
		maxX = max(toAppend.X, maxX)
		maxY = max(toAppend.Y, maxY)
		dots = append(dots, toAppend)
	}

	grid := NewGrid(maxX+1, maxY+1)
	for _, dot := range dots {
		grid.Set(dot.X, dot.Y)
	}

	for _, line := range foldLines {
		line = strings.TrimPrefix(line, "fold along ")
		parts := strings.Split(line, "=")
		foldX := parts[0] == "x"
		along := atoi(parts[1])
		if foldX {
			grid.FoldAlongX(along)
		} else {
			grid.FoldAlongY(along)
		}
	}

	fmt.Println(grid.String())
}

type Grid struct {
	data   []bool
	width  int
	height int
}

func NewGrid(x, y int) *Grid {
	return &Grid{
		data:   make([]bool, x*y),
		width:  x,
		height: y,
	}
}

func (g *Grid) Set(x, y int) {
	g.data[g.idx(x, y)] = true
}

func (g *Grid) Get(x, y int) bool {
	if g.oob(x, y) {
		return false
	}
	return g.data[g.idx(x, y)]
}

func (g *Grid) idx(x, y int) int {
	return (y * g.width) + x
}

func (g *Grid) oob(x, y int) bool {
	return x >= g.width || y >= g.height || x < 0 || y < 0
}

func (g *Grid) FoldAlongX(foldX int) {
	newGrid := NewGrid(foldX, g.height)
	for y := 0; y < g.height; y++ {
		for x := 0; x < foldX; x++ {
			if g.Get(x, y) {
				newGrid.Set(x, y)
			}
		}
	}

	for y := 0; y < g.height; y++ {
		for x := foldX + 1; x < g.width; x++ {
			if g.Get(x, y) {
				newGrid.Set((2*foldX)-x, y)
			}
		}
	}

	*g = *newGrid
}

func (g *Grid) FoldAlongY(foldY int) {
	for y := foldY; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.Get(x, y) {
				g.Set(x, foldY-(y-foldY))
			}
		}
	}

	g.height = foldY
	g.data = g.data[:g.width*g.height]
}

func (g *Grid) Count() int {
	ret := 0
	for _, point := range g.data {
		if point {
			ret++
		}
	}

	return ret
}

func (g *Grid) String() string {
	ret := make([]byte, 0, (g.width*g.height)+g.height-1)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.Get(x, y) {
				ret = append(ret, '#')
			} else {
				ret = append(ret, '.')
			}
		}

		ret = append(ret, '\n')
	}

	return string(ret)
}

type Point struct {
	X int
	Y int
}

func atoi(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert string `%s' to int: %s", s, err))
	}

	return ret
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
