package main

import (
	"fmt"
	"math"

	"github.com/ablicq/adventofcode/utils"
)

func computeRootsSplit(T, D int) int {
	r_delta := math.Sqrt(float64(T*T - 4*D))
	H1 := (float64(T) - r_delta) / 2
	H2 := (float64(T) + r_delta) / 2

	rootSplit := 0
	if H2 == math.Floor(H2) {
		// Handle case were the upper root lands on an int
		// Roots beings the hold times for which the distance travelled is exactly D,
		// They are not considered valid solutions
		rootSplit = int(math.Floor(H2)) - int(math.Floor(H1)) - 1
	} else {
		rootSplit = int(math.Floor(H2)) - int(math.Floor(H1))
	}

	return rootSplit
}

func main() {
	input := utils.ParseInput("../input.txt")

	times := utils.ParseIntList(input[0])
	distances := utils.ParseIntList(input[1])

	product := 1

	for i := 0; i < len(times); i++ {
		product *= computeRootsSplit(times[i], distances[i])
	}

	fmt.Println(product)
}
