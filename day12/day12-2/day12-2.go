package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

type SpringProblem struct {
	symbols string
	blocks  []int
}

func convert(s []int) string {
	str := make([]string, 0)
	for _, i := range s {
		str = append(str, strconv.FormatInt(int64(i), 10))
	}
	return strings.Join(str, ",")
}

func (p SpringProblem) String() string {
	return fmt.Sprintf("%s %s", p.symbols, convert(p.blocks))
}

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

func solve(problem SpringProblem, resMap map[string]int) int {
	res, ok := resMap[problem.String()]
	if ok {
		return res
	}
	if len(problem.blocks) == 0 {
		if strings.Count(problem.symbols, "#") == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(problem.symbols) == 0 {
		return 0
	}

	firstBlock := problem.blocks[0]
	testedMotif := fmt.Sprintf(".%s.", strings.Repeat("#", firstBlock))
	symbolsExtended := fmt.Sprintf(".%s.", problem.symbols)
	lastTestedPosition := min(
		firstOccurrence(symbolsExtended, '#')+firstBlock+1,
		len(symbolsExtended)-(sumSlice(problem.blocks[1:])+len(problem.blocks[1:])),
	)

	sum := 0

	for i := 0; i < lastTestedPosition-len(testedMotif)+1; i++ {
		testedSlice := symbolsExtended[i : i+len(testedMotif)]
		if isCompatible(testedSlice, testedMotif) {
			next_symbols := problem.symbols[min(i+firstBlock+1, len(problem.symbols)):]
			next_blocks := problem.blocks[1:]
			next_problem := SpringProblem{next_symbols, next_blocks}
			res := solve(next_problem, resMap)
			resMap[next_problem.String()] = res
			sum += res
		}
	}

	return sum
}

func expandSymbols(symbols string) string {
	return strings.Join([]string{symbols, symbols, symbols, symbols, symbols}, "?")
}

func expandBlocks(blocks []int) []int {
	ret := make([]int, 0, 5*len(blocks))
	for i := 0; i < 5; i++ {
		ret = append(ret, blocks...)
	}
	return ret
}

func main() {
	input := utils.ParseInput("../input.txt")
	sum := 0
	for _, line := range input {
		symbols, blocks, _ := ParseLine(line)
		expandedSymbols, expandedBlocks := expandSymbols(symbols), expandBlocks(blocks)
		problem := SpringProblem{expandedSymbols, expandedBlocks}
		v := solve(problem, make(map[string]int, 0))
		sum += v
	}

	fmt.Println(sum)
}
