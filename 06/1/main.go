package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const numDays = 80

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	fishiesRaw := strings.Split(strings.TrimSpace(string(input)), ",")
	fishies := make([]*FishieGroup, 7)
	for i := int64(0); i < 7; i++ {
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
		toAppend := []*FishieGroup{}
		for _, group := range fishies {
			toAdd := group.Advance()
			if toAdd != nil {
				toAppend = append(toAppend, toAdd)
			}
		}

		fishies = append(fishies, toAppend...)
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

func (f *FishieGroup) Advance() *FishieGroup {
	var ret *FishieGroup
	if f.timer == 0 {
		ret = newFishieGroup(8, f.amount)
		f.timer = 7
	}

	f.timer--
	return ret
}

func (f *FishieGroup) Amount() int64 { return f.amount }
