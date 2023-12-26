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

func read() map[string][]string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	connections := make(map[string][]string, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		connections[parts[0]] = strings.Split(parts[1], " ")
	}

	return connections
}

func process(connections map[string][]string) int {
	adjList := makeGraph(connections)
	edges := makeEdges(adjList)
	for i1 := 0; i1 < len(edges)-2; i1++ {
		e1 := edges[i1]
		delete(adjList[e1[0]], e1[1])
		delete(adjList[e1[1]], e1[0])
		for i2 := i1 + 1; i2 < len(edges)-1; i2++ {
			e2 := edges[i2]
			delete(adjList[e2[0]], e2[1])
			delete(adjList[e2[1]], e2[0])
			if removeBridge(adjList) {
				visited := make([]bool, len(adjList))
				size := islandSize(adjList, visited, 0)
				if size < len(adjList) {
					return size * (len(adjList) - size)
				}
			}
			adjList[e2[0]][e2[1]] = struct{}{}
			adjList[e2[1]][e2[0]] = struct{}{}
		}
		adjList[e1[0]][e1[1]] = struct{}{}
		adjList[e1[1]][e1[0]] = struct{}{}
	}

	panic("bridges not found")
}

func removeBridge(adjList []map[int]struct{}) bool {
	// It's based on Tarjan's Algorithm
	visited := make([]bool, len(adjList))
	disc := make([]int, len(adjList))
	low := make([]int, len(adjList))
	parent := make([]int, len(adjList))
	time := 0

	edge, found := findBridge(0, adjList, visited, disc, low, parent, &time)
	if !found {
		return false
	}

	delete(adjList[edge[0]], edge[1])
	delete(adjList[edge[1]], edge[0])

	return true
}

func findBridge(
	node int, adjList []map[int]struct{}, visited []bool, disc, low, parent []int, time *int,
) (edge [2]int, found bool) {
	visited[node] = true

	*time++
	disc[node] = *time
	low[node] = *time

	for adjNode := range adjList[node] {
		if !visited[adjNode] {
			parent[adjNode] = node
			edge, found = findBridge(adjNode, adjList, visited, disc, low, parent, time)
			if found {
				return edge, found
			}
			low[node] = min(low[node], low[adjNode])
			if low[adjNode] > disc[node] {
				edge[0] = node
				edge[1] = adjNode
				return edge, true
			}
			continue
		}

		if adjNode != parent[node] {
			low[node] = min(low[node], disc[adjNode])
		}
	}

	return edge, false
}

func islandSize(adjList []map[int]struct{}, visited []bool, node int) int {
	if visited[node] {
		return 0
	}

	visited[node] = true

	count := 1
	for adjNode := range adjList[node] {
		count += islandSize(adjList, visited, adjNode)
	}

	return count
}

func makeEdges(adjList []map[int]struct{}) [][2]int {
	var edges [][2]int
	for i, adj := range adjList {
		for j := range adj {
			if i < j {
				edges = append(edges, [2]int{i, j})
			}
		}
	}

	return edges
}

func makeGraph(connections map[string][]string) (adjList []map[int]struct{}) {
	nameToIdx := make(map[string]int, len(connections))
	getIndex := func(component string) int {
		if idx, ok := nameToIdx[component]; ok {
			return idx
		}
		idx := len(adjList)
		nameToIdx[component] = idx
		adjList = append(adjList, map[int]struct{}{})
		return idx
	}
	for component, linkedComponents := range connections {
		compIdx := getIndex(component)
		for _, lc := range linkedComponents {
			lcIdx := getIndex(lc)
			adjList[compIdx][lcIdx] = struct{}{}
			adjList[lcIdx][compIdx] = struct{}{}
		}
	}

	return adjList
}
