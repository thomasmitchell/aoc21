package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var terminators = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	scores := []int{}
	for _, line := range lines {
		p := NewParser(line)
		s := p.Parse()
		if s > 0 {
			scores = append(scores, s)
		}
	}

	sort.Ints(scores)
	fmt.Println(scores[(len(scores) / 2)])
}

type Parser struct {
	s string
}

func NewParser(s string) *Parser {
	return &Parser{s: s}
}

func (p *Parser) Parse() int {
	stack := []Chunk{}
	for _, c := range p.s {
		switch c {
		case '(':
			stack = append(stack, NewParenChunk())
		case '[':
			stack = append(stack, NewSquareChunk())
		case '{':
			stack = append(stack, NewCurlyChunk())
		case '<':
			stack = append(stack, NewAngleChunk())
		case ')', ']', '}', '>':
			if len(stack) == 0 || !stack[len(stack)-1].Close(c) {
				return 0
			}

			stack = stack[:len(stack)-1]
		}
	}

	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += terminators[stack[i].closeDelim]
	}

	return score
}

type Chunk struct {
	closeDelim rune
}

func NewChunk(close rune) Chunk {
	return Chunk{
		closeDelim: close,
	}
}

func (c Chunk) Close(with rune) bool {
	return c.closeDelim == with
}

func NewParenChunk() Chunk {
	return NewChunk(')')
}

func NewSquareChunk() Chunk {
	return NewChunk(']')
}

func NewCurlyChunk() Chunk {
	return NewChunk('}')
}

func NewAngleChunk() Chunk {
	return NewChunk('>')
}
