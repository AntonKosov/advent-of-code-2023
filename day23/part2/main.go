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
		for j, v := range maze[i] {
			if v != cellEmpty && v != cellForest {
				maze[i][j] = cellEmpty
			}
		}
	}

	return maze
}

func process(maze [][]byte) int {
	adjList := makeGraph(maze)
	maxSteps := 0
	visited := make([]bool, len(adjList))

	dfs(adjList, visited, 0, 0, &maxSteps)

	return maxSteps
}

func dfs(adjList []map[int]int, visited []bool, currentNode, distance int, maxSteps *int) {
	adjNodes := adjList[currentNode]
	for adjNode, dst := range adjNodes {
		if visited[adjNode] {
			continue
		}
		if adjNode == len(adjList)-1 {
			*maxSteps = max(*maxSteps, distance+dst)
			continue
		}
		visited[adjNode] = true
		dfs(adjList, visited, adjNode, distance+dst, maxSteps)
		visited[adjNode] = false
	}
}

func makeGraph(maze [][]byte) (adjList []map[int]int) {
	nodes := findNodes(maze)
	adjList = make([]map[int]int, len(nodes))
	for i := range adjList {
		adjList[i] = map[int]int{}
	}

	fillAdjList(maze, nodes, adjList, aoc.NewVector2(1, 0), aoc.NewVector2(1, -1), aoc.NewVector2(1, 0), 0)

	return adjList
}

func fillAdjList(
	maze [][]byte, nodes map[aoc.Vector2]int, adjList []map[int]int,
	from aoc.Vector2, prevPos, currentPos aoc.Vector2, distance int,
) {
	idxFrom := nodes[from]
	for _, dir := range dirs {
		np := currentPos.Add(dir)
		if np.Y < 0 || np.Y >= len(maze) || np == prevPos || maze[np.Y][np.X] == cellForest {
			continue
		}

		if idxTo, ok := nodes[np]; ok {
			if _, connected := adjList[idxFrom][idxTo]; connected {
				continue
			}
			adjList[idxFrom][idxTo] = distance + 1
			adjList[idxTo][idxFrom] = distance + 1
			fillAdjList(maze, nodes, adjList, np, currentPos, np, 0)
			continue
		}

		fillAdjList(maze, nodes, adjList, from, currentPos, np, distance+1)
	}
}

func findNodes(maze [][]byte) map[aoc.Vector2]int {
	nodes := map[aoc.Vector2]int{}
	for r, row := range maze {
		for c, v := range row {
			if v == cellForest {
				continue
			}

			pos := aoc.NewVector2(c, r)
			if r == 0 || r == len(row)-1 {
				nodes[pos] = len(nodes)
			}

			emptyAround := 0
			for _, dir := range dirs {
				p := pos.Add(dir)
				if p.Y < 0 || p.Y >= len(maze) || maze[p.Y][p.X] == cellEmpty {
					emptyAround++
				}
			}

			if emptyAround > 2 {
				nodes[pos] = len(nodes)
			}
		}
	}

	return nodes
}

const (
	cellEmpty  = '.'
	cellForest = '#'
)

var dirs []aoc.Vector2

func init() {
	dirs = []aoc.Vector2{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
}
