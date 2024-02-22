package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type position struct {
	i int
	j int
}

type number struct {
	value int
	pos   []position
}

func getNumbers(schematics [][]rune) []number {
	numbers := make([]number, 0)
	current_number := make([]rune, 0)
	current_pos := make([]position, 0)
	for i, line := range schematics {
		for j, char := range line {
			if char >= '0' && char <= '9' {
				current_number = append(current_number, char)
				current_pos = append(current_pos, position{i, j})
			} else if len(current_number) != 0 {
				current_value, err := strconv.Atoi(string(current_number))
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number{current_value, current_pos})
				current_number = nil
				current_pos = nil
			}
		}
		if len(current_number) != 0 {
			current_value, err := strconv.Atoi(string(current_number))
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number{current_value, current_pos})
		}
		current_number = nil
		current_pos = nil
	}
	return numbers
}

func getNumbersMask(numbers []number, schematics [][]rune) [][]int {
	mask := make([][]int, len(schematics))
	for i := 0; i < len(schematics); i++ {
		mask[i] = make([]int, len(schematics[i]))
		for j := 0; j < len(schematics[i]); j++ {
			mask[i][j] = -1
		}
	}
	for idx, n := range numbers {
		for _, p := range n.pos {
			mask[p.i][p.j] = idx
		}
	}
	return mask
}

func checkGearPos(i, j int, mask [][]int, numbers []number) int {
	height := len(mask)
	width := len(mask[0])
	gearPosMap := make(map[int]bool)
	for di := -1; di <= 1; di++ {
		ci := i + di
		if ci < 0 || ci >= height {
			continue
		}
		for dj := -1; dj <= 1; dj++ {
			cj := j + dj
			if cj < 0 || cj >= width {
				continue
			}
			if mask[ci][cj] >= 0 {
				gearPosMap[mask[ci][cj]] = true
			}
		}
	}
	if len(gearPosMap) == 2 {
		mul := 1
		for k := range gearPosMap {
			mul *= numbers[k].value
		}
		return mul
	}
	return 0
}

func computeSum(schematics [][]rune, mask [][]int, numbers []number) int {
	sum := 0
	for i, line := range schematics {
		for j, char := range line {
			if char == '*' {
				sum += checkGearPos(i, j, mask, numbers)
			}
		}
	}
	return sum
}

func parseSchematics(scanner *bufio.Scanner) [][]rune {
	scanner.Split(bufio.ScanLines)

	schematics := make([][]rune, 0)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		schematics = append(schematics, line)
	}

	return schematics
}

func main() {
	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	schematics := parseSchematics(fileScanner)
	numbers := getNumbers(schematics)
	mask := getNumbersMask(numbers, schematics)

	fmt.Println(computeSum(schematics, mask, numbers))
}
