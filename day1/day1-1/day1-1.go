package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func firstDigit(line []rune) rune {
	for i := 0; i < len(line); i++ {
		r := line[i]
		if r >= '0' && r <= '9' {
			return r
		}
	}
	return '0'
}

func lastDigit(line []rune) rune {
	for i := len(line) - 1; i >= 0; i-- {
		r := line[i]
		if r >= '0' && r <= '9' {
			return r
		}
	}
	return '0'
}

func processLine(line string) int {
	lineRunes := []rune(line)
	firstDigit, lastDigit := firstDigit(lineRunes), lastDigit(lineRunes)
	lineDigitStr := string([]rune{firstDigit, lastDigit})
	lineDigit, err := strconv.Atoi(lineDigitStr)

	if err != nil {
		panic(err)
	}

	return lineDigit
}

func main() {
	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		sum += processLine(fileScanner.Text())
	}

	fmt.Println(sum)

	readFile.Close()
}
