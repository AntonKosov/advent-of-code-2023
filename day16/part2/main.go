package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]byte {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	device := make([][]byte, len(lines))
	for i, line := range lines {
		device[i] = []byte(line)
	}

	return device
}

func process(device [][]byte) int {
	width, height := len(device[0]), len(device)
	maxEnergizedCells := 0

	for x := 0; x < width; x++ {
		maxEnergizedCells = max(
			maxEnergizedCells,
			energizedCells(device, x, 0, dirDown),
			energizedCells(device, x, height-1, dirUp),
		)
	}

	for y := 0; y < height; y++ {
		maxEnergizedCells = max(
			maxEnergizedCells,
			energizedCells(device, 0, y, dirRight),
			energizedCells(device, width-1, y, dirLeft),
		)
	}

	return maxEnergizedCells
}

func energizedCells(device [][]byte, x, y int, dir lightDir) int {
	lightMap := makeLightMap(len(device[0]), len(device))
	calculateBeam(device, lightMap, aoc.NewVector2(x, y), dir)

	energized := 0
	for _, row := range lightMap {
		for _, v := range row {
			if v != 0 {
				energized++
			}
		}
	}

	return energized
}

func calculateBeam(device [][]byte, lightMap [][]lightDir, pos aoc.Vector2[int], ld lightDir) {
	if pos.Y < 0 || pos.X < 0 || pos.Y >= len(device) || pos.X >= len(device[0]) {
		return
	}

	if lightMap[pos.Y][pos.X]&ld != 0 {
		return
	}

	lightMap[pos.Y][pos.X] |= ld

	deviceCell := rune(device[pos.Y][pos.X])
	if deviceCell == '.' {
		calculateBeam(device, lightMap, pos.Add(offsets[ld]), ld)
		return
	}

	ref := reflection{ld: ld, mirror: deviceCell}
	if refDir, ok := reflections[ref]; ok {
		calculateBeam(device, lightMap, pos.Add(offsets[refDir]), refDir)
		return
	}

	if refDirs, ok := splits[ref]; ok {
		for _, refDir := range refDirs {
			calculateBeam(device, lightMap, pos.Add(offsets[refDir]), refDir)
		}
		return
	}

	panic("unexpected state")
}

func makeLightMap(w, h int) [][]lightDir {
	m := make([][]lightDir, h)
	for i := 0; i < len(m); i++ {
		m[i] = make([]lightDir, w)
	}

	return m
}

type lightDir byte

const (
	dirUp lightDir = 1 << iota
	dirDown
	dirLeft
	dirRight
)

var offsets map[lightDir]aoc.Vector2[int]

type reflection struct {
	ld     lightDir
	mirror rune
}

var reflections map[reflection]lightDir

var splits map[reflection][]lightDir

func init() {
	offsets = map[lightDir]aoc.Vector2[int]{
		dirUp:    aoc.NewVector2(0, -1),
		dirDown:  aoc.NewVector2(0, 1),
		dirLeft:  aoc.NewVector2(-1, 0),
		dirRight: aoc.NewVector2(1, 0),
	}

	reflections = map[reflection]lightDir{
		{ld: dirUp, mirror: '\\'}:    dirLeft,
		{ld: dirDown, mirror: '\\'}:  dirRight,
		{ld: dirRight, mirror: '\\'}: dirDown,
		{ld: dirLeft, mirror: '\\'}:  dirUp,
		{ld: dirUp, mirror: '/'}:     dirRight,
		{ld: dirDown, mirror: '/'}:   dirLeft,
		{ld: dirRight, mirror: '/'}:  dirUp,
		{ld: dirLeft, mirror: '/'}:   dirDown,
	}

	splits = map[reflection][]lightDir{
		{ld: dirUp, mirror: '-'}:    {dirLeft, dirRight},
		{ld: dirDown, mirror: '-'}:  {dirLeft, dirRight},
		{ld: dirRight, mirror: '-'}: {dirRight},
		{ld: dirLeft, mirror: '-'}:  {dirLeft},
		{ld: dirUp, mirror: '|'}:    {dirUp},
		{ld: dirDown, mirror: '|'}:  {dirDown},
		{ld: dirRight, mirror: '|'}: {dirUp, dirDown},
		{ld: dirLeft, mirror: '|'}:  {dirUp, dirDown},
	}
}
