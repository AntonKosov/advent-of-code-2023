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

func read() []hand {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	hands := make([]hand, len(lines))
	for i, line := range lines {
		hands[i] = parseHand(line)
	}

	return hands
}

func process(hands []hand) int {
	slices.SortFunc(hands, func(a, b hand) int { return a.cmp(b) })

	winnings := 0
	for i, h := range hands {
		winnings += (i + 1) * h.bid
	}

	return winnings
}

type hand struct {
	cards []byte
	bid   int
	value int
}

func (h1 hand) cmp(h2 hand) int {
	if diff := h2.value - h1.value; diff != 0 {
		return diff
	}

	for i := range h2.cards {
		if diff := labels[h1.cards[i]] - labels[h2.cards[i]]; diff != 0 {
			return diff
		}
	}

	return 0
}

func (h hand) calculateValue() int {
	matchingCards := make(map[byte]int, len(h.cards))
	for _, c := range h.cards {
		matchingCards[c]++
	}

	list := make([]int, 0, len(matchingCards))
	for _, c := range matchingCards {
		list = append(list, c)
	}
	slices.Sort(list)

	switch len(list) {
	case 1: // five of a kind
		return 0
	case 2:
		return list[0] // four of a kind (1) or full house (2)
	case 3:
		return list[1] + 2 // three of a kind (3) or two pair (4)
	case 4: // one pair
		return 5
	case 5: // high card
		return 6
	}

	panic("unexpected length")
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")

	h := hand{
		cards: []byte(parts[0]),
		bid:   aoc.StrToInt(parts[1]),
	}

	h.value = h.calculateValue()

	return h
}

var labels map[byte]int

func init() {
	labelsInOrder := []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	labels = make(map[byte]int, len(labelsInOrder))
	for i, label := range labelsInOrder {
		labels[label] = i
	}
}
