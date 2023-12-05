package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() ([]interval, []mapping) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	seedIntervals := parseSeedsIntervals(lines[0])

	lines = lines[2:]
	var mappings []mapping
	for i := 1; i < len(lines); i += 2 {
		var transformations []transformation
		for ; i < len(lines) && lines[i] != ""; i++ {
			nums := aoc.StrToInts(lines[i])
			t := transformation{
				sourceStart: nums[1],
				targetStart: nums[0],
				count:       nums[2],
			}
			transformations = append(transformations, t)
		}

		mappings = append(mappings, newMapping(transformations))
	}

	return seedIntervals, mappings
}

func parseSeedsIntervals(line string) []interval {
	nums := aoc.StrToInts(line)
	var intervals []interval
	for i := 0; i < len(nums); i += 2 {
		count := nums[i+1]
		start := nums[i]
		intervals = append(intervals, interval{start: start, end: start + count - 1})
	}

	return intervals
}

func process(seedIntervals []interval, mappings []mapping) int {
	minValue := math.MaxInt
	for _, seedInterval := range seedIntervals {
		targetIntervals := []interval{seedInterval}
		for _, m := range mappings {
			targetIntervals = m.convertIntervals(targetIntervals)
		}

		minTarget := math.MaxInt
		for _, t := range targetIntervals {
			minTarget = min(minTarget, t.start)
		}

		minValue = min(minValue, minTarget)
	}

	return minValue
}

type interval struct {
	start int
	end   int
}

func (i interval) valid() bool {
	return i.start <= i.end
}

type transformation struct {
	sourceStart int
	targetStart int
	count       int
}

func (t transformation) convert(source int) int {
	return t.targetStart + (source - t.sourceStart)
}

func (t transformation) sourceEnd() int {
	return t.sourceStart + t.count - 1
}

type mapping struct {
	transformations []transformation
}

func newMapping(transformations []transformation) mapping {
	slices.SortFunc(transformations, func(a, b transformation) int {
		return a.sourceStart - b.sourceStart
	})

	return mapping{transformations: transformations}
}

func (m mapping) convertIntervals(intervals []interval) []interval {
	var targets []interval
	for _, i := range intervals {
		targets = append(targets, m.convertInterval(i)...)
	}

	return targets
}

func (m mapping) convertInterval(sourceInterval interval) []interval {
	var targets []interval
	for i := 0; i < len(m.transformations) && sourceInterval.valid(); {
		t := m.transformations[i]
		if t.sourceStart > sourceInterval.end {
			break
		}

		if t.sourceEnd() < sourceInterval.start {
			i++
			continue
		}

		if sourceInterval.start < t.sourceStart {
			targets = append(targets, interval{
				start: sourceInterval.start,
				end:   t.sourceStart - 1,
			})
			sourceInterval.start = t.sourceStart
			continue
		}

		transformIntervalEnd := min(sourceInterval.end, t.sourceEnd())
		targets = append(targets, interval{
			start: t.convert(sourceInterval.start),
			end:   t.convert(transformIntervalEnd),
		})
		sourceInterval.start = transformIntervalEnd + 1

		i++
	}

	if sourceInterval.valid() {
		targets = append(targets, sourceInterval)
	}

	return targets
}
