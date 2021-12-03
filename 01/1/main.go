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

	var lastVal int64
	increments := 0
	for _, val := range values {
		thisVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic("couldn't parse int " + val)
		}

		if thisVal > lastVal {
			increments++
		}

		lastVal = thisVal
	}

	fmt.Println(increments - 1)
}
