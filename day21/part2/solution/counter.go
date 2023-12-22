package solution

import (
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func Parse(lines []string) ([][]byte, aoc.Vector2[int]) {
	var start aoc.Vector2[int]
	garden := make([][]byte, len(lines))
	for i, line := range lines {
		startIndex := strings.IndexRune(line, 'S')
		row := []byte(line)
		if startIndex >= 0 {
			start = aoc.NewVector2(startIndex, i)
			row[startIndex] = CellEmpty
		}
		garden[i] = row
	}

	return garden, start
}

func Count(garden [][]byte, start aoc.Vector2[int], steps int) int {
	// * the map is a square (131x131)
	// * the start is in the center (65, 65)
	// * there are no vertical and horizontal obsticles from the start point
	// * the first neighbours will be reached in 66 moves
	// * the initial map will be completely filled in 65+65=130 moves
	// * the edges have no rocks
	n := len(garden)
	if steps <= n/2 {
		return count(garden, start, steps)
	}

	r := radius(steps, n)
	straightMoves := (steps - n/2 - 1) % n
	fullEvenGardens, fullOddGardens := r*r, (r-1)*(r-1)
	if r%2 == 1 {
		fullEvenGardens, fullOddGardens = fullOddGardens, fullEvenGardens
	}

	countFullEvenSteps := count(garden, start, n*2)
	countFullOddSteps := count(garden, start, n*2+1)
	countTopSteps := count(garden, aoc.NewVector2(start.X, n-1), straightMoves)
	countBottomSteps := count(garden, aoc.NewVector2(start.X, 0), straightMoves)
	countLeftSteps := count(garden, aoc.NewVector2(n-1, start.Y), straightMoves)
	countRightSteps := count(garden, aoc.NewVector2(0, start.Y), straightMoves)

	count := fullEvenGardens*countFullEvenSteps + fullOddGardens*countFullOddSteps +
		countTopSteps + countBottomSteps + countLeftSteps + countRightSteps

	top := aoc.NewVector2(0, steps)
	bottom := top.Mul(-1)
	left := aoc.NewVector2(-steps, 0)
	right := left.Mul(-1)
	count += countEdge(garden, steps, top, left, aoc.NewVector2(n-1, n-1))
	count += countEdge(garden, steps, left, bottom, aoc.NewVector2(n-1, 0))
	count += countEdge(garden, steps, bottom, right, aoc.NewVector2(0, 0))
	count += countEdge(garden, steps, right, top, aoc.NewVector2(0, n-1))

	return count
}

func countEdge(garden [][]byte, steps int, from, to, corner aoc.Vector2[int]) int {
	n := len(garden)
	counted := map[aoc.Vector2[int]]bool{}
	cache := map[int]int{}
	total := 0
	offset := to.Sub(from)
	offset.X /= aoc.Abs(offset.X)
	offset.Y /= aoc.Abs(offset.Y)

	for pos := from.Add(offset); pos != to; pos = pos.Add(offset) {
		gardenPos := aoc.NewVector2((aoc.Abs(pos.X)+n/2)/n, (aoc.Abs(pos.Y)+n/2)/n)
		if counted[gardenPos] {
			continue
		}
		counted[gardenPos] = true
		if gardenPos.X == 0 || gardenPos.Y == 0 {
			continue
		}
		firstStep := aoc.NewVector2(n*(gardenPos.X-1)+n/2+1, n*(gardenPos.Y-1)+n/2+1)
		absPos := aoc.NewVector2(aoc.Abs(pos.X), aoc.Abs(pos.Y))
		gardenSteps := firstStep.ManhattanDst(absPos)
		filled, ok := cache[gardenSteps]
		if !ok {
			filled = count(garden, corner, gardenSteps)
			cache[gardenSteps] = filled
		}
		total += filled
	}

	return total
}

func radius(steps, n int) int {
	if steps < n-1 {
		return 0
	}

	if steps == n-1 {
		return 1
	}

	return (steps + 1) / n
}

func count(garden [][]byte, start aoc.Vector2[int], steps int) int {
	if steps < 0 {
		return 0
	}
	dirs := []aoc.Vector2[int]{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	currentPositions := map[aoc.Vector2[int]]struct{}{start: {}}
	nextPositions := map[aoc.Vector2[int]]struct{}{}
	width, height := len(garden[0]), len(garden)

	for i := 0; i < steps; i++ {
		for p := range currentPositions {
			for _, dir := range dirs {
				np := p.Add(dir)
				if np.X < 0 || np.Y < 0 || np.X >= width || np.Y >= height {
					continue
				}
				if garden[np.Y][np.X] == CellEmpty {
					nextPositions[np] = struct{}{}
				}
			}
		}
		currentPositions, nextPositions = nextPositions, currentPositions
		clear(nextPositions)
	}

	return len(currentPositions)
}

const (
	CellEmpty = '.'
	CellRock  = '#'
)
