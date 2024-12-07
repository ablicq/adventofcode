package main

import (
	"fmt"

	"github.com/ablicq/adventofcode/utils"
)

func diff(seq []int) []int {
	dSeq := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		dSeq[i] = seq[i+1] - seq[i]
	}
	return dSeq
}

func isAllZero(seq []int) bool {
	for _, e := range seq {
		if e != 0 {
			return false
		}
	}
	return true
}

func forecast(numbers []int) int {
	ret := 0
	f := 1
	seq := numbers
	for !isAllZero(seq) {
		ret += f * seq[0]
		f *= -1
		seq = diff(seq)
	}
	return ret
}

func main() {
	input := utils.ParseInput("../input.txt")

	sum := 0
	for _, line := range input {
		numbers := utils.ParseIntList(line)
		sum += forecast(numbers)
	}

	fmt.Println(sum)
}
