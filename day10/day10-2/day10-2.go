package main

import (
	"fmt"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

type pos struct {
	i, j int
}

func equalPos(p1, p2 pos) bool {
	return p1.i == p2.i && p1.j == p2.j
}

func StartPos(grid []string) pos {
	for i, line := range grid {
		j := strings.Index(line, "S")
		if j != -1 {
			return pos{i, j}
		}
	}
	return pos{-1, -1}
}

func StartScouts(grid []string, start pos) (pos, pos) {
	scouts := make([]pos, 0, 2)
	upPos := pos{start.i - 1, start.j}
	if upPos.i >= 0 && strings.ContainsRune("7|F", rune(grid[upPos.i][upPos.j])) {
		scouts = append(scouts, upPos)
	}
	downPos := pos{start.i + 1, start.j}
	if downPos.i < len(grid) && strings.ContainsRune("J|L", rune(grid[downPos.i][downPos.j])) {
		scouts = append(scouts, downPos)
	}
	leftPos := pos{start.i, start.j - 1}
	if leftPos.j >= 0 && strings.ContainsRune("L-F", rune(grid[leftPos.i][leftPos.j])) {
		scouts = append(scouts, leftPos)
	}
	rightPos := pos{start.i, start.j + 1}
	if rightPos.j < len(grid[start.i]) && strings.ContainsRune("J-7", rune(grid[rightPos.i][rightPos.j])) {
		scouts = append(scouts, rightPos)
	}
	return scouts[0], scouts[1]
}

func StartTile(start, s1, s2 pos) rune {
	tiles := [3][3]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}
	tiles[s1.i-start.i+1][s1.j-start.j+1] = true
	tiles[s2.i-start.i+1][s2.j-start.j+1] = true

	switch {
	case tiles[1][0] && tiles[1][2]:
		return '-'
	case tiles[1][0] && tiles[0][1]:
		return 'J'
	case tiles[1][0] && tiles[2][1]:
		return '7'
	case tiles[2][1] && tiles[1][2]:
		return 'F'
	case tiles[0][1] && tiles[1][2]:
		return 'L'
	case tiles[0][1] && tiles[2][1]:
		return '|'

	}
	return '.'
}

func nextScout(grid []string, previousScout pos, scoutPos pos) pos {
	switch grid[scoutPos.i][scoutPos.j] {
	case '|':
		up := pos{scoutPos.i - 1, scoutPos.j}
		if equalPos(previousScout, up) {
			return pos{scoutPos.i + 1, scoutPos.j}
		}
		return up
	case '-':
		left := pos{scoutPos.i, scoutPos.j - 1}
		if equalPos(previousScout, left) {
			return pos{scoutPos.i, scoutPos.j + 1}
		}
		return left
	case 'L':
		up := pos{scoutPos.i - 1, scoutPos.j}
		if equalPos(previousScout, up) {
			return pos{scoutPos.i, scoutPos.j + 1}
		}
		return up
	case 'J':
		up := pos{scoutPos.i - 1, scoutPos.j}
		if equalPos(previousScout, up) {
			return pos{scoutPos.i, scoutPos.j - 1}
		}
		return up
	case '7':
		left := pos{scoutPos.i, scoutPos.j - 1}
		if equalPos(previousScout, left) {
			return pos{scoutPos.i + 1, scoutPos.j}
		}
		return left
	case 'F':
		right := pos{scoutPos.i, scoutPos.j + 1}
		if equalPos(previousScout, right) {
			return pos{scoutPos.i + 1, scoutPos.j}
		}
		return right
	}
	return pos{-1, -1}
}

func GetWeight(tile rune) float64 {
	// 'F7' and 'LJ' do not count as a crossing but 'FJ' and 'L7' count as 1, hence the 0.5 and -0.5 weights
	switch tile {
	case '|':
		return 1.0
	case 'F', 'J':
		return 0.5
	case 'L', '7':
		return -0.5
	}
	return 0.0
}

func main() {
	grid := utils.ParseInput("../input.txt")

	start := StartPos(grid)
	s1, s2 := StartScouts(grid, start)

	startTile := StartTile(start, s1, s2)

	loopMask := make([][]bool, len(grid))
	for i, line := range grid {
		loopMask[i] = make([]bool, len(line))
	}
	loopMask[start.i][start.j] = true

	previousScout, scout := start, s1
	for !equalPos(scout, start) {
		loopMask[scout.i][scout.j] = true
		previousScout, scout = scout, nextScout(grid, previousScout, scout)
	}

	// Count the crossings of the loop between each tile not on the loop and the edge of the board
	// If the number of crossings is even, then the tile is outside of the loop (tile and border on the same side of the loop)
	// If the number of crossings is odd, then the tile is inside of the loop (tile and border on opposite sides of the loop)
	interior := 0
	for i, line := range grid {
		nCrossings := 0.0
		for j, tile := range line {
			switch {
			case tile == 'S':
				nCrossings += GetWeight(startTile)
			case loopMask[i][j]:
				nCrossings += GetWeight(rune(tile))
			case int(nCrossings)%2 != 0:
				interior += 1
			}
		}
	}

	fmt.Println(interior)
}
