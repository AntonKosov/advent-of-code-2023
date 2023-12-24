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
	path, rot := findPath(maze, startPos)
	return countEnclosedArea(maze, path, rot < 0)
}

func countEnclosedArea(maze [][]byte, path []aoc.Vector2[int], leftRotation bool) int {
	visited := buildVisited(maze, path)
	count := 0
	for i := 1; i < len(path); i++ {
		prevPos, pos := path[i-1], path[i]
		incomingDir := pos.Sub(prevPos)
		enclosedDir := incomingDir.RotateRight()
		if leftRotation {
			enclosedDir = incomingDir.RotateLeft()
		}
		count += measureArea(maze, visited, prevPos.Add(enclosedDir))
		count += measureArea(maze, visited, pos.Add(enclosedDir))
	}

	return count
}

func buildVisited(maze [][]byte, path []aoc.Vector2[int]) [][]bool {
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	for _, p := range path {
		visited[p.Y][p.X] = true
	}

	return visited
}

func measureArea(maze [][]byte, visited [][]bool, pos aoc.Vector2[int]) int {
	if pos.X < 0 || pos.Y < 0 || pos.Y >= len(maze) || pos.X >= len(maze[0]) {
		return 0
	}

	if visited[pos.Y][pos.X] || maze[pos.Y][pos.X] == enclosedChar {
		return 0
	}

	maze[pos.Y][pos.X] = enclosedChar
	area := 1

	for _, dir := range dirs {
		area += measureArea(maze, visited, pos.Add(dir))
	}

	return area
}

func findPath(maze [][]byte, startPos aoc.Vector2[int]) (path []aoc.Vector2[int], rotationSum int) {
	prevPos := aoc.NewVector2(-1, -1)
	currentPos := startPos
	for len(path) == 0 || currentPos != startPos {
		neighboursDir := connections[maze[currentPos.Y][currentPos.X]]
		dir := neighboursDir[0]
		if currentPos.Add(dir.direction) == prevPos {
			dir = neighboursDir[1]
		}
		nextPos := currentPos.Add(dir.direction)
		rotationSum += dir.turn
		path = append(path, currentPos)
		prevPos = currentPos
		currentPos = nextPos
	}

	return path, rotationSum
}

func setStartPipe(maze [][]byte, startPos aoc.Vector2[int]) {
	height, width := len(maze), len(maze[0])
	for pipe, outputs := range connections {
		count := 0
		for _, output := range outputs {
			pos := startPos.Add(output.direction)
			if pos.X < 0 || pos.Y < 0 || pos.X >= width || pos.Y >= height {
				break
			}
			outputs2 := connections[maze[pos.Y][pos.X]]
			if pos.Add(outputs2[0].direction) == startPos || pos.Add(outputs2[1].direction) == startPos {
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

const enclosedChar = 'I'

var connections map[byte][2]rotation
var dirs []aoc.Vector2[int]

type rotation struct {
	direction aoc.Vector2[int]
	turn      int
}

func init() {
	up := aoc.NewVector2(0, -1)
	down := up.Mul(-1)
	left := aoc.NewVector2(-1, 0)
	right := left.Mul(-1)

	dirs = []aoc.Vector2[int]{up, down, left, right}

	direct, leftTurn, rightTurn := 0, -1, 1

	connections = map[byte][2]rotation{
		'|': {{direction: up, turn: direct}, {direction: down, turn: direct}},
		'-': {{direction: left, turn: direct}, {direction: right, turn: direct}},
		'L': {{direction: up, turn: rightTurn}, {direction: right, turn: leftTurn}},
		'J': {{direction: up, turn: leftTurn}, {direction: left, turn: rightTurn}},
		'7': {{direction: left, turn: leftTurn}, {direction: down, turn: rightTurn}},
		'F': {{direction: down, turn: leftTurn}, {direction: right, turn: rightTurn}},
	}
}
