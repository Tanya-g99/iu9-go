package main

import (
	"fmt"
)

func econom(str string) int {
	expr := []rune(str)
	hashTable := make(map[string]bool)
	var beginIndex []int
	for index, elem := range expr {
		if elem == '(' {
			beginIndex = append(beginIndex, index+1)
		} else if elem == ')' {
			end := len(beginIndex) - 1
			minexpr := string(expr[beginIndex[end]:index])
			beginIndex = beginIndex[:end]
			//fmt.Println(minexpr)
			hashTable[minexpr] = true
		}
	}
	return len(hashTable)
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(econom(str))
}
