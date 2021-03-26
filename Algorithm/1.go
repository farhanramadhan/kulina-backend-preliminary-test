package main

import "fmt"

func main() {
	numberOfSocks := 9
	socks := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}

	fmt.Println(pairOfSocks(numberOfSocks, socks))
}

func pairOfSocks(n int, socks []int) int {
	result := 0
	mapOfSocks := make(map[int]int)

	for _, sock := range socks {
		mapOfSocks[sock]++
	}

	for _, sock := range mapOfSocks {
		result += sock / 2
	}

	return result
}
