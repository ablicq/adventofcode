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

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func slice_equal(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
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

func tilt(grid []string, direction string) []string {
	switch direction {
	case "north":
		tilted_grid := make([]string, 0, len(grid[0]))
		for _, col := range transpose(grid) {
			tilted_col := tilt_one_column(col)
			tilted_grid = append(tilted_grid, tilted_col)
		}
		return transpose(tilted_grid)
	case "west":
		tilted_grid := make([]string, 0, len(grid[0]))
		for _, col := range grid {
			tilted_col := tilt_one_column(col)
			tilted_grid = append(tilted_grid, tilted_col)
		}
		return tilted_grid
	case "south":
		tilted_grid := make([]string, 0, len(grid[0]))
		for _, col := range transpose(grid) {
			tilted_col := reverse(tilt_one_column(reverse(col)))
			tilted_grid = append(tilted_grid, tilted_col)
		}
		return transpose(tilted_grid)
	case "east":
		tilted_grid := make([]string, 0, len(grid[0]))
		for _, col := range grid {
			tilted_col := reverse(tilt_one_column(reverse(col)))
			tilted_grid = append(tilted_grid, tilted_col)
		}
		return tilted_grid
	default:
		return grid
	}
}

func spin_cycle(grid []string) []string {
	n_tilt := tilt(grid, "north")
	w_tilt := tilt(n_tilt, "west")
	s_tilt := tilt(w_tilt, "south")
	e_tilt := tilt(s_tilt, "east")
	return e_tilt
}

func detect_period(grid []string) (int, int) {
	turtoise := spin_cycle(grid)
	hare := spin_cycle(spin_cycle(grid))
	for !slice_equal(turtoise, hare) {
		turtoise = spin_cycle(turtoise)
		hare = spin_cycle(spin_cycle(hare))
	}

	mu := 0
	turtoise = grid
	for !slice_equal(turtoise, hare) {
		turtoise = spin_cycle(turtoise)
		hare = spin_cycle(hare)
		mu++
	}

	lam := 1
	hare = spin_cycle(turtoise)
	for !slice_equal(turtoise, hare) {
		hare = spin_cycle(hare)
		lam++
	}

	return lam, mu
}

func iter_1B(grid []string) []string {
	lam, mu := detect_period(grid)
	fmt.Println(lam, mu)
	n_spin_cycles := (1e9-mu)%lam + mu

	for range n_spin_cycles {
		grid = spin_cycle(grid)
	}

	return grid
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
	tilted_grid := iter_1B(grid)
	fmt.Println(compute_weight(tilted_grid))
}
