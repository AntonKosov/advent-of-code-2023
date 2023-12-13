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

func read() []row {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	rows := make([]row, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		rows[i] = row{
			cells:      []byte(parts[0]),
			conditions: aoc.StrToInts(parts[1]),
		}
	}

	return rows
}

func process(rows []row) int {
	unfold(rows)

	arrangements := 0
	cache := map[state]int{}
	for _, r := range rows {
		arrangements += countArrangements(r.cells, r.conditions, cache)
		clear(cache)
	}

	return arrangements
}

func countArrangements(cells []byte, conditions []int, cache map[state]int) (count int) {
	state := newState(cells, conditions)
	if c, ok := cache[state]; ok {
		return c
	}

	defer func() { cache[state] = count }()

	for len(cells) > 0 && cells[0] == cellEmpty {
		cells = cells[1:]
	}

	if len(cells) < sum(conditions)+len(conditions)-1 {
		return 0
	}

	if len(conditions) == 0 {
		if slices.Contains(cells, cellDamaged) {
			return 0
		}
		return 1
	}

	if cells[0] == cellUnknown {
		count += countArrangements(cells[1:], conditions, cache)
	}

	damaged := conditions[0]
	if len(cells) > damaged && cells[damaged] == cellDamaged {
		return count
	}

	if slices.Contains(cells[:damaged], cellEmpty) {
		return count
	}

	if len(cells) <= damaged+1 {
		return count + countArrangements(nil, conditions[1:], cache)
	}

	return count + countArrangements(cells[damaged+1:], conditions[1:], cache)
}

func unfold(rows []row) {
	for i := range rows {
		rows[i].unfold()
	}
}

type row struct {
	cells      []byte
	conditions []int
}

func (r *row) unfold() {
	scale := 5
	newCells := make([]byte, 0, len(r.cells)*scale+scale-1)
	newConditions := make([]int, 0, len(r.conditions)*scale)
	for i := 0; i < scale; i++ {
		if i > 0 {
			newCells = append(newCells, '?')
		}
		newCells = append(newCells, r.cells...)
		newConditions = append(newConditions, r.conditions...)
	}
	r.cells = newCells
	r.conditions = newConditions
}

func sum(values []int) int {
	s := 0
	for _, v := range values {
		s += v
	}

	return s
}

type state struct {
	cellsLeft         int
	encodedConditions int
}

func newState(cells []byte, conditions []int) state {
	encodedConditions := 0
	for _, c := range conditions {
		encodedConditions = encodedConditions*100 + c
	}

	return state{cellsLeft: len(cells), encodedConditions: encodedConditions}
}

const (
	cellEmpty   = '.'
	cellDamaged = '#'
	cellUnknown = '?'
)
