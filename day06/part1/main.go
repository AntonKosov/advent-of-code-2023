package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []race {
	lines := aoc.ReadAllInput()
	times := aoc.StrToInts(lines[0])
	distances := aoc.StrToInts(lines[1])

	races := make([]race, len(times))
	for i := range times {
		races[i] = race{time: times[i], distance: distances[i]}
	}

	return races
}

func process(races []race) int {
	mul := 1
	for _, r := range races {
		mul *= r.winningOptions()
	}

	return mul
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
