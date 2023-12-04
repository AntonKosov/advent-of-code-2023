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

func read() []int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	points := make([]int, len(lines))
	for i, line := range lines {
		c := parseCard(line)
		points[i] = c.matching()
	}

	return points
}

func process(points []int) int {
	copies := make([]int, len(points))
	for i := range copies {
		copies[i] = 1
	}

	for i, p := range points {
		addCopies(i, p, copies)
	}

	sum := 0
	for _, c := range copies {
		sum += c
	}

	return sum
}

func addCopies(cardIndex, count int, copies []int) {
	cards := copies[cardIndex]
	for i := cardIndex + 1; i < len(copies) && i <= cardIndex+count; i++ {
		copies[i] += cards
	}
}

type card struct {
	winningNumbers map[int]bool
	numbers        []int
}

func parseCard(line string) card {
	parts := strings.Split(line, ": ")
	parts = strings.Split(parts[1], "|")
	c := card{
		winningNumbers: map[int]bool{},
		numbers:        aoc.StrToInts(parts[1]),
	}
	for _, num := range aoc.StrToInts(parts[0]) {
		c.winningNumbers[num] = true
	}

	return c
}

func (c card) matching() int {
	matches := 0
	for _, num := range c.numbers {
		if !c.winningNumbers[num] {
			continue
		}
		matches++
	}

	return matches
}
