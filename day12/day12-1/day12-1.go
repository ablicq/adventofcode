package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ablicq/adventofcode/utils"
	"gonum.org/v1/gonum/stat/combin"
)

const (
	Unknown = '?'
	Broken  = '#'
	Ok      = '.'
)

func ParseLine(line string) (string, []int, bool) {
	record, summary, ok := strings.Cut(line, " ")
	if !ok {
		return "", make([]int, 0), false
	}

	chunksStr := strings.Split(summary, ",")
	chunksInt := make([]int, 0, len(chunksStr))
	for _, s := range chunksStr {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return "", make([]int, 0), false
		}
		chunksInt = append(chunksInt, int(i))
	}

	return record, chunksInt, true
}

func sumSlice(s []int) int {
	sum := 0
	for _, e := range s {
		sum += e
	}
	return sum
}

func GetUnknownIndices(record string) []int {
	idx := make([]int, 0, len(record))
	for i, c := range record {
		if c == Unknown {
			idx = append(idx, i)
		}
	}
	return idx
}

func describe(record []rune) []int {
	description := make([]int, 0)
	combo := 0
	for _, c := range record {
		if c == Broken {
			combo++
		} else if combo != 0 {
			description = append(description, combo)
			combo = 0
		}
	}
	if combo != 0 {
		description = append(description, combo)
	}
	return description
}

func ValidCombinations(line string) int {
	record, summary, ok := ParseLine(line)
	if !ok {
		panic(fmt.Sprintf("Ill-formatted line: %s", line))
	}

	nUnknown := strings.Count(record, "?")
	nBroken := strings.Count(record, "#")
	totalBroken := sumSlice(summary)

	missingBroken := totalBroken - nBroken

	unknownIndices := GetUnknownIndices(record)

	cpt := 0
	for _, combination := range combin.Combinations(nUnknown, missingBroken) {
		testedRecord := []rune(record)
		for _, i := range combination {
			testedRecord[unknownIndices[i]] = Broken
		}
		description := describe(testedRecord)
		if slices.Equal(description, summary) {
			cpt++
		}
	}

	return cpt
}

func main() {
	input := utils.ParseInput("../input.txt")
	sum := 0
	for _, line := range input {
		sum += ValidCombinations(line)
	}

	fmt.Println(sum)
}
