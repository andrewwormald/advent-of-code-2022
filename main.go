package main

import (
	"adventOfCode/challenges"
	"adventOfCode/inputs"
	"fmt"
)

func main() {
	res := challenges.MonkeyInTheMiddle(inputs.Monkey, 10000)
	fmt.Println(res)
}
