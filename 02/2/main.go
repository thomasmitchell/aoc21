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
		panic("Couldn't read file")
	}

	commands := strings.Split(strings.TrimSpace(string(input)), "\n")

	var depth, horizontal, aim int64

	for _, command := range commands {
		commandParts := strings.Split(command, " ")
		dir := commandParts[0]
		dist, err := strconv.ParseInt(commandParts[1], 10, 32)
		if err != nil {
			panic("couldn't parse the distance")
		}
		switch dir {
		case "forward":
			horizontal += dist
			depth += dist * aim
		case "down":
			aim += dist
		case "up":
			aim -= dist
		default:
			panic("unknown command")
		}
	}

	fmt.Println(depth * horizontal)
}
