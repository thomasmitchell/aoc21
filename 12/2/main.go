package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"unicode"
)

const (
	NameStart = "start"
	NameEnd   = "end"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read file: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	nodes := map[string]*Node{}
	//build graph
	for _, line := range lines {
		parts := strings.Split(line, "-")
		for _, part := range parts {
			if nodes[part] == nil {
				nodes[part] = NewNode(part)
			}
		}

		from, to := parts[0], parts[1]
		nodes[from].Connect(nodes[to])
	}

	paths := nodes[NameStart].Walk(nil)
	pathStrings := make([]string, len(paths))
	for i, path := range paths {
		pathStrings[i] = path.String()
	}

	sort.Strings(pathStrings)
	fmt.Println(strings.Join(pathStrings, "\n"))

	fmt.Println(len(paths))
}

type Node struct {
	payload     string
	connections []*Node
}

func NewNode(name string) *Node {
	return &Node{
		payload: name,
	}
}

func (n *Node) Connect(n2 *Node) {
	n.connections = append(n.connections, n2)
	n2.connections = append(n2.connections, n)
}

func (n *Node) Retraversable() bool      { return unicode.IsUpper(rune(n.payload[0])) }
func (n *Node) Start() bool              { return n.HasName(NameStart) }
func (n *Node) End() bool                { return n.HasName(NameEnd) }
func (n *Node) HasName(name string) bool { return n.Name() == name }
func (n *Node) Name() string             { return n.payload }

func (n *Node) Walk(walked Path) []Path {
	base := walked.Append(n)
	if n.End() {
		return []Path{base}
	}

	ret := []Path{}
	for _, connection := range n.connections {
		if connection.Start() ||
			(!connection.Retraversable() &&
				base.HasTraveled(connection) &&
				base.AnyNonRetraversableNodeTraveledTwice()) {
			continue
		}

		ret = append(ret, connection.Walk(base)...)
	}

	return ret
}

type Path []*Node

func (p Path) Append(n *Node) Path {
	ret := make(Path, len(p)+1)
	copy(ret, p)
	ret[len(ret)-1] = n
	return ret
}

func (p Path) HasTraveled(n *Node) bool {
	for _, walkedNode := range p {
		if walkedNode.HasName(n.Name()) {
			return true
		}
	}

	return false
}

func (p Path) AnyNonRetraversableNodeTraveledTwice() bool {
	walked := map[string]bool{}
	for _, walkedNode := range p {
		if walkedNode.Retraversable() {
			continue
		}

		if walked[walkedNode.Name()] {
			return true
		}

		walked[walkedNode.Name()] = true
	}

	return false
}

func (p Path) String() string {
	nodeNames := make([]string, len(p))
	for i, n := range p {
		nodeNames[i] = n.Name()
	}

	return strings.Join(nodeNames, ",")
}
