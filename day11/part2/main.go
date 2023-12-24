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
	image := make([][]byte, len(lines))
	for i, line := range lines {
		image[i] = []byte(line)
	}

	return image
}

func process(image [][]byte) int {
	positions := galaxiesPositions(image)
	distance := 0
	for i := 0; i < len(positions)-1; i++ {
		g := positions[i]
		for j := i + 1; j < len(positions); j++ {
			distance += g.ManhattanDst(positions[j])
		}
	}

	return distance
}

func galaxiesPositions(image [][]byte) []aoc.Vector2[int] {
	ci, ri := columnIndeces(image), rowsIndeces(image)
	var galaxies []aoc.Vector2[int]
	for r, row := range image {
		for c, v := range row {
			if v == galaxy {
				galaxies = append(galaxies, aoc.NewVector2(ri[r], ci[c]))
			}
		}
	}

	return galaxies
}

func columnIndeces(image [][]byte) []int {
	indeces := make([]int, len(image[0]))
	idx := -1
	for c := range indeces {
		offset := scale
		for r := range image {
			if image[r][c] == galaxy {
				offset = 1
				break
			}
		}
		idx += offset
		indeces[c] = idx
	}

	return indeces
}

func rowsIndeces(image [][]byte) []int {
	indeces := make([]int, len(image))
	idx := -1
	for r, row := range image {
		offset := scale
		for _, v := range row {
			if v == galaxy {
				offset = 1
				break
			}
		}
		idx += offset
		indeces[r] = idx
	}

	return indeces
}

const galaxy = '#'

const scale = 1_000_000
