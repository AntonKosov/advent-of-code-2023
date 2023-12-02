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
	bag := set{red: 12, green: 13, blue: 14}
	sum := 0
	for _, g := range games {
		impossible := slices.ContainsFunc(g.sets, func(s set) bool {
			return s.red > bag.red || s.green > bag.green || s.blue > bag.blue
		})
		if !impossible {
			sum += g.id
		}
	}

	return sum
}

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}
