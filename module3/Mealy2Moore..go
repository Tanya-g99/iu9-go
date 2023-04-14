package main

import (
	"fmt"
)

type vertex struct {
	q int
	y string
}

func find(X []string, x string) int {
	for i, v := range X {
		if v == x {
			return i
		}
	}
	return 0
}

func main() {
	var x, y, countq int
	fmt.Scan(&x)
	X := make([]string, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&X[i])
	}
	fmt.Scan(&y)
	Y := make([]string, y)
	for i := 0; i < y; i++ {
		fmt.Scan(&Y[i])
	}
	fmt.Scan(&countq)
	T := make([][]int, countq)
	O := make([][]int, countq)

	hashTable := make(map[vertex]int)

	for i := 0; i < countq; i++ {
		T[i] = make([]int, x)
		for j := 0; j < x; j++ {
			fmt.Scan(&T[i][j])

		}
	}
	c := 0
	for i := 0; i < countq; i++ {
		O[i] = make([]int, x)
		for j := 0; j < x; j++ {
			fmt.Scan(&O[i][j])
			_, ok := hashTable[vertex{T[i][j], Y[O[i][j]]}]
			if !ok {
				hashTable[vertex{T[i][j], Y[O[i][j]]}] = c
				c++
			}
		}
	}

	fmt.Println("digraph {\n\trankdir = LR")

	for i := 0; i < countq; i++ {
		for k := 0; k < y; k++ {
			cv, ok := hashTable[vertex{i, Y[k]}]
			if ok {
				fmt.Printf("\t%v      [label = \"(%v,%v)\"]\n", cv, i, Y[k])
				for j := 0; j < x; j++ {
					fmt.Printf("\t%v -> %v [label = \"%v\"]\n", cv, hashTable[vertex{T[i][j], Y[O[i][j]]}], X[j])
				}
			}
		}
	}
	fmt.Println("}")
}
