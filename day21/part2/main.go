package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
	"github.com/AntonKosov/advent-of-code-2023/day21/part2/solution"
)

func main() {
	garden, start := read()
	answer := solution.Count(garden, start, 26501365)
	fmt.Printf("Answer: %v\n", answer)
}

func read() ([][]byte, aoc.Vector2[int]) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	return solution.Parse(lines)
}
