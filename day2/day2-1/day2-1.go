package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func processLine(line string) int {
	parts := strings.Split(line, ": ")

	turns := strings.Split(parts[1], "; ")
	for _, turn := range turns {
		cubes := strings.Split(turn, ", ")
		for _, cube := range cubes {
			props := strings.Split(cube, " ")

			nCubes := atoi(props[0])
			color := props[1]
			if nCubes > limits[color] {
				return 0
			}
		}
	}

	return atoi(parts[0][5:])
}

func main() {
	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		sum += processLine(fileScanner.Text())
	}

	fmt.Println(sum)

	readFile.Close()
}
