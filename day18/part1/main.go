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
	// todo remove rotation, it's always right
	terrain, internalCells := dig(instructions)
	for _, c := range internalCells {
		fill(terrain, c)
	}

	return len(terrain)
}

func fill(terrain map[aoc.Vector2[int]]bool, pos aoc.Vector2[int]) {
	if terrain[pos] {
		return
	}

	terrain[pos] = true
	for _, dir := range dirs {
		fill(terrain, pos.Add(dir))
	}
}

func dig(instructions []instruction) (terrain map[aoc.Vector2[int]]bool, internalCells []aoc.Vector2[int]) {
	leftCells := make([]aoc.Vector2[int], 0, len(instructions))
	rightCells := make([]aoc.Vector2[int], 0, len(instructions))
	terrain = make(map[aoc.Vector2[int]]bool, len(instructions))
	rotation := 0
	currentDir := aoc.NewVector2(1, 0)
	currentPos := aoc.NewVector2(0, 0)
	currentDirLetter := instructions[0].dir
	for _, instr := range instructions {
		rot := rotations[currentDirLetter][instr.dir]
		if rot != 0 {
			rotation += rot
			if rot < 0 {
				currentDir = currentDir.RotateLeft()
			} else {
				currentDir = currentDir.RotateRight()
			}
		}
		currentDirLetter = instr.dir
		for i := 0; i < instr.meters; i++ {
			currentPos = currentPos.Add(currentDir)
			terrain[currentPos] = true
			leftCells = append(leftCells, currentPos.Add(currentDir.RotateLeft()))
			rightCells = append(rightCells, currentPos.Add(currentDir.RotateRight()))
		}
	}

	internalCells = leftCells
	if rotation > 0 {
		internalCells = rightCells
	}

	return terrain, internalCells
}

type instruction struct {
	dir    byte
	meters int
}

var dirs []aoc.Vector2[int]
var rotations map[byte]map[byte]int

func init() {
	left, right := -1, 1
	rotations = map[byte]map[byte]int{
		'U': {'L': left, 'R': right},
		'D': {'L': right, 'R': left},
		'R': {'U': left, 'D': right},
		'L': {'U': right, 'D': left},
	}
	dirs = []aoc.Vector2[int]{
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
	}
}
