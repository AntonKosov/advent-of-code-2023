package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	data := make([][]int, len(lines))
	for i, line := range lines {
		data[i] = aoc.StrToInts(line)
	}

	return data
}

func process(data [][]int) int {
	sum := 0
	for _, line := range data {
		sum += extrapolate(line)
	}

	return sum
}

func extrapolate(line []int) int {
	lastNumbers := findLastNumbers(line)
	v := 0
	for i := len(lastNumbers) - 1; i >= 0; i-- {
		v += lastNumbers[i]
	}

	return v
}

func findLastNumbers(line []int) []int {
	var lastNumbers []int
	for i := len(line) - 1; i > 0; i-- {
		lastNumbers = append(lastNumbers, line[i])

		allZeros := true
		for j := 0; j < i; j++ {
			v := line[j+1] - line[j]
			allZeros = allZeros && v == 0
			line[j] = v
		}

		if allZeros {
			return lastNumbers
		}
	}

	panic("invalid input")
}
