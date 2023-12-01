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
	lineChars := []byte(line)
	for i := range lineChars {
		if d, ok := startsWithDigit(lineChars[i:]); ok {
			return d
		}
	}

	panic("first digit not found")
}

func lastDigit(line string) int {
	lineChars := []byte(line)
	for i := len(lineChars) - 1; i >= 0; i-- {
		if d, ok := startsWithDigit(lineChars[i:]); ok {
			return d
		}
	}

	panic("last digit not found")
}

func startsWithDigit(line []byte) (int, bool) {
	if c := line[0]; c >= '0' && c <= '9' {
		return int(c - '0'), true
	}

	for word, value := range digits {
		if startsWithWord(line, []byte(word)) {
			return value, true
		}
	}

	return 0, false
}

func startsWithWord(line, word []byte) bool {
	if len(line) < len(word) {
		return false
	}

	for i, v := range word {
		if line[i] != v {
			return false
		}
	}

	return true
}

var digits map[string]int

func init() {
	digits = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
}
