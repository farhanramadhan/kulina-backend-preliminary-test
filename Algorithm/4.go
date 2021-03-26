package main

import "fmt"

func main() {
	mapOfLamps := make(map[int]bool)

	// first trip
	for i := 1; i <= 100; i++ {
		mapOfLamps[i] = true
	}

	// O(n^2) :))))
	for i := 2; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if j%i == 0 {
				if mapOfLamps[j] {
					mapOfLamps[j] = false
				} else {
					mapOfLamps[j] = true
				}
			}
		}
	}

	for i, v := range mapOfLamps {
		if v {
			fmt.Println(i)
		}
	}
}
