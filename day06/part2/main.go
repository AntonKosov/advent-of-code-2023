package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() race {
	lines := aoc.ReadAllInput()

	return race{
		time:     combineNums(aoc.StrToInts(lines[0])),
		distance: combineNums(aoc.StrToInts(lines[1])),
	}
}

func combineNums(nums []int) int {
	var str strings.Builder
	for _, num := range nums {
		str.WriteString(strconv.Itoa(num))
	}

	return aoc.StrToInt(str.String())
}

func process(r race) int {
	return r.winningOptions()
}

type race struct {
	time     int
	distance int
}

func (r race) winningOptions() int {
	count := 0
	for i := 1; i < r.time; i++ {
		timeLeft := r.time - i
		dst := i * timeLeft
		if dst > r.distance {
			count++
		}
	}

	return count
}
