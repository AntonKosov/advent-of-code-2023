package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []game {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	games := make([]game, len(lines))
	for i, line := range lines {
		lineParts := strings.Split(line, ": ")
		games[i].id = aoc.StrToInts(lineParts[0])[0]
		setsLine := lineParts[1]
		sets := strings.Split(setsLine, "; ")
		for _, setLine := range sets {
			var s set
			cubes := strings.Split(setLine, ", ")
			for _, cube := range cubes {
				parts := strings.Split(cube, " ")
				count := aoc.StrToInt(parts[0])
				switch color := parts[1]; color {
				case "red":
					s.red = count
				case "green":
					s.green = count
				case "blue":
					s.blue = count
				default:
					panic("unexpected color: " + color)
				}
			}
			games[i].sets = append(games[i].sets, s)
		}
	}

	return games
}

func process(games []game) int {
	sum := 0
	for _, g := range games {
		sum += g.minPossible().power()
	}

	return sum
}

type set struct {
	red   int
	green int
	blue  int
}

func (s set) power() int {
	return s.red * s.green * s.blue
}

type game struct {
	id   int
	sets []set
}

func (g game) minPossible() set {
	var m set
	for _, s := range g.sets {
		m.red = max(m.red, s.red)
		m.green = max(m.green, s.green)
		m.blue = max(m.blue, s.blue)
	}

	return m
}
