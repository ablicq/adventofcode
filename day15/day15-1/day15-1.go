package main

import (
	"fmt"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = (h + int(c)) * 17 % 256
	}
	return h
}

func main() {
	input := strings.Split(utils.ParseInput("../input.txt")[0], ",")
	sum := 0
	for _, s := range input {
		sum += hash(s)
	}
	fmt.Println(sum)
}
