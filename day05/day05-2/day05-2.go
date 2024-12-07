package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

type Range struct {
	start  int
	length int
}

type RangeMapEntry struct {
	srcStart int
	dstStart int
	length   int
}

type RangeMap []RangeMapEntry

func (m RangeMap) get(key int) int {
	for _, r := range m {
		if key >= r.srcStart && key < r.srcStart+r.length {
			return r.dstStart + key - r.srcStart
		}
	}
	return key
}

func intersect(r Range, rme RangeMapEntry) (Range, bool) {
	if r.start < rme.srcStart+rme.length && rme.srcStart < r.start+r.length {
		start := max(r.start, rme.srcStart)
		end := min(r.start+r.length, rme.srcStart+rme.length)
		return Range{start, end - start}, true
	}
	return Range{-1, 0}, false
}

func getHoles(fullRange Range, parts []Range) []Range {
	slices.SortFunc(parts, func(r1, r2 Range) int { return r1.start - r2.start })
	currentStart := fullRange.start
	holes := make([]Range, 0)

	for _, r := range parts {
		if currentStart < r.start {
			holes = append(holes, Range{currentStart, r.start - currentStart})
		}
		currentStart = r.start + r.length
	}
	if currentStart < fullRange.start+fullRange.length {
		holes = append(holes, Range{currentStart, fullRange.start + fullRange.length - currentStart})
	}

	return holes
}

func (m RangeMap) applyRange(keyRange Range) []Range {
	srcRanges := make([]Range, 0)
	for _, r := range m {
		if inter, ok := intersect(keyRange, r); ok {
			srcRanges = append(srcRanges, inter)
		}
	}

	holes := getHoles(keyRange, srcRanges)

	dstRanges := make([]Range, len(srcRanges))
	for i, r := range srcRanges {
		dstRanges[i] = Range{m.get(r.start), r.length}
	}

	return append(dstRanges, holes...)
}

func (m RangeMap) applyRanges(keyRanges []Range) []Range {
	valueRanges := make([]Range, 0)
	for _, r := range keyRanges {
		valueRanges = append(valueRanges, m.applyRange(r)...)
	}
	return valueRanges
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func extractSeedRanges(s string) []Range {
	ss := strings.Split(s, " ")
	ret := make([]Range, len(ss)/2)
	for i := 1; i < len(ss); i += 2 {
		ret[i/2] = Range{
			atoi(ss[i]),
			atoi(ss[i+1]),
		}
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
	rangeMap := make([]RangeMapEntry, len(input))
	for i, line := range input {
		s := strings.Split(line, " ")
		rangeMap[i] = RangeMapEntry{
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

func solve(rangeMaps []RangeMap, seeds []Range) int {
	ranges := seeds
	for _, rm := range rangeMaps {
		ranges = rm.applyRanges(ranges)
	}

	return slices.MinFunc(ranges, func(r1, r2 Range) int { return r1.start - r2.start }).start
}

func main() {
	input := utils.ParseInput("../input.txt")
	seedRanges := extractSeedRanges(input[0])
	rangeMaps := extractRangeMaps(input)

	fmt.Println(solve(rangeMaps, seedRanges))
}
