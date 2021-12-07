package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const numDays = 256

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	fishiesRaw := strings.Split(strings.TrimSpace(string(input)), ",")
	fishies := [9]*FishieGroup{}
	for i := int64(0); i < 9; i++ {
		fishies[i] = newFishieGroup(i, 0)
	}

	for _, fishie := range fishiesRaw {
		timer, err := strconv.ParseInt(fishie, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not convert timer `%s' to int: %s", fishie, err))
		}

		fishies[timer].Add(1)
	}

	for i := 0; i < numDays; i++ {
		newFishies := int64(0)
		for _, group := range fishies {
			newFishies += group.Advance()
		}

		eightIdx := i % 9
		sixIdx := (i + 7) % 9
		fishies[eightIdx].Add(newFishies)
		fishies[sixIdx].Add(newFishies)
	}

	total := int64(0)
	for _, group := range fishies {
		total += group.Amount()
	}

	fmt.Println(total)
}

type FishieGroup struct {
	amount int64
	timer  int64
}

func newFishieGroup(timer, amount int64) *FishieGroup {
	return &FishieGroup{timer: timer, amount: amount}
}

func (f *FishieGroup) Add(num int64) { f.amount += num }

func (f *FishieGroup) Advance() int64 {
	var ret int64
	if f.timer == 0 {
		ret = f.amount
		f.amount = 0
		f.timer = 9
	}

	f.timer--
	return ret
}

func (f *FishieGroup) Amount() int64 { return f.amount }

func (f *FishieGroup) Timer() int64 { return f.timer }
