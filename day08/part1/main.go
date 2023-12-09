package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (map[string]node, []byte) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	directions := []byte(lines[0])
	lines = lines[2:]
	nodes := make(map[string]node, len(lines))
	for _, line := range lines {
		value, left, right := line[:3], line[7:10], line[12:15]
		nodes[value] = node{left: left, right: right}
	}

	return nodes, directions
}

func process(nodes map[string]node, directions []byte) int {
	count := 0
	for i, nodeName := 0, "AAA"; nodeName != "ZZZ"; i = (i + 1) % len(directions) {
		count++
		n := nodes[nodeName]
		if directions[i] == 'L' {
			nodeName = n.left
		} else {
			nodeName = n.right
		}
	}

	return count
}

type node struct {
	left  string
	right string
}
