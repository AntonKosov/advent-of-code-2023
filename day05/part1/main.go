package main

import (
	"fmt"
	"math"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() ([]int, []mapping) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	seeds := aoc.StrToInts(lines[0])
	lines = lines[2:]

	var mappings []mapping
	for i := 1; i < len(lines); i += 2 {
		var m mapping
		for ; i < len(lines) && lines[i] != ""; i++ {
			nums := aoc.StrToInts(lines[i])
			m.addRange(nums[1], nums[0], nums[2])
		}

		mappings = append(mappings, m)
	}

	return seeds, mappings
}

func process(seeds []int, mappings []mapping) int {
	minValue := math.MaxInt
	for _, value := range seeds {
		for _, m := range mappings {
			value = m.convert(value)
		}

		minValue = min(minValue, value)
	}

	return minValue
}

type mappingRange struct {
	source int
	target int
	count  int
}

type mapping struct {
	ranges []mappingRange
}

func (m *mapping) addRange(source, target, count int) {
	m.ranges = append(m.ranges, mappingRange{
		source: source,
		target: target,
		count:  count,
	})
}

func (m *mapping) convert(source int) int {
	for _, r := range m.ranges {
		if source >= r.source && source < r.source+r.count {
			return r.target + (source - r.source)
		}
	}

	return source
}
