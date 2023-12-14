package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]byte {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	field := make([][]byte, len(lines))
	for i, line := range lines {
		field[i] = []byte(line)
	}

	return field
}

func process(field [][]byte) int {
	tiltNorth(field)

	totalLoad := 0
	for r, row := range field {
		countRound := 0
		for _, v := range row {
			if v == round {
				countRound++
			}
		}
		totalLoad += countRound * (len(field) - r)
	}

	return totalLoad
}

func tiltNorth(field [][]byte) {
	for c := 0; c < len(field[0]); c++ {
		emptyRow := -1
		for r := range field {
			switch field[r][c] {
			case empty:
				if emptyRow < 0 {
					emptyRow = r
				}
			case square:
				emptyRow = -1
			case round:
				if emptyRow >= 0 {
					field[emptyRow][c] = round
					field[r][c] = empty
					emptyRow++
				}
			}
		}
	}
}

const (
	empty  = '.'
	round  = 'O'
	square = '#'
)
