package main

import (
	"fmt"
	"strings"

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

func tilt_one_column(col string) string {
	tilted_col := make([]rune, 0, len(col))
	for idx, tile := range col {
		switch tile {
		case 'O':
			tilted_col = append(tilted_col, 'O')
		case '#':
			for range idx - len(tilted_col) {
				tilted_col = append(tilted_col, '.')
			}
			tilted_col = append(tilted_col, '#')
		default:
			continue
		}
	}
	for range len(col) - len(tilted_col) {
		tilted_col = append(tilted_col, '.')
	}

	return string(tilted_col)
}

func compute_weight(grid []string) int {
	weight := 0
	for i, row := range grid {
		weight += (len(grid) - i) * strings.Count(row, "O")
	}
	return weight
}

func main() {
	grid := utils.ParseInput("../input.txt")
	transposed_tilted_grid := make([]string, 0, len(grid[0]))
	for _, col := range transpose(grid) {
		tilted_col := tilt_one_column(col)
		transposed_tilted_grid = append(transposed_tilted_grid, tilted_col)
	}
	tilted_grid := transpose(transposed_tilted_grid)
	fmt.Println(compute_weight(tilted_grid))
}
