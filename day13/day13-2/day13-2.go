package main

import (
	"fmt"

	"github.com/ablicq/adventofcode/utils"
)

func transpose(block []string) []string {
	newBlock := make([]string, 0, len(block[0]))
	for i := range block[0] {
		newRow := make([]uint8, 0, len(block))
		for j := range block {
			newRow = append(newRow, block[j][i])
		}
		newBlock = append(newBlock, string(newRow))
	}
	return newBlock
}

func nDifferences(block []string, i int) int {
	nDiff := 0
	for delta := 0; i-delta >= 0 && i+delta+1 < len(block); delta++ {
		l1, l2 := block[i-delta], block[i+delta+1]
		for c := range l1 {
			if l1[c] != l2[c] {
				nDiff++
			}
		}
	}
	return nDiff
}

func findReflectionWithSmudge(block []string) int {
	for i := 0; i < len(block)-1; i++ {
		if nDifferences(block, i) == 1 {
			return i
		}
	}
	return -1
}

func processBlock(block []string) int {
	reflectIdx := findReflectionWithSmudge(block)

	if reflectIdx >= 0 {
		return 100 * (reflectIdx + 1)
	}

	return findReflectionWithSmudge(transpose(block)) + 1
}

func splitBlocks(input []string) [][]string {
	blocks := make([][]string, 0)
	newBlock := make([]string, 0)
	for _, line := range input {
		if line == "" {
			blocks = append(blocks, newBlock)
			newBlock = make([]string, 0)
		} else {
			newBlock = append(newBlock, line)
		}
	}
	blocks = append(blocks, newBlock)
	return blocks
}

func main() {
	input := utils.ParseInput("../input.txt")

	blocks := splitBlocks(input)

	sum := 0

	for _, block := range blocks {
		sum += processBlock(block)
	}

	fmt.Println(sum)
}
