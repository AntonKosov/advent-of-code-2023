package main

import (
	"container/heap"
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	heatMap := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, h := range line {
			row[j] = int(h - '0')
		}

		heatMap[i] = row
	}

	return heatMap
}

func process(heatMap [][]int) int {
	width, height := len(heatMap[0]), len(heatMap)
	target := aoc.NewVector2(width-1, height-1)
	ps := PriorityState{
		{direction: aoc.NewVector2(1, 0)},
		{direction: aoc.NewVector2(0, 1)},
	}
	processedStates := map[processedState]int{}
	for len(ps) > 0 {
		s := heap.Pop(&ps).(state)
		if s.position == target {
			return s.lostHeat
		}
		prState := processedState{
			position:      s.position,
			direction:     s.direction,
			straightCount: s.straightCount,
		}
		if ls, ok := processedStates[prState]; ok && ls <= s.lostHeat {
			continue
		}
		processedStates[prState] = s.lostHeat

		addValid := func(pos, dir aoc.Vector2, straightCount int) {
			if pos.X < 0 || pos.Y < 0 || pos.X >= width || pos.Y >= height {
				return
			}
			heap.Push(&ps, state{
				position:      pos,
				direction:     dir,
				straightCount: straightCount,
				lostHeat:      s.lostHeat + heatMap[pos.Y][pos.X],
			})
		}

		if s.straightCount < 3 {
			addValid(s.position.Add(s.direction), s.direction, s.straightCount+1)
		}

		leftDir, rightDir := s.direction.RotateLeft(), s.direction.RotateRight()
		addValid(s.position.Add(leftDir), leftDir, 1)
		addValid(s.position.Add(rightDir), rightDir, 1)
	}

	panic("path not found")
}

type processedState struct {
	position      aoc.Vector2
	direction     aoc.Vector2
	straightCount int
}

type state struct {
	position      aoc.Vector2
	direction     aoc.Vector2
	straightCount int
	lostHeat      int
}

type PriorityState []state

func (ps PriorityState) Len() int {
	return len(ps)
}

func (ps PriorityState) Less(i, j int) bool {
	psi, psj := ps[i], ps[j]

	if psi.lostHeat == psj.lostHeat {
		return psi.position.X+psi.position.Y > psj.position.X+psj.position.Y
	}

	return psi.lostHeat < psj.lostHeat
}

func (ps PriorityState) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps *PriorityState) Push(x any) {
	*ps = append(*ps, x.(state))
}

func (ps *PriorityState) Pop() any {
	old := *ps
	n := len(old)
	s := old[n-1]
	*ps = old[:n-1]

	return s
}
