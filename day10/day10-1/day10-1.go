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

func main() {
	grid := utils.ParseInput("../input.txt")

	start := StartPos(grid)
	previousScout1, previousScout2 := start, start
	scout1, scout2 := StartScouts(grid, start)
	cpt := 1
	for !equalPos(scout1, scout2) {
		previousScout1, scout1 = scout1, nextScout(grid, previousScout1, scout1)
		previousScout2, scout2 = scout2, nextScout(grid, previousScout2, scout2)
		cpt++
	}

	fmt.Println(cpt)
}
