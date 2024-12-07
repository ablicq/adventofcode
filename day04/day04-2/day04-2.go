package main

import (
	"fmt"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

func getCardNumbers(cardString string) map[string]bool {
	uniqueNumbers := make(map[string]bool, 0)
	for _, num := range strings.Split(cardString, " ") {
		if len(num) > 0 {
			uniqueNumbers[num] = true
		}
	}
	return uniqueNumbers
}

func getCardMatches(line string) int {
	_, card, f1 := strings.Cut(line, ": ")
	if !f1 {
		panic("Invalid input format")
	}

	mine, winning, f2 := strings.Cut(card, " | ")
	if !f2 {
		panic("Invalid input format")
	}
	mineSet, winningSet := getCardNumbers(mine), getCardNumbers(winning)

	commonCount := 0

	for k := range mineSet {
		if winningSet[k] {
			commonCount++
		}
	}

	return commonCount
}

func main() {
	input := utils.ParseInput("../input.txt")

	nCards := len(input)

	copies := make([]int, nCards)

	for cardIdx, card := range input {
		nMatches := getCardMatches(card)
		for i := cardIdx + 1; i <= cardIdx+nMatches; i++ {
			copies[i] += copies[cardIdx] + 1
		}
	}

	sum := nCards
	for _, c := range copies {
		sum += c
	}

	fmt.Println(sum)
}
