package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read from file")
	}

	pointCounts := map[string]int{}
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		points := calcLine(parts[0], parts[1])
		for _, point := range points {
			pointCounts[point]++
		}
	}

	result := 0
	for _, count := range pointCounts {
		if count > 1 {
			result++
		}
	}

	fmt.Println(result)
}

func calcLine(from, to string) []string {
	fromParts := strings.Split(from, ",")
	toParts := strings.Split(to, ",")
	fromX, fromY := parseInt(fromParts[0]), parseInt(fromParts[1])
	toX, toY := parseInt(toParts[0]), parseInt(toParts[1])
	if fromX != toX && fromY != toY {
		return nil
	}

	ret := []string{}
	for fromX != toX || fromY != toY {
		ret = append(ret, fmt.Sprintf("%d,%d", fromX, fromY))

		if fromX < toX {
			fromX++
		} else if fromX > toX {
			fromX--
		}

		if fromY < toY {
			fromY++
		} else if fromY > toY {
			fromY--
		}
	}

	ret = append(ret, fmt.Sprintf("%d,%d", toX, toY))
	//fmt.Printf("From: %s, To: %s\n", from, to)
	//fmt.Printf("Points: %+v\n", ret)
	return ret
}

func parseInt(s string) int64 {
	ret, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse int `%s': %s", s, err))
	}

	return ret
}
