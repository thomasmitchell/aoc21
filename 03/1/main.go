package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Couldn't read input file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	lineLen := len(lines[0])
	bits := make([]int, lineLen)
	for _, line := range lines {
		for i, c := range line {
			if c == '0' {
				bits[i]--
			} else {
				bits[i]++
			}
		}
	}

	gamma := calc(bits)
	for i := range bits {
		bits[i] = -bits[i]
	}
	epsilon := calc(bits)
	fmt.Println(gamma * epsilon)
}

func calc(bits []int) int64 {
	var ret int64
	for i, bit := range bits {
		if bit < 0 {
			continue
		}
		ret |= 1 << (len(bits) - (i + 1))
	}

	return ret
}
