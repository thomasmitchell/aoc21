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
		panic("failed to read file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	chain := NewPolymerChain([]byte(lines[0]))
	ruleset := NewPairInsertionRuleset(lines[2:])

	const numRounds = 10
	for i := 0; i < numRounds; i++ {
		nextPolymer := chain
		for nextPolymer != nil {
			nextPolymer = nextPolymer.Insert(ruleset)
		}
	}

	counts := chain.Count()
	var mostCommon int64
	var leastCommon int64 = math.MaxInt64
	for _, count := range counts {
		mostCommon = max(mostCommon, count)
		leastCommon = min(leastCommon, count)
	}

	fmt.Println(mostCommon - leastCommon)
}

type PolymerChain struct {
	next    *PolymerChain
	payload byte
}

func NewPolymerChain(template []byte) *PolymerChain {
	ret := &PolymerChain{
		payload: template[0],
	}

	cur := ret
	for _, polymer := range template[1:] {
		cur.next = &PolymerChain{payload: polymer}
		cur = cur.next
	}

	return ret
}

//Insert puts a new polymer into the chain if the current pair matches a
// rule in the given ruleset, and returns the next item in the chain to insert,
// skipping over an inserted object if there was one.
func (p *PolymerChain) Insert(rules *PairInsertionRuleset) *PolymerChain {
	ret := p.next
	if p.next != nil {
		toInsert := rules.Lookup([]byte{p.payload, p.next.payload})
		if toInsert != 0 {
			p.next = &PolymerChain{payload: toInsert, next: ret}
		}
	}

	return ret
}

func (p *PolymerChain) Count() map[byte]int64 {
	ret := map[byte]int64{}
	cur := p
	for cur != nil {
		ret[cur.payload]++
		cur = cur.next
	}

	return ret
}

func (p *PolymerChain) String() string {
	ret := []byte{}
	cur := p
	for cur != nil {
		ret = append(ret, cur.payload)
		cur = cur.next
	}

	return string(ret)
}

type PairInsertionRuleset struct {
	rules map[string]byte
}

func NewPairInsertionRuleset(rules []string) *PairInsertionRuleset {
	ret := &PairInsertionRuleset{rules: make(map[string]byte, len(rules))}
	for _, rule := range rules {
		parts := strings.Split(rule, " -> ")
		ret.rules[parts[0]] = []byte(parts[1])[0]
	}

	return ret
}

func (p *PairInsertionRuleset) Lookup(pair []byte) byte {
	return p.rules[string(pair)]
}

func max(i, j int64) int64 {
	if j > i {
		i = j
	}

	return i
}

func min(i, j int64) int64 {
	if j < i {
		i = j
	}

	return i
}
