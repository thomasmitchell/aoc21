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
		panic("couldn't read file")
	}

	values := strings.Split(strings.TrimSpace(string(input)), "\n")

	intValues := make([]int32, len(values))

	for i := 0; i < len(values); i++ {
		val, err := strconv.ParseInt(values[i], 10, 32)
		if err != nil {
			panic("couldn't parse int " + values[i])
		}

		intValues[i] = int32(val)

	}

	var lastVal int32
	increments := 0

	for i := 2; i < len(values); i++ {
		thisVal := intValues[i] + intValues[i-1] + intValues[i-2]
		if thisVal > lastVal {
			increments++
		}

		lastVal = thisVal
	}

	fmt.Println(increments - 1)
}
