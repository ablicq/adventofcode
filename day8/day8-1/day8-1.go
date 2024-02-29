package main

import (
	"fmt"

	"github.com/ablicq/adventofcode/utils"
)

type Node struct {
	left  string
	right string
}

func main() {
	input := utils.ParseInput("../input.txt")
	steps := input[0]
	net := map[string]Node{}
	for _, line := range input[2:] {
		net[line[:3]] = Node{line[7:10], line[12:15]}
	}
	i := 0
	node := "AAA"
	for node != "ZZZ" {
		for _, step := range steps {
			switch step {
			case 'L':
				node = net[node].left
			case 'R':
				node = net[node].right
			}
			i++
		}
	}
	fmt.Println(i)
}
