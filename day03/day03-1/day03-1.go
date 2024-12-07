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

func checkNumber(n number, schematics [][]rune) bool {
	height := len(schematics)
	width := len(schematics[0])
	for _, pos := range n.pos {
		for di := -1; di <= 1; di++ {
			i := pos.i + di
			if i < 0 || i >= height {
				continue
			}
			for dj := -1; dj <= 1; dj++ {
				j := pos.j + dj
				if j < 0 || j >= width {
					continue
				}
				if (schematics[i][j] < '0' || schematics[i][j] > '9') && schematics[i][j] != '.' {
					return true
				}
			}
		}
	}
	return false
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
	sum := 0

	for _, n := range numbers {
		if checkNumber(n, schematics) {
			sum += n.value
		}
	}

	fmt.Println(sum)
}
