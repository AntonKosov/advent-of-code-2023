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
	cards     []byte
	bid       int
	typeValue int
}

func (h1 hand) cmp(h2 hand) int {
	if diff := h2.typeValue - h1.typeValue; diff != 0 {
		return diff
	}

	for i := range h2.cards {
		if diff := labels[h1.cards[i]] - labels[h2.cards[i]]; diff != 0 {
			return diff
		}
	}

	return 0
}

func (h hand) calculateType() int {
	matchingCards := make(map[byte]int, len(h.cards))
	for _, c := range h.cards {
		matchingCards[c]++
	}

	list := make([]int, 0, len(matchingCards))
	for _, c := range matchingCards {
		list = append(list, c)
	}
	slices.Sort(list)

	jokers := matchingCards[joker]
	switch len(list) {
	case 1:
		return handFiveOfAKind
	case 2:
		if list[0] == 1 { // four of a kind
			if jokers > 0 {
				return handFiveOfAKind
			}
			return handFourOfAKind
		}
		// full house
		switch jokers {
		case 0:
			return handFullHouse
		case 1:
			return handFourOfAKind
		}
		return handFiveOfAKind
	case 3:
		v := list[1] + 2 // three of a kind (3) or two pair (4)
		switch jokers {
		case 1:
			if v == 3 {
				return handFourOfAKind
			}
			return handFullHouse
		case 2:
			if v == 3 {
				return handFiveOfAKind
			}
			return handFourOfAKind
		case 3:
			return handFourOfAKind
		}
		return v
	case 4: // one pair
		switch jokers {
		case 0:
			return handOnePair
		case 3:
			return handFiveOfAKind
		}
		return handThreeOfAKind
	case 5: // high card
		return 6 - jokers // high card (6) or one pair (5)
	}

	panic("unexpected value")
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")

	h := hand{
		cards: []byte(parts[0]),
		bid:   aoc.StrToInt(parts[1]),
	}

	h.typeValue = h.calculateType()

	return h
}

const (
	joker = 'J'

	handFiveOfAKind  = 0
	handFourOfAKind  = 1
	handFullHouse    = 2
	handThreeOfAKind = 3
	handTwoPair      = 4
	handOnePair      = 5
	handHighCard     = 6
)

var labels map[byte]int

func init() {
	labelsInOrder := []byte{joker, '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	labels = make(map[byte]int, len(labelsInOrder))
	for i, label := range labelsInOrder {
		labels[label] = i
	}
}
