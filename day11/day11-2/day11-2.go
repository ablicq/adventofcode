package main

import (
	"fmt"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

type galaxy struct {
	i int
	j int
}

func (g galaxy) dist(other galaxy) int {
	return max(g.i, other.i) - min(g.i, other.i) + max(g.j, other.j) - min(g.j, other.j)
}

func findGalaxies(grid []string) []galaxy {
	galaxies := make([]galaxy, 0)
	for i, line := range grid {
		for j, tile := range line {
			if tile == '#' {
				galaxies = append(galaxies, galaxy{i, j})
			}
		}
	}
	return galaxies
}

func isColumnEmpty(grid []string, j int) bool {
	for i := 0; i < len(grid); i++ {
		if grid[i][j] != '.' {
			return false
		}
	}
	return true
}

func applyExpansion(grid []string, galaxies []galaxy) []galaxy {
	factors_i := make([]int, len(grid))
	factors_j := make([]int, len(grid[0]))
	cpt := 0
	for i, line := range grid {
		if line == strings.Repeat(".", len(line)) {
			cpt += 999999
		}
		factors_i[i] = cpt
	}

	cpt = 0
	for j := 0; j < len(grid[0]); j++ {
		if isColumnEmpty(grid, j) {
			cpt += 999999
		}
		factors_j[j] = cpt
	}

	expanded_galaxies := make([]galaxy, len(galaxies))
	for i, g := range galaxies {
		expanded_galaxies[i] = galaxy{g.i + factors_i[g.i], g.j + factors_j[g.j]}
	}

	return expanded_galaxies
}

func main() {
	grid := utils.ParseInput("../input.txt")

	galaxies := findGalaxies(grid)
	expanded_galaxies := applyExpansion(grid, galaxies)

	cpt := 0
	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < i; j++ {
			cpt += expanded_galaxies[i].dist(expanded_galaxies[j])
		}
	}

	fmt.Println(cpt)
}
