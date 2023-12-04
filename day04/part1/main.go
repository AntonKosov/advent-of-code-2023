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

func read() []card {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	cards := make([]card, len(lines))
	for i, line := range lines {
		cards[i] = parseCard(line)
	}

	return cards
}

func process(cards []card) int {
	sum := 0
	for _, c := range cards {
		sum += c.points()
	}

	return sum
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

func (c card) points() int {
	p := 0
	for _, num := range c.numbers {
		if !c.winningNumbers[num] {
			continue
		}
		if p == 0 {
			p = 1
			continue
		}
		p *= 2
	}

	return p
}
