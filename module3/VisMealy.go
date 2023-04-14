package main

import "fmt"

func Char(i int) rune {
	return rune('a' + i)
}
func main() {
	var n, m, q0 int
	fmt.Scanf("%v\n%v\n%v\n", &n, &m, &q0)

	matrixTransition := make([][]int, n)
	matrixOutput := make([][]string, n)

	for i := 0; i < n; i++ {
		matrixTransition[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrixTransition[i][j])
		}
	}
	for i := 0; i < n; i++ {
		matrixOutput[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrixOutput[i][j])
		}
	}

	fmt.Println("digraph {\n\trankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("\t%v -> %v [label = \"%c(%v)\"]\n", i, matrixTransition[i][j], Char(j), matrixOutput[i][j])
		}
	}
	fmt.Println("}")
}
