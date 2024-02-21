package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func processLine(line string) int {
	parts := strings.Split(line, ": ")

	var colorCounts = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	turns := strings.Split(parts[1], "; ")
	for _, turn := range turns {
		cubes := strings.Split(turn, ", ")
		for _, cube := range cubes {
			props := strings.Split(cube, " ")

			nCubes := atoi(props[0])
			color := props[1]
			if nCubes > colorCounts[color] {
				colorCounts[color] = nCubes
			}
		}
	}

	return colorCounts["red"] * colorCounts["green"] * colorCounts["blue"]
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
