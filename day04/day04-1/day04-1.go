package main

import (
	"fmt"
	"math"
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

func process(line string) int {
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

	if commonCount == 0 {
		return 0
	}
	return int(math.Pow(2., float64(commonCount)-1))
}

func main() {
	parsedInput := utils.ParseInput("../input.txt")

	sum := 0

	for _, line := range parsedInput {
		sum += process(line)
	}

	fmt.Println(sum)
}
