package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []instruction {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	instructions := make([]instruction, len(lines))
	dirs := map[byte]byte{'0': 'R', '1': 'D', '2': 'L', '3': 'U'}
	for i, line := range lines {
		parts := strings.Split(line, " (#")
		part1 := parts[1]
		meters, err := strconv.ParseInt(part1[:5], 16, 64)
		aoc.Must(err)
		instructions[i] = instruction{
			dir:    dirs[part1[5]],
			meters: int(meters),
		}
	}

	return instructions
}

func process(instructions []instruction) int {
	corners := dig(instructions)
	compressedMap, verticalValues, horizontalValues := compress(corners)
	fillCompressedMap(corners, compressedMap, verticalValues, horizontalValues)
	subEdges(compressedMap)

	return sumArea(compressedMap, verticalValues, horizontalValues)
}

func sumArea(compressedMap [][]int, verticalValues, horizontalValues []int) int {
	sum := 0
	for y, row := range compressedMap {
		h := horizontalValues[y+1] - horizontalValues[y]
		for x, v := range row {
			if v > 1 {
				continue
			}
			w := verticalValues[x+1] - verticalValues[x]
			sum += w * h
		}
	}

	return sum
}

func subEdges(compressedMap [][]int) {
	for y := range compressedMap {
		sub(compressedMap, aoc.NewVector2(0, y))
		sub(compressedMap, aoc.NewVector2(len(compressedMap[y])-1, y))
	}
	for x := 0; x < len(compressedMap[0]); x++ {
		sub(compressedMap, aoc.NewVector2(x, 0))
		sub(compressedMap, aoc.NewVector2(x, len(compressedMap)-1))
	}
}

func sub(compressedMap [][]int, pos aoc.Vector2[int]) {
	if pos.X < 0 || pos.Y < 0 || pos.X >= len(compressedMap[0]) || pos.Y >= len(compressedMap) {
		return
	}

	if compressedMap[pos.Y][pos.X] != 0 {
		return
	}

	compressedMap[pos.Y][pos.X] = 2

	for _, dir := range dirs {
		sub(compressedMap, pos.Add(dir))
	}
}

func fillCompressedMap(corners []aoc.Vector2[int], compressedMap [][]int, verticalValues, horizontalValues []int) {
	currentPos := aoc.NewVector2(0, 0)
	if vv, b := slices.BinarySearch(verticalValues, 0); !b {
		panic("origin not found")
	} else {
		currentPos.X = vv
	}
	if hv, b := slices.BinarySearch(horizontalValues, 0); !b {
		panic("origin not found")
	} else {
		currentPos.Y = hv
	}

	for i := 0; i < len(corners)-1; i++ {
		cornerB, cornerE := corners[i], corners[i+1]
		dir := cornerE.Sub(cornerB).Norm()
		for {
			currentPos = currentPos.Add(dir)
			if dir.X == 1 {
				compressedMap[currentPos.Y][currentPos.X-1] = 1
			} else if dir.X == -1 {
				compressedMap[currentPos.Y-1][currentPos.X] = 1
			} else if dir.Y == 1 {
				compressedMap[currentPos.Y-1][currentPos.X-1] = 1
			} else {
				compressedMap[currentPos.Y][currentPos.X] = 1
			}
			if cornerE.X == verticalValues[currentPos.X] && cornerE.Y == horizontalValues[currentPos.Y] {
				break
			}
		}
	}
}

func compress(corners []aoc.Vector2[int]) (compressedMap [][]int, verticalValues, horizontalValues []int) {
	vm, hm := map[int]struct{}{}, map[int]struct{}{}
	for _, corner := range corners {
		vm[corner.X] = struct{}{}
		hm[corner.Y] = struct{}{}
	}
	verticalValues = convertMapToSortedSlice(vm)
	horizontalValues = convertMapToSortedSlice(hm)
	compressedMap = make([][]int, len(horizontalValues)-1)
	for i := range compressedMap {
		compressedMap[i] = make([]int, len(verticalValues)-1)
	}

	return compressedMap, verticalValues, horizontalValues
}

func convertMapToSortedSlice(m map[int]struct{}) []int {
	s := make([]int, 0, len(m))
	for k := range m {
		s = append(s, k)
	}

	slices.Sort(s)

	return s
}

func dig(instructions []instruction) (corners []aoc.Vector2[int]) {
	currentPos := aoc.NewVector2(0, 0)
	corners = make([]aoc.Vector2[int], 0, len(instructions)+1)
	corners = append(corners, currentPos)
	currentDirLetter := instructions[0].dir
	var currentDir aoc.Vector2[int]
	switch currentDirLetter {
	case 'U':
		currentDir = aoc.NewVector2(0, -1)
	case 'D':
		currentDir = aoc.NewVector2(0, 1)
	case 'R':
		currentDir = aoc.NewVector2(1, 0)
	case 'L':
		currentDir = aoc.NewVector2(-1, 0)
	default:
		panic("unexpected direction")
	}
	prevDir := currentDir.RotateLeft()
	first := true
	for _, instr := range instructions {
		rot := rotations[currentDirLetter][instr.dir]
		if !first && rot == 0 {
			panic("no rotation")
		}
		leftRot := rot < 0
		if !first {
			prevDir = currentDir
		}
		offset := 0
		if leftRot {
			if !first {
				currentDir = currentDir.RotateLeft()
				offset--
			}
		} else {
			if !first {
				currentDir = currentDir.RotateRight()
				currentPos = currentPos.Add(prevDir)
				corners[len(corners)-1] = currentPos
			}
		}
		currentDirLetter = instr.dir
		currentPos = currentPos.Add(currentDir.Mul(instr.meters + offset))
		corners = append(corners, currentPos)
		first = false
	}

	currentPos = currentPos.Add(currentDir)
	corners[len(corners)-1] = currentPos

	return corners
}

type instruction struct {
	dir    byte
	meters int
}

var dirs []aoc.Vector2[int]
var rotations map[byte]map[byte]int

func init() {
	left, right := -1, 1
	rotations = map[byte]map[byte]int{
		'U': {'L': left, 'R': right},
		'D': {'L': right, 'R': left},
		'R': {'U': left, 'D': right},
		'L': {'U': right, 'D': left},
	}
	dirs = []aoc.Vector2[int]{
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
	}
}
