package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := aoc.ReadAllInput()
	return lines[:len(lines)-1]
}

func process(data []string) int {
	restored := restore(data)
	sum := 0
	for _, v := range restored {
		sum += v
	}

	return sum
}

func restore(data []string) []int {
	restored := make([]int, len(data))

	for i, line := range data {
		restored[i] = firstDigit(line)*10 + lastDigit(line)
	}

	return restored
}

func firstDigit(line string) int {
	for _, c := range line {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}

	panic("first digit not found")
}

func lastDigit(line string) int {
	lineChars := []byte(line)
	for i := len(lineChars) - 1; i >= 0; i-- {
		c := lineChars[i]
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}

	panic("last digit not found")
}
