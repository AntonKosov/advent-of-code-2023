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

func read() []step {
	lines := aoc.ReadAllInput()
	seq := strings.Split(lines[0], ",")
	steps := make([]step, len(seq))
	for i, s := range seq {
		if s[len(s)-1] == '-' {
			steps[i] = step{
				label:  s[:len(s)-1],
				remove: true,
			}
			continue
		}
		parts := strings.Split(s, "=")
		steps[i] = step{
			label:       parts[0],
			focalLength: aoc.StrToInt(parts[1]),
		}
	}

	return steps
}

func process(steps []step) int {
	boxes := loadBoxes(steps)

	totalFocusingPower := 0
	for i, b := range boxes {
		totalFocusingPower += b.focusingPower(i + 1)
	}

	return totalFocusingPower
}

func loadBoxes(steps []step) []box {
	boxes := make([]box, numberOfBoxes)
	for _, s := range steps {
		bi := boxIndex(s.label)
		boxes[bi].applyStep(s)
	}

	return boxes
}

func boxIndex(label string) int {
	idx := 0
	for _, v := range label {
		idx = ((idx + int(v)) * 17) % numberOfBoxes
	}

	return idx
}

type step struct {
	label       string
	remove      bool
	focalLength int
}

type box struct {
	slots []step
}

func (b *box) applyStep(s step) {
	if s.remove {
		b.removeLense(s.label)
		return
	}

	for i, setSlot := range b.slots {
		if setSlot.label == s.label {
			b.slots[i] = s
			return
		}
	}

	b.slots = append(b.slots, s)
}

func (b *box) removeLense(label string) {
	for i, s := range b.slots {
		if s.label != label {
			continue
		}
		copy(b.slots[i:], b.slots[i+1:])
		b.slots = b.slots[:len(b.slots)-1]
		break
	}
}

func (b *box) focusingPower(order int) int {
	power := 0
	for i, s := range b.slots {
		power += order * (i + 1) * s.focalLength
	}

	return power
}

const numberOfBoxes = 256
