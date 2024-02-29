package main

import (
	"fmt"

	"github.com/ablicq/adventofcode/utils"
)

type Node struct {
	left  string
	right string
}

func startNodes(net map[string]Node) []string {
	starts := []string{}
	for k := range net {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}
	return starts
}

func journey(net map[string]Node, steps string, start string) int {
	i := 0
	node := start
	for node[2] != 'Z' {
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
	return i
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm2(a, b int) int {
	return a * b / gcd(a, b)
}

func lcm(nums ...int) int {
	ret := 1
	for _, n := range nums {
		ret = lcm2(ret, n)
	}
	return ret
}

func main() {
	input := utils.ParseInput("../input.txt")
	steps := input[0]
	net := map[string]Node{}
	for _, line := range input[2:] {
		net[line[:3]] = Node{line[7:10], line[12:15]}
	}
	nodes := startNodes(net)
	journeys := make([]int, len(nodes))
	for i, n := range nodes {
		journeys[i] = journey(net, steps, n)
	}

	fmt.Println(lcm(journeys...))
}
