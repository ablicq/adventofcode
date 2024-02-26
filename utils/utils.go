package utils

import (
	"bufio"
	"os"
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
