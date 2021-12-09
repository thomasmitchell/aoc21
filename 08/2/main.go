package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	sum := 0
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "|")
		inputs := strings.Split(strings.TrimSpace(parts[0]), " ")
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		decoder := newDecoder(inputs)
		value := decoder.decode(outputs)
		sum += value
	}

	fmt.Println(sum)
}

type displayDigit uint8

func parseDisplayDigit(s string) displayDigit {
	var ret displayDigit
	for _, c := range s {
		ret |= 1 << (c - 'a')
	}
	return ret
}

func (d displayDigit) and(d2 displayDigit) displayDigit {
	return d & d2
}

func (d displayDigit) or(d2 displayDigit) displayDigit {
	return d | d2
}

func (d displayDigit) len() int {
	ret := 0
	for i := 0; i < 7; i++ {
		ret += int((d & (1 << i)) >> i)
	}

	return ret
}

type Decoder struct {
	decoded [10]displayDigit
}

func newDecoder(inputs []string) *Decoder {
	ret := &Decoder{}
	parsedInputs := make([]displayDigit, 0, 6)
	for _, input := range inputs {
		dis := parseDisplayDigit(input)
		switch dis.len() {
		case 2 /*1*/ :
			ret.decoded[1] = dis
		case 4 /*4*/ :
			ret.decoded[4] = dis
		case 3 /*7*/ :
			ret.decoded[7] = dis
		case 7 /*8*/ :
			ret.decoded[8] = dis
		default:
			parsedInputs = append(parsedInputs, dis)
		}
	}

	remInputs := make([]displayDigit, 0, 3)
	for _, input := range parsedInputs {
		if input.len() != 6 {
			remInputs = append(remInputs, input)
			continue
		}

		//len is 6
		switch {
		case input.or(ret.decoded[4]).len() == 6:
			ret.decoded[9] = input
		case input.and(ret.decoded[1]).len() == 2:
			ret.decoded[0] = input
		default:
			ret.decoded[6] = input
		}
	}

	for _, input := range remInputs {
		//len is 5
		switch {
		case input.or(ret.decoded[1]).len() == 5:
			ret.decoded[3] = input
		case input.or(ret.decoded[6]).len() == 6:
			ret.decoded[5] = input
		default:
			ret.decoded[2] = input
		}
		continue
	}

	return ret
}

func (d *Decoder) decode(outputs []string) int {
	ret := 0
	lookup := d.genLookupTable()
	for i, output := range outputs {
		parsed := parseDisplayDigit(output)
		val := lookup[parsed]
		ret += val * int(math.Pow(10, float64(len(outputs)-(i+1))))
	}

	return ret
}

func (d *Decoder) genLookupTable() map[displayDigit]int {
	lookup := make(map[displayDigit]int, 10)
	for i, val := range d.decoded {
		lookup[val] = i
	}

	return lookup
}
