package main

import (
	"fmt"

	"github.com/ablicq/adventofcode/utils"
)

type Beam struct {
	x         int
	y         int
	direction string
}

type Tile struct {
	tileType string
	beams    []string
}

func beam_in_tile(b Beam, t Tile) bool {
	for _, d := range t.beams {
		if b.direction == d {
			return true
		}
	}
	return false
}

func tick(grid [][]Tile, beams []Beam) ([][]Tile, []Beam) {
	beamCandidates := make([]Beam, 0)

	for _, beam := range beams {
		tile := grid[beam.x][beam.y]
		switch tile.tileType {
		case ".":
			switch beam.direction {
			case "left":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y - 1, beam.direction})
			case "right":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y + 1, beam.direction})
			case "down":
				beamCandidates = append(beamCandidates, Beam{beam.x + 1, beam.y, beam.direction})
			case "up":
				beamCandidates = append(beamCandidates, Beam{beam.x - 1, beam.y, beam.direction})
			}
		case "\\":
			switch beam.direction {
			case "left":
				beamCandidates = append(beamCandidates, Beam{beam.x - 1, beam.y, "up"})
			case "right":
				beamCandidates = append(beamCandidates, Beam{beam.x + 1, beam.y, "down"})
			case "down":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y + 1, "right"})
			case "up":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y - 1, "left"})
			}
		case "/":
			switch beam.direction {
			case "left":
				beamCandidates = append(beamCandidates, Beam{beam.x + 1, beam.y, "down"})
			case "right":
				beamCandidates = append(beamCandidates, Beam{beam.x - 1, beam.y, "up"})
			case "down":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y - 1, "left"})
			case "up":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y + 1, "right"})
			}
		case "-":
			switch beam.direction {
			case "left":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y - 1, beam.direction})
			case "right":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y + 1, beam.direction})
			case "down":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y + 1, "right"})
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y - 1, "left"})
			case "up":
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y + 1, "right"})
				beamCandidates = append(beamCandidates, Beam{beam.x, beam.y - 1, "left"})
			}
		case "|":
			switch beam.direction {
			case "left":
				beamCandidates = append(beamCandidates, Beam{beam.x + 1, beam.y, "down"})
				beamCandidates = append(beamCandidates, Beam{beam.x - 1, beam.y, "up"})
			case "right":
				beamCandidates = append(beamCandidates, Beam{beam.x + 1, beam.y, "down"})
				beamCandidates = append(beamCandidates, Beam{beam.x - 1, beam.y, "up"})
			case "down":
				beamCandidates = append(beamCandidates, Beam{beam.x + 1, beam.y, beam.direction})
			case "up":
				beamCandidates = append(beamCandidates, Beam{beam.x - 1, beam.y, beam.direction})
			}
		}
	}

	newGrid := make([][]Tile, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]Tile, len(grid[i]))
	}
	for i := range newGrid {
		for j := range newGrid[i] {
			newGrid[i][j] = grid[i][j]
		}
	}

	newBeams := make([]Beam, 0)

	for _, beam := range beamCandidates {
		if beam.x >= 0 && beam.x < len(newGrid) && beam.y >= 0 && beam.y < len(newGrid[0]) && !beam_in_tile(beam, newGrid[beam.x][beam.y]) {
			newBeams = append(newBeams, beam)
			newGrid[beam.x][beam.y].beams = append(newGrid[beam.x][beam.y].beams, beam.direction)
		}
	}

	return newGrid, newBeams
}

func tickUntilStable(startGrid [][]Tile, startBeams []Beam) [][]Tile {
	grid, beams := startGrid, startBeams

	for len(beams) != 0 {
		grid, beams = tick(grid, beams)
	}

	return grid
}

func parseGrid(gridStr []string) [][]Tile {
	grid := make([][]Tile, len(gridStr))
	for i := range grid {
		grid[i] = make([]Tile, len(gridStr[i]))
		for j, c := range gridStr[i] {
			grid[i][j] = Tile{string(c), make([]string, 0, 4)}
		}
	}

	return grid
}

func countEnergizedTiles(grid [][]Tile) int {
	nEnergizedTiles := 0
	for i := range grid {
		for j := range grid[i] {
			if len(grid[i][j].beams) != 0 {
				nEnergizedTiles += 1
			}
		}
	}
	return nEnergizedTiles
}

func main() {
	gridStr := utils.ParseInput("../input.txt")

	startGrid := parseGrid(gridStr)
	startBeams := []Beam{{0, 0, "right"}}

	startGrid[0][0].beams = []string{"right"}

	endGrid := tickUntilStable(startGrid, startBeams)

	fmt.Println(countEnergizedTiles(endGrid))
}
