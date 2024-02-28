package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseInput(inputPath string) []string {
	readFile, err := os.Open(inputPath)

	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	ret := make([]string, 0)

	for fileScanner.Scan() {
		ret = append(ret, fileScanner.Text())
	}

	return ret
}

func ParseIntList(line string) []int {
	parts := strings.Split(line, " ")
	ret := make([]int, 0, len(parts))
	for _, p := range parts {
		if i, err := strconv.Atoi(p); err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
