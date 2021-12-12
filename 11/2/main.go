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
		panic("Unable to read file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	width := len(lines[0])
	height := len(lines)
	g := NewGrid(width, height)
	for y, line := range lines {
		for x, c := range line {
			cInt, err := strconv.Atoi(string(c))
			if err != nil {
				panic("Failed to parse int")
			}
			g.Set(cInt, x, y)
		}
	}

	var turn int = 1
	for {
		thisFlashed := g.BumpAll()
		turnFlashed := []Point{}

		for len(thisFlashed) > 0 {
			turnFlashed = append(turnFlashed, thisFlashed...)
			nextFlashed := []Point{}
			for _, toFlash := range thisFlashed {
				nextFlashed = append(nextFlashed, g.Flash(toFlash.X, toFlash.Y)...)
			}

			thisFlashed = nextFlashed
		}

		if len(turnFlashed) == width*height {
			break
		}

		for _, point := range turnFlashed {
			g.Set(0, point.X, point.Y)
		}

		turn++
	}

	fmt.Println(turn)
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
	if g.oob(x, y) {
		return
	}

	g.data[g.idx(x, y)] = val
}

func (g *Grid) Get(x, y int) int {
	if g.oob(x, y) {
		return -1
	}
	return g.data[g.idx(x, y)]
}

//returns true and sets to -1 if equal to 9 before bump
//values < 0 won't bump
func (g *Grid) Bump(x, y int) bool {
	i := g.Get(x, y)
	if i < 0 {
		return false
	}
	if i == 9 {
		i = -2
	}
	g.Set(i+1, x, y)
	return i == -2
}

func (g *Grid) BumpAll() []Point {
	ret := []Point{}
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			shouldFlash := g.Bump(x, y)
			if shouldFlash {
				ret = append(ret, Point{x, y})
			}
		}
	}

	return ret
}

func (g *Grid) Flash(x, y int) []Point {
	ret := []Point{}
	toBump := []Point{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
		{x + 1, y + 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
		{x - 1, y - 1},
	}

	for _, point := range toBump {
		shouldFlash := g.Bump(point.X, point.Y)
		if shouldFlash {
			ret = append(ret, point)
		}
	}

	return ret
}

func (g *Grid) idx(x, y int) int {
	return (y * g.width) + x
}

func (g *Grid) oob(x, y int) bool {
	return x >= g.width || y >= g.height || x < 0 || y < 0
}

type Point struct {
	X int
	Y int
}
