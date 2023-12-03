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
	scheme := make([][]byte, len(lines))
	for i, line := range lines {
		scheme[i] = []byte(line)
	}

	return scheme
}

func process(scheme [][]byte) int {
	partNumbers := parsePartNumbers(scheme)
	sum := 0
	for _, pn := range partNumbers {
		sum += pn
	}

	return sum
}

func parsePartNumbers(scheme [][]byte) []int {
	var partNumbers []int
	for y, line := range scheme {
		for x := 0; x < len(line); x++ {
			num, length, found := parsePartNumber(scheme, x, y)
			if found {
				partNumbers = append(partNumbers, num)
			}
			x += length
		}
	}

	return partNumbers
}

func parsePartNumber(scheme [][]byte, x, y int) (num int, length int, found bool) {
	line := scheme[y]

	if _, ok := parseDigit(line[x]); !ok {
		return 0, 0, false
	}

	for i := x; i < len(line); i++ {
		d, ok := parseDigit(line[i])
		if !ok {
			break
		}
		num = num*10 + d
		length++
	}

	return num, length, symbolsAround(scheme, x, x+length-1, y)
}

func symbolsAround(scheme [][]byte, x1, x2, y int) bool {
	x1 = max(0, x1-1)
	x2 = min(len(scheme[y])-1, x2+1)
	if symbol(scheme[y][x1]) || symbol(scheme[y][x2]) {
		return true
	}

	for x := x1; x <= x2; x++ {
		if y > 0 && symbol(scheme[y-1][x]) {
			return true
		}
		if y < len(scheme)-1 && symbol(scheme[y+1][x]) {
			return true
		}
	}

	return false
}

func parseDigit(v byte) (int, bool) {
	if v < '0' || v > '9' {
		return 0, false
	}

	return int(v - '0'), true
}

func symbol(v byte) bool {
	return v != '.' && (v < '0' || v > '9')
}
