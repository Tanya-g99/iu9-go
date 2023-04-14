package main

import "fmt"

func main() {

	//var hashTable map[[][][]string]bool
	//var delta [][][]string

	var n, m int
	fmt.Scan(&n)

	delta := make([][][]string, n)
	for i := 0; i < n; i++ {
		delta[i] = make([][]string, n)
	}

	fmt.Scan(&m)

	X := make(map[string]bool)

	for i := 0; i < m; i++ {
		var q1, q2 int
		var symbol string
		fmt.Scanf("%v %v %v\n", &q1, &q2, &symbol)
		_, ok := X[symbol]
		if !ok && symbol != "lambda" {
			X[symbol] = true
		}
		delta[q1][q2] = append(delta[q1][q2], symbol)
	}

}
