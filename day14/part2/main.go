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

func read() platform {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	plat := make(platform, len(lines))
	for i, line := range lines {
		plat[i] = []byte(line)
	}

	return plat
}

func process(plat platform) int {
	totalCycles := 1_000_000_000
	var platforms []platform
	for i := 0; i < totalCycles; i++ {
		platforms = append(platforms, plat.clone())
		plat.runCycle()
		for j, prevPlatform := range platforms {
			if plat.equals(prevPlatform) {
				cycleLen := i - j + 1
				idx := (totalCycles-j)%cycleLen + j
				return platforms[idx].totalLoad()
			}
		}
	}

	return plat.totalLoad()
}

type platform [][]byte

func (p platform) equals(p2 platform) bool {
	for r := range p {
		if !slices.Equal(p[r], p2[r]) {
			return false
		}
	}

	return true
}

func (p platform) clone() platform {
	clone := make(platform, len(p))
	for i, row := range p {
		cloneRow := make([]byte, len(row))
		copy(cloneRow, row)
		clone[i] = cloneRow
	}

	return clone
}

func (p platform) runCycle() {
	p.tiltNorth()
	p.tiltWest()
	p.tiltSouth()
	p.tiltEast()
}

func (p platform) tiltNorth() {
	for c := 0; c < len(p[0]); c++ {
		emptyRow := -1
		for r := range p {
			switch p[r][c] {
			case empty:
				if emptyRow < 0 {
					emptyRow = r
				}
			case square:
				emptyRow = -1
			case round:
				if emptyRow >= 0 {
					p[emptyRow][c] = round
					p[r][c] = empty
					emptyRow++
				}
			}
		}
	}
}

func (p platform) tiltSouth() {
	for c := 0; c < len(p[0]); c++ {
		emptyRow := -1
		for r := len(p) - 1; r >= 0; r-- {
			switch p[r][c] {
			case empty:
				if emptyRow < 0 {
					emptyRow = r
				}
			case square:
				emptyRow = -1
			case round:
				if emptyRow >= 0 {
					p[emptyRow][c] = round
					p[r][c] = empty
					emptyRow--
				}
			}
		}
	}
}

func (p platform) tiltWest() {
	for _, row := range p {
		emptyCol := -1
		for c, v := range row {
			switch v {
			case empty:
				if emptyCol < 0 {
					emptyCol = c
				}
			case square:
				emptyCol = -1
			case round:
				if emptyCol >= 0 {
					row[emptyCol] = round
					row[c] = empty
					emptyCol++
				}
			}
		}
	}
}

func (p platform) tiltEast() {
	for _, row := range p {
		emptyCol := -1
		for c := len(row) - 1; c >= 0; c-- {
			switch row[c] {
			case empty:
				if emptyCol < 0 {
					emptyCol = c
				}
			case square:
				emptyCol = -1
			case round:
				if emptyCol >= 0 {
					row[emptyCol] = round
					row[c] = empty
					emptyCol--
				}
			}
		}
	}
}

func (p platform) totalLoad() int {
	load := 0
	for r, row := range p {
		countRound := 0
		for _, v := range row {
			if v == round {
				countRound++
			}
		}
		load += countRound * (len(p) - r)
	}

	return load
}

const (
	empty  = '.'
	round  = 'O'
	square = '#'
)
