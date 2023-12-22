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

func read() ([][]byte, aoc.Vector2[int]) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	var start aoc.Vector2[int]
	garden := make([][]byte, len(lines))
	for i, line := range lines {
		startIndex := strings.IndexRune(line, cellStart)
		row := []byte(line)
		if startIndex >= 0 {
			start = aoc.NewVector2(startIndex, i)
			row[startIndex] = cellEmpty
		}
		garden[i] = row
	}

	return garden, start
}

func process(garden [][]byte, start aoc.Vector2[int]) int {
	dirs := []aoc.Vector2[int]{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	currentPositions := map[aoc.Vector2[int]]struct{}{start: {}}
	nextPositions := map[aoc.Vector2[int]]struct{}{}
	width, height := len(garden[0]), len(garden)

	for i := 0; i < 64; i++ {
		for p := range currentPositions {
			for _, dir := range dirs {
				np := p.Add(dir)
				if np.X < 0 || np.Y < 0 || np.X >= width || np.Y >= height {
					continue
				}
				if garden[np.Y][np.X] == cellEmpty {
					nextPositions[np] = struct{}{}
				}
			}
		}
		currentPositions, nextPositions = nextPositions, currentPositions
		clear(nextPositions)
	}

	return len(currentPositions)
}

const (
	cellEmpty = '.'
	cellRock  = '#'
	cellStart = 'S'
)
