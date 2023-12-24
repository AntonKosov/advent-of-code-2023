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
	maze := make([][]byte, len(lines))
	for i, line := range lines {
		maze[i] = []byte(line)
	}

	return maze
}

func process(maze [][]byte) int {
	targetPos := aoc.NewVector2(len(maze[0])-2, len(maze)-1)
	steps := 0
	positions := []position{
		{
			current:  aoc.NewVector2(1, 0),
			previous: aoc.NewVector2(1, -1),
		},
	}
	var nextPositions []position
	for ; len(positions) > 0; steps++ {
		for _, pos := range positions {
			nextPositions = append(nextPositions, move(maze, pos)...)
		}
		positions = positions[:0]
		for _, np := range nextPositions {
			if np.current != targetPos {
				positions = append(positions, np)
			}
		}
		nextPositions = nextPositions[:0]
	}

	return steps
}

func move(maze [][]byte, pos position) []position {
	nextPositions := make([]position, 0, 3)
	if c := maze[pos.current.Y][pos.current.X]; c != cellEmpty {
		nextPositions = append(nextPositions, position{
			previous: pos.current,
			current:  pos.current.Add(slope[c]),
		})

		return nextPositions
	}

	for _, dir := range dirs {
		nextPos := pos.current.Add(dir)
		if nextPos == pos.previous {
			continue
		}
		c := maze[nextPos.Y][nextPos.X]
		if c == cellForest {
			continue
		}
		if c == cellEmpty || nextPos.Add(slope[c]) != pos.current {
			nextPositions = append(nextPositions, position{
				previous: pos.current,
				current:  nextPos,
			})
		}
	}

	return nextPositions
}

type position struct {
	current  aoc.Vector2[int]
	previous aoc.Vector2[int]
}

const (
	cellEmpty  = '.'
	cellForest = '#'
)

var (
	dirs  []aoc.Vector2[int]
	slope map[byte]aoc.Vector2[int]
)

func init() {
	dirs = []aoc.Vector2[int]{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	slope = map[byte]aoc.Vector2[int]{
		'>': aoc.NewVector2(1, 0),
		'<': aoc.NewVector2(-1, 0),
		'v': aoc.NewVector2(0, 1),
		'^': aoc.NewVector2(0, -1),
	}
}
