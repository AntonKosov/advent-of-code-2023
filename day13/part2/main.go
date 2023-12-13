package main

import (
	"fmt"
	"slices"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []pattern {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	var patterns []pattern
	for i := 0; i < len(lines); i++ {
		var p pattern
		for ; i < len(lines) && lines[i] != ""; i++ {
			p = append(p, []byte(lines[i]))
		}

		patterns = append(patterns, p)
	}

	return patterns
}

func process(patterns []pattern) int {
	sum := 0
	for _, p := range patterns {
		sum += summarize(p)
	}

	return sum
}

func summarize(p pattern) int {
	reverse := map[byte]byte{empty: rock, rock: empty}
	originalVM, originalHM := -1, -1
	if vm, ok := p.verticalMirror(nil); ok {
		originalVM = vm
	}
	if hm, ok := p.horizontalMirror(nil); ok {
		originalHM = hm
	}

	for _, row := range p {
		for c, v := range row {
			row[c] = reverse[v]

			if vm, ok := p.verticalMirror(&originalVM); ok {
				return vm
			}

			if hm, ok := p.horizontalMirror(&originalHM); ok {
				return 100 * hm
			}

			row[c] = v
		}
	}

	panic("mirror not found")
}

type pattern [][]byte

func (p pattern) verticalMirror(ignore *int) (int, bool) {
	width := len(p[0])
nextCol:
	for vm := 1; vm < width; vm++ {
		if ignore != nil && vm == *ignore {
			continue
		}
		dst := min(vm, width-vm)
		for _, row := range p {
			for i := 0; i < dst; i++ {
				if row[vm-i-1] != row[vm+i] {
					continue nextCol
				}
			}
		}

		return vm, true
	}

	return 0, false
}

func (p pattern) horizontalMirror(ignore *int) (int, bool) {
	height := len(p)
nextRow:
	for hm := 1; hm < height; hm++ {
		if ignore != nil && hm == *ignore {
			continue
		}
		dst := min(hm, height-hm)
		for i := 0; i < dst; i++ {
			if !slices.Equal(p[hm-i-1], p[hm+i]) {
				continue nextRow
			}
		}

		return hm, true
	}

	return 0, false
}

const (
	empty = '.'
	rock  = '#'
)
