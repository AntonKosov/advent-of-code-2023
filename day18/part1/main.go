package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []instruction {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " (")
		part0 := parts[0]
		instructions[i] = instruction{
			dir:    part0[0],
			meters: aoc.StrToInts(part0)[0],
		}
	}

	return instructions
}

func process(instructions []instruction) int {
	// Pick's theorem and Shoelace (Trapezoid) formula are used
	boundaryPoints := 4
	doubledInteriorPoints := 0
	currentPos := aoc.NewVector2(0, 0)
	for _, ins := range instructions {
		nextPos := currentPos.Add(dirs[ins.dir].Mul(ins.meters))
		doubledInteriorPoints += (currentPos.Y + nextPos.Y) * (currentPos.X - nextPos.X)
		currentPos = nextPos
		boundaryPoints += ins.meters
	}

	return aoc.Abs(doubledInteriorPoints/2) + boundaryPoints/2 - 1
}

type instruction struct {
	dir    byte
	meters int
}

var dirs map[byte]aoc.Vector2[int]

func init() {
	dirs = map[byte]aoc.Vector2[int]{
		'U': aoc.NewVector2(0, -1),
		'D': aoc.NewVector2(0, 1),
		'R': aoc.NewVector2(1, 0),
		'L': aoc.NewVector2(-1, 0),
	}
}
