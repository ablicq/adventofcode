package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

type HandType int

const (
	High HandType = iota + 1
	Pair
	TwoPair
	ThreeOfAKind
	Full
	FourOfAKind
	FiveOfAKind
)

type Hand [5]int

type Score [6]int

type ScoredHand struct {
	hand  Hand
	bid   int
	score []int
}

var heightMap = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 0,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func (h Hand) Type() HandType {
	// construct card histogram
	hist := make(map[int]int, 5)
	for _, card := range h {
		_, ok := hist[card]
		if ok {
			hist[card] += 1
		} else {
			hist[card] = 1
		}
	}

	// dispatch the jokers
	jokers, ok := hist[0]
	wcKey, wcValue := 0, 0
	if ok && jokers != 5 {
		for k, v := range hist {
			if (k != 0) && (v > wcValue || (v == wcValue && k > wcKey)) {
				wcKey, wcValue = k, v
			}
		}
	}
	if wcKey != 0 {
		hist[wcKey] += jokers
		delete(hist, 0)
	}

	// infer the card type from the histogram values
	switch len(hist) {
	case 5:
		return High
	case 4:
		return Pair
	case 3:
		for _, v := range hist {
			if v == 2 {
				return TwoPair
			} else if v == 3 {
				return ThreeOfAKind
			}
		}
	case 2:
		for _, v := range hist {
			if v == 4 || v == 1 {
				return FourOfAKind
			} else {
				return Full
			}
		}
	case 1:
		return FiveOfAKind
	}

	return 0
}

func (h Hand) Score() []int {
	ret := make([]int, 6)
	ret[0] = int(h.Type())
	for i, card := range h {
		ret[i+1] = card
	}
	return ret
}

func getScoredhand(line string) ScoredHand {
	var hand Hand
	parts := strings.Split(line, " ")
	for i, r := range parts[0] {
		hand[i] = heightMap[r]
	}
	bid := utils.ParseInt(parts[1])
	return ScoredHand{hand, bid, hand.Score()}
}

func main() {
	input := utils.ParseInput("../input.txt")
	scoredHands := make([]ScoredHand, len(input))
	for i, l := range input {
		scoredHands[i] = getScoredhand(l)
	}

	slices.SortFunc(scoredHands, func(s1, s2 ScoredHand) int {
		for i := range s1.score {
			d := s1.score[i] - s2.score[i]
			if d == 0 {
				continue
			}
			return d
		}
		return 0
	})

	sum := 0
	for i, h := range scoredHands {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)
}
