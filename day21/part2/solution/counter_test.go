package solution_test

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
	"github.com/AntonKosov/advent-of-code-2023/day21/part2/solution"
)

func runTest(t *testing.T, input []string, steps, expectedSteps int) {
	if steps%2 == 0 {
		t.Fatalf("even number of steps: %v", steps)
	}
	if len(input)%2 == 0 {
		t.Fatal("even number of lines")
	}
	w := len(input[0])
	if w%2 == 0 {
		t.Fatal("even number of columns")
	}
	for _, line := range input {
		if len(line) != w {
			t.Fatal("invalid input")
		}
	}

	garden, start := solution.Parse(input)
	actualSteps := solution.Count(garden, start, steps)

	if actualSteps != expectedSteps {
		t.Logf("expected steps: %v, actual steps: %v\n", expectedSteps, actualSteps)
		t.Fail()
	}
}

func TestEmpty3x3Steps5(t *testing.T) {
	input := []string{
		"...",
		".S.",
		"...",
	}
	runTest(t, input, 5, 36)
}

func TestEmpty3x3Steps7(t *testing.T) {
	input := []string{
		"...",
		".S.",
		"...",
	}
	runTest(t, input, 7, 8*1+4*4+4*4+1*4+4*5)
}

func TestEmpty3x3Steps13(t *testing.T) {
	input := []string{
		"...",
		".S.",
		"...",
	}
	runTest(t, input, 13, 4*4+4*4+1*4+4*5+12*5+16*4+4*4)
}

func TestEmpty5x5Steps5(t *testing.T) {
	// no diagonals
	input := []string{
		".....",
		".....",
		"..S..",
		".....",
		".....",
	}
	runTest(t, input, 5, 36)
}

func TestEmpty5x5Steps7(t *testing.T) {
	input := []string{
		".....",
		".....",
		"..S..",
		".....",
		".....",
	}
	runTest(t, input, 7, 12+4*11+4*2)
}

func Test1Rock5x5Steps5(t *testing.T) {
	input := []string{
		".....",
		".#...",
		"..S..",
		".....",
		".....",
	}
	runTest(t, input, 5, 36-2)
}

func Test2Rocks5x5Steps5(t *testing.T) {
	input := []string{
		".....",
		".#...",
		"..S..",
		".#...",
		".....",
	}
	runTest(t, input, 5, 36-4)
}

func Test1Rock5x5Steps7(t *testing.T) {
	input := []string{
		".....",
		".#...",
		"..S..",
		".....",
		".....",
	}
	runTest(t, input, 7, 60)
}

func Test2Rocks5x5Steps7(t *testing.T) {
	input := []string{
		".....",
		".#...",
		"..S..",
		".#...",
		".....",
	}
	runTest(t, input, 7, 56)
}

func TestEmpty5x5Steps11(t *testing.T) {
	input := []string{
		".....",
		".....",
		"..S..",
		".....",
		".....",
	}
	runTest(t, input, 11, 12+4*13+4*2+4*8+4*10)
}

func Test2Rocks5x5Steps11(t *testing.T) {
	input := []string{
		".....",
		".#...",
		"..S..",
		".#...",
		".....",
	}
	runTest(t, input, 11, 12+4*13+4*2+4*8+4*10-8)
}

func TestPart1(t *testing.T) {
	input := aoc.ReadAllInputFromFile("../input.txt")
	input = input[:len(input)-1]
	runTest(t, input, 63, 3589)
}

func TestPart2(t *testing.T) {
	input := aoc.ReadAllInputFromFile("../input.txt")
	input = input[:len(input)-1]
	runTest(t, input, 26501365, 621289922886149)
}
