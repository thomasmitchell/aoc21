package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	numEasyDigits := 0

	lines := strings.Split(strings.TrimSpace(string(in)), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "|")
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, output := range outputs {
			switch len(output) {
			case 2 /*1*/, 4 /*4*/, 3 /*7*/, 7 /*8*/ :
				numEasyDigits += 1
			}
		}
	}

	fmt.Println(numEasyDigits)
}
