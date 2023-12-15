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

func read() []string {
	lines := aoc.ReadAllInput()

	return strings.Split(lines[0], ",")
}

func process(initSeq []string) int {
	sum := 0
	for _, s := range initSeq {
		sum += hash(s)
	}

	return sum
}

func hash(seq string) int {
	h := 0
	for _, v := range seq {
		h = ((h + int(v)) * 17) % 256
	}

	return h
}
