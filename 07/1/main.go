package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to read file: %s", err))
	}

	positionsRaw := strings.Split(strings.TrimSpace(string(in)), ",")
	positions := make([]int, len(positionsRaw))
	for i, position := range positionsRaw {
		positionParsed, err := strconv.ParseInt(position, 10, 32)
		if err != nil {
			panic(fmt.Sprintf("failed to parse int `%s': %s", position, err))
		}

		positions[i] = int(positionParsed)
	}

	sort.Ints(positions)

	targetPosition := median(positions)
	fmt.Printf("Median: %d\n", targetPosition)

	fmt.Println(sumDifferences(positions, targetPosition))
}

func median(in []int) int {
	if len(in) == 0 {
		return 0
	}

	if len(in)%2 == 1 {
		return in[(len(in)/2)+1]
	}

	return (in[len(in)/2] + in[(len(in)/2)-1]) / 2
}

func sumDifferences(in []int, target int) int {
	ret := 0
	for _, val := range in {
		lesser, greater := val, target
		if target < val {
			lesser, greater = target, val
		}

		ret += greater - lesser
	}

	return ret
}
