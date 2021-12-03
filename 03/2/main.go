package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

//FIXME: Need to discard the strings that don't contain the bit before counting bits again
func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Couldn't read input file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	sort.Strings(lines)
	oxy := buildInt(getValue(lines, false))
	co2 := buildInt(getValue(lines, true))
	fmt.Println(oxy * co2)
}

func getValue(lines []string, less bool) string {
	lineLen := len(lines[0])
	for i := 0; i < lineLen; i++ {
		if len(lines) <= 1 {
			break
		}
		zero := keepZero(lines, i)
		if less {
			zero = !zero
		}
		lines = cut(lines, i, zero)
	}

	fmt.Println(lines[0])
	return lines[0]
}

func keepZero(lines []string, idx int) bool {
	bit := lines[len(lines)/2][idx]
	if len(lines)%2 == 0 && lines[(len(lines)/2)-1][idx] != bit { //then it's a tie
		return false
	}

	return bit == '0'
}

func cut(lines []string, strIdx int, keepZero bool) []string {
	idx := sort.Search(len(lines), func(i int) bool {
		return lines[i][strIdx] == '1'
	})

	if keepZero {
		return lines[:idx]
	}

	return lines[idx:]
}

func buildInt(bin string) int {
	var ret int
	for i, c := range bin {
		if c == '0' {
			continue
		}

		ret |= 1 << (len(bin) - (i + 1))
	}

	return ret
}
