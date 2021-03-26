package main

import "fmt"

func main() {
	number := 1345679

	printPerNumber(fmt.Sprint(number))
}

func printPerNumber(n string) {
	resultArr := []string{}
	zeroString := ""
	length := len(n) - 1
	for i := 0; i < length; i++ {
		zeroString += "0"
	}

	for _, v := range n {
		resultArr = append(resultArr, string(v)+zeroString)
		if len(zeroString) > 0 {
			zeroString = zeroString[:len(zeroString)-1]
		}
	}

	for _, v := range resultArr {
		fmt.Println(v)
	}
}
