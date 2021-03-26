package main

import (
	"fmt"
)

var up = "U"
var down = "D"

func main() {
	numberOfStep := 8
	steps := []string{up, down, down, down, up, down, up, up}

	fmt.Println(numberOfValleys(numberOfStep, steps))
}

func numberOfValleys(n int, steps []string) int {
	level := 0
	isInValley := false
	result := 0

	for _, step := range steps {
		if step == up {
			level++
		} else {
			level--
		}

		if level < 0 && !isInValley {
			result++
			isInValley = true
		}

		if level == 0 && isInValley {
			isInValley = false
		}
	}

	return result
}
