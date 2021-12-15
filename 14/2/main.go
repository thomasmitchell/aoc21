package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

func main() {
	in, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("failed to read file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	//chain := NewPolymerChain([]byte(lines[0]))
	chain := NewPolymerChain([]byte(lines[0]))
	ruleset := NewPairInsertionRuleset(lines[2:])

	const numRounds = 40
	for i := 0; i < numRounds; i++ {
		startedAt := time.Now()
		fmt.Printf("round %d: ", i+1)
		chain.Insert(ruleset)
		fmt.Println(time.Since(startedAt).String())
	}

	fmt.Println(Score(chain.Count()))
}

func Score(counts map[byte]int64) int64 {
	var mostCommon int64
	var leastCommon int64 = math.MaxInt64
	for _, count := range counts {
		mostCommon = max(mostCommon, count)
		leastCommon = min(leastCommon, count)
	}

	return mostCommon - leastCommon
}

type PolymerChain struct {
	pairs map[string]int64
	first byte
	last  byte
}

func NewPolymerChain(template []byte) *PolymerChain {
	ret := &PolymerChain{
		pairs: map[string]int64{},
		first: template[0],
		last:  template[len(template)-1],
	}
	for i := 0; i < len(template)-1; i++ {
		ret.bumpPair(template[i], template[i+1], 1)
	}

	return ret
}

func (p *PolymerChain) bumpPair(b1, b2 byte, num int64) {
	p.pairs[string([]byte{b1, b2})] += num
}

//Insert puts a new polymer into the chain if the current pair matches a
// rule in the given ruleset, and returns the next item in the chain to insert,
// skipping over an inserted object if there was one.
func (p *PolymerChain) Insert(rules *PairInsertionRuleset) {
	next := *p
	next.pairs = map[string]int64{}
	for pair, num := range p.pairs {
		toInsert := rules.Lookup(pair)
		next.bumpPair(pair[0], toInsert, num)
		next.bumpPair(toInsert, pair[1], num)
	}

	*p = next
}

func (p *PolymerChain) Count() map[byte]int64 {
	ret := map[byte]int64{}
	for pair, num := range p.pairs {
		ret[pair[0]] += num
		ret[pair[1]] += num
	}

	for token := range ret {
		ret[token] /= 2
	}

	ret[p.first]++
	ret[p.last]++

	return ret
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

func (p *PairInsertionRuleset) Lookup(pair string) byte {
	return p.rules[pair]
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
