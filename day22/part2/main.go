package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []brick {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	bricks := make([]brick, len(lines))
	for i, line := range lines {
		bricks[i] = parseBrick(line)
	}

	return bricks
}

func process(bricks []brick) int {
	settle(bricks)

	count := 0
	for i := range bricks {
		ignored := map[int]bool{i: true}
		for j := i + 1; j < len(bricks); j++ {
			if canFall(bricks, j, ignored) {
				ignored[j] = true
				count++
			}
		}
	}

	return count
}

func settle(bricks []brick) {
	slices.SortFunc(bricks, func(a, b brick) int { return a.bottom.Z - b.bottom.Z })
	for i := range bricks {
		for canFall(bricks, i, nil) {
			bricks[i].moveDown()
		}
	}
}

func canFall(bricks []brick, testedIndex int, ignoredIndeces map[int]bool) bool {
	b := bricks[testedIndex]
	if b.bottom.Z == 1 {
		return false
	}
	for i := 0; i < testedIndex; i++ {
		if ignoredIndeces != nil && ignoredIndeces[i] {
			continue
		}
		if bricks[i].supports(b) {
			return false
		}
	}

	return true
}

type brick struct {
	bottom aoc.Vector3
	top    aoc.Vector3
}

func (b *brick) moveDown() {
	b.bottom.Z--
	b.top.Z--
}

func (b *brick) supports(b2 brick) bool {
	if b.top.Z != b2.bottom.Z-1 {
		return false
	}

	b1x1, b1x2, b1y1, b1y2 := b.bottom.X, b.top.X, b.bottom.Y, b.top.Y
	b1x1, b1x2 = min(b1x1, b1x2), max(b1x1, b1x2)
	b1y1, b1y2 = min(b1y1, b1y2), max(b1y1, b1y2)

	b2x1, b2x2, b2y1, b2y2 := b2.bottom.X, b2.top.X, b2.bottom.Y, b2.top.Y
	b2x1, b2x2 = min(b2x1, b2x2), max(b2x1, b2x2)
	b2y1, b2y2 = min(b2y1, b2y2), max(b2y1, b2y2)

	x1 := max(b1x1, b2x1)
	x2 := min(b1x2, b2x2)
	if x2-x1 < 0 {
		return false
	}

	y1 := max(b1y1, b2y1)
	y2 := min(b1y2, b2y2)

	return y2-y1 >= 0
}

func parseBrick(value string) brick {
	parts := strings.Split(value, "~")
	return brick{
		bottom: parsePosition(parts[0]),
		top:    parsePosition(parts[1]),
	}
}

func parsePosition(value string) aoc.Vector3 {
	c := aoc.StrToInts(value)
	return aoc.NewVector3(c[0], c[1], c[2])
}
