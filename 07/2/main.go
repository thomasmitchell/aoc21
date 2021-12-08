package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	/*
		bestI := -1
		val := math.MaxInt64
		for i := 0; i < 2000; i++ {
			thisVal := calcDifferences(positions, i)
			if thisVal < val {
				bestI = i
				val = thisVal
			}
		}

		fmt.Printf("Target: %d\n", bestI)
		fmt.Printf("Val: %d\n", val)
	*/
	target := average(positions)
	fmt.Printf("target: %d\n", target)
	floorDifference := calcDifferences(positions, target)
	ceilDifference := calcDifferences(positions, target+1)
	diff := floorDifference
	if ceilDifference < floorDifference {
		diff = ceilDifference
	}

	fmt.Printf("difference: %d\n", diff)
}

func average(in []int) int {
	var total float64
	for _, val := range in {
		total += float64(val)
	}

	return int(math.Floor(total / float64(len(in))))
}

func calcDifferences(in []int, position int) int {
	sum := 0
	for _, val := range in {
		lesser, greater := val, position
		if val > position {
			lesser, greater = position, val
		}
		sum += triangleNumber(greater - lesser)
	}

	return sum
}

func triangleNumber(i int) int {
	return (i * (i + 1)) / 2
}
