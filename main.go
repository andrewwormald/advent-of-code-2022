package main

import (
	"fmt"

	"adventOfCode/challenges"
	"adventOfCode/inputs"
)

func main() {
	result := challenges.CountPairsOfOverlap(inputs.CampCleanup)
	fmt.Println(result)
}
