package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

type Range struct {
	srcStart int
	dstStart int
	length   int
}

type RangeMap []Range

func (m RangeMap) get(key int) int {
	for _, r := range m {
		if key >= r.srcStart && key < r.srcStart+r.length {
			return r.dstStart + key - r.srcStart
		}
	}
	return key
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func extractSeeds(s string) []int {
	ss := strings.Split(s, " ")
	ret := make([]int, len(ss)-1)
	for i := 1; i < len(ss); i++ {
		ret[i-1] = atoi(ss[i])
	}
	return ret
}

func getBlankLines(input []string) []int {
	ret := make([]int, 8)
	idx := 0
	for i, line := range input {
		if line == "" {
			ret[idx] = i
			idx++
		}
	}
	ret[idx] = len(input)
	return ret
}

func extractRangeMap(input []string) RangeMap {
	rangeMap := make([]Range, len(input))
	for i, line := range input {
		s := strings.Split(line, " ")
		rangeMap[i] = Range{
			atoi(s[1]),
			atoi(s[0]),
			atoi(s[2]),
		}
	}
	return rangeMap
}

func extractRangeMaps(input []string) []RangeMap {
	blankLines := getBlankLines(input)
	ret := make([]RangeMap, 7)

	for i := 0; i < len(blankLines)-1; i++ {
		ret[i] = extractRangeMap(input[blankLines[i]+2 : blankLines[i+1]])
	}

	return ret
}

func applyMaps(rangeMaps []RangeMap, seeds []int) []int {
	locations := make([]int, len(seeds))
	for i, seed := range seeds {
		v := seed
		for _, m := range rangeMaps {
			v = m.get(v)
		}
		locations[i] = v
	}
	return locations
}

func main() {
	input := utils.ParseInput("../input.txt")
	seeds := extractSeeds(input[0])
	rangeMaps := extractRangeMaps(input)
	locations := applyMaps(rangeMaps, seeds)

	fmt.Println(slices.Min(locations))
}
