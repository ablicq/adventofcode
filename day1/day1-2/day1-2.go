package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numbers = map[string]rune{
	"0":     '0',
	"1":     '1',
	"2":     '2',
	"3":     '3',
	"4":     '4',
	"5":     '5',
	"6":     '6',
	"7":     '7',
	"8":     '8',
	"9":     '9',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func findDigit(s []rune) (rune, bool) {
	if len(s) >= 5 {
		s5 := string(s[:5])
		if r, ok := numbers[s5]; ok {
			return r, true
		}
	}
	if len(s) >= 4 {
		s4 := string(s[:4])
		if r, ok := numbers[s4]; ok {
			return r, true
		}
	}
	if len(s) >= 3 {
		s3 := string(s[:3])
		if r, ok := numbers[s3]; ok {
			return r, true
		}
	}
	if len(s) >= 1 {
		s1 := string(s[:1])
		if r, ok := numbers[s1]; ok {
			return r, true
		}
	}
	return 0, false
}

func firstDigit(line []rune) rune {
	for i := 0; i < len(line); i++ {
		if r, ok := findDigit(line[i:]); ok {
			return r
		}
	}
	return '0'
}

func lastDigit(line []rune) rune {
	for i := len(line) - 1; i >= 0; i-- {
		if r, ok := findDigit(line[i:]); ok {
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
