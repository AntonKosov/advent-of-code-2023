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
	ratios := findRatios(scheme)
	sum := 0
	for _, ratio := range ratios {
		sum += ratio
	}

	return sum
}

func findRatios(scheme [][]byte) []int {
	var ratios []int
	for y, line := range scheme {
		for x, v := range line {
			if v != '*' {
				continue
			}
			if nums := findExactlyTwoNumbersAround(scheme, x, y); nums != nil {
				ratios = append(ratios, nums[0]*nums[1])
			}
		}
	}

	return ratios
}

func findExactlyTwoNumbersAround(scheme [][]byte, x, y int) []int {
	minX, maxX := max(0, x-1), min(len(scheme[y])-1, x+1)
	minY, maxY := max(0, y-1), min(len(scheme)-1, y+1)

	var numbers []int
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			num, lastX, found := parsePartNumber(scheme[y], x)
			if !found {
				continue
			}
			if len(numbers) == 2 {
				return nil
			}
			numbers = append(numbers, num)
			x = lastX + 1
		}
	}

	if len(numbers) < 2 {
		return nil
	}

	return numbers
}

func parsePartNumber(line []byte, x int) (num, lastX int, found bool) {
	if _, ok := parseDigit(line[x]); !ok {
		return 0, 0, false
	}

	for ; x > 0 && digit(line[x-1]); x-- {
	}

	for i := x; i < len(line); i++ {
		d, ok := parseDigit(line[i])
		if !ok {
			break
		}
		num = num*10 + d
		lastX = i
	}

	return num, lastX, true
}

func parseDigit(v byte) (int, bool) {
	if !digit(v) {
		return 0, false
	}

	return int(v - '0'), true
}

func digit(v byte) bool {
	return v >= '0' && v <= '9'
}
