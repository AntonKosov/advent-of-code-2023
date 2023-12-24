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
	startPos := findStart(maze)
	setStartPipe(maze, startPos)
	length := loopLength(maze, startPos)
	return length / 2
}

func loopLength(maze [][]byte, startPos aoc.Vector2[int]) int {
	count := 1
	prevPos := aoc.NewVector2(-1, -1)
	currentPos := startPos
	for {
		neighboursDir := connections[maze[currentPos.Y][currentPos.X]]
		for _, dir := range neighboursDir {
			nextPos := currentPos.Add(dir)
			if nextPos == prevPos {
				continue
			}
			count++
			if nextPos == startPos {
				return count
			}
			prevPos = currentPos
			currentPos = nextPos
			break
		}
	}
}

func setStartPipe(maze [][]byte, startPos aoc.Vector2[int]) {
	height, width := len(maze), len(maze[0])
	for pipe, outputs := range connections {
		count := 0
		for _, output := range outputs {
			pos := startPos.Add(output)
			if pos.X < 0 || pos.Y < 0 || pos.X >= width || pos.Y >= height {
				break
			}
			outputs2 := connections[maze[pos.Y][pos.X]]
			if pos.Add(outputs2[0]) == startPos || pos.Add(outputs2[1]) == startPos {
				count++
			}
		}
		if count == 2 {
			maze[startPos.Y][startPos.X] = pipe
			return
		}
	}

	panic("start pipe not found")
}

func findStart(maze [][]byte) aoc.Vector2[int] {
	for r, row := range maze {
		for c, v := range row {
			if v == 'S' {
				return aoc.NewVector2(c, r)
			}
		}
	}

	panic("start not found")
}

var connections map[byte][2]aoc.Vector2[int]

func init() {
	down := aoc.NewVector2(0, 1)
	up := down.Mul(-1)
	right := aoc.NewVector2(1, 0)
	left := right.Mul(-1)

	connections = map[byte][2]aoc.Vector2[int]{
		'|': {up, down},
		'-': {left, right},
		'L': {up, right},
		'J': {up, left},
		'7': {left, down},
		'F': {down, right},
	}
}
