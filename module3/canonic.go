package main

import "fmt"

func CanonicNumeration(matrixTransition [][]int, q0 int, n int) (int, []int, []int) {
	time := 0
	numerationT := make([]int, n)
	numeration := make([]int, n)
	for i := 0; i < n; i++ {
		numeration[i] = -1
	}
	var DFS func(int)
	DFS = func(i int) {
		if numeration[i] == -1 {
			numeration[i] = time
			numerationT[time] = i
			time++
			for j := 0; j < len(matrixTransition[i]); j++ {
				DFS(matrixTransition[i][j])
			}
		}
	}
	DFS(q0)
	return time, numerationT, numeration
}

func CanonicMatrix(matrixTransition [][]int, matrixOutput [][]string, q0 int, n int, m int) ([][]int, [][]string) {
	maxTime, numerationT, numeration := CanonicNumeration(matrixTransition, q0, n)
	mT, mO := make([][]int, maxTime), make([][]string, maxTime)
	for i := 0; i < maxTime; i++ {
		mT[i], mO[i] = make([]int, m), make([]string, m)
	}
	for i := 0; i < maxTime; i++ {
		num := numerationT[i]
		for j := 0; j < m; j++ {
			mT[i][j] = numeration[matrixTransition[num][j]]
			mO[i][j] = matrixOutput[num][j]
		}
	}
	return mT, mO

}

func main() {
	var m, n, q0 int
	fmt.Scanf("%v\n%v\n%v", &n, &m, &q0)
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

	resT, resO := CanonicMatrix(matrixTransition, matrixOutput, q0, n, m)
	fmt.Printf("%v\n%v\n0\n", n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%v ", resT[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%v ", resO[i][j])
		}
		fmt.Println()
	}

}

/* 4
3
0
1 3 3
1 1 2
2 2 2
1 2 3
x y y
y y x
x x x
x y y

3
2
1
1 0
2 0
2 2
x y
y x
x y*/
