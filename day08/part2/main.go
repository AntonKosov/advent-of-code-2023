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
		nodes[value] = node{left: left, right: right, finish: value[2] == 'Z'}
	}

	return nodes, directions
}

func process(nodes map[string]node, directions []byte) int {
	loops := getLoops(nodes, directions)

	return aoc.LCM(loops...)
}

func getLoops(nodes map[string]node, directions []byte) []int {
	var loops []int
	for v := range nodes {
		if v[2] == 'A' {
			loops = append(loops, getLoop(v, nodes, directions))
		}
	}

	return loops
}

func getLoop(startNodeName string, nodes map[string]node, directions []byte) int {
	currentNode := nodes[startNodeName]
	for i := 0; ; i++ {
		dir := directions[i%len(directions)]
		if dir == 'L' {
			currentNode = nodes[currentNode.left]
		} else {
			currentNode = nodes[currentNode.right]
		}

		if currentNode.finish {
			return i + 1
		}
	}
}

type node struct {
	left   string
	right  string
	finish bool
}
