package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

func ParseLine(line string) (string, []int, bool) {
	symbols, blocks, ok := strings.Cut(line, " ")
	if !ok {
		return "", make([]int, 0), false
	}

	blocksStr := strings.Split(blocks, ",")
	blocksInt := make([]int, 0, len(blocksStr))
	for _, s := range blocksStr {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return "", make([]int, 0), false
		}
		blocksInt = append(blocksInt, int(i))
	}

	return symbols, blocksInt, true
}

func sumSlice(s []int) int {
	sum := 0
	for _, e := range s {
		sum += e
	}
	return sum
}

func isCompatible(testedSlice string, testedMotif string) bool {
	for i := range testedMotif {
		if testedSlice[i] != '?' && testedSlice[i] != testedMotif[i] {
			return false
		}
	}
	return true
}

func firstOccurrence(symbols string, c uint8) int {
	for i := 0; i < len(symbols); i++ {
		if symbols[i] == c {
			return i
		}
	}
	return len(symbols)
}

func solve(symbols string, blocks []int) int {
	if len(blocks) == 0 {
		if strings.Count(symbols, "#") == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(symbols) == 0 {
		return 0
	}

	firstBlock := blocks[0]
	testedMotif := fmt.Sprintf(".%s.", strings.Repeat("#", firstBlock))
	symbolsExtended := fmt.Sprintf(".%s.", symbols)
	lastTestedPosition := min(
		firstOccurrence(symbolsExtended, '#')+firstBlock+1,
		len(symbolsExtended)-(sumSlice(blocks[1:])+len(blocks[1:])),
	)

	sum := 0

	for i := 0; i < lastTestedPosition-len(testedMotif)+1; i++ {
		testedSlice := symbolsExtended[i : i+len(testedMotif)]
		if isCompatible(testedSlice, testedMotif) {
			next_symbols := symbols[min(i+firstBlock+1, len(symbols)):]
			next_blocks := blocks[1:]
			sum += solve(next_symbols, next_blocks)
		}
	}

	return sum
}

func main() {
	input := utils.ParseInput("../input.txt")
	sum := 0
	for _, line := range input {
		symbols, blocks, _ := ParseLine(line)
		sum += solve(symbols, blocks)
	}

	fmt.Println(sum)
}
