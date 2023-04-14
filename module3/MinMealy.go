package main

import "fmt"

var parent, depth []int

func Find(x int) int {
	if x == parent[x] {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

func Union(x, y int) {
	x, y = Find(x), Find(y)
	if x != y {
		if depth[x] < depth[y] {
			x, y = y, x
		}
		parent[y] = x
		if depth[x] == depth[y] {
			depth[x]++
		}
	}
}

func NewClass(n int) {
	parent, depth = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}
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

func Split1(m_ int, pi []int, n int, m int, T [][]int, O [][]string) (int, []int) {
	m_ = n
	NewClass(n)
	for qi := 0; qi < n; qi++ {
		for qj := qi + 1; qj < n; qj++ {
			if Find(qi) != Find(qj) {
				eq := true
				for x := 0; x < m; x++ {
					if O[qi][x] != O[qj][x] {
						eq = false
						break
					}
				}
				if eq {
					Union(qi, qj)
					m_--
				}
			}
		}
	}
	for qi := 0; qi < n; qi++ {
		pi[qi] = Find(qi)
	}

	return m_, pi
}

func Split(m_ int, pi []int, n int, m int, T [][]int, O [][]string) (int, []int) {
	m_ = n
	NewClass(n)
	for qi := 0; qi < n; qi++ {
		for qj := qi + 1; qj < n; qj++ {
			if pi[qi] == pi[qj] && Find(qi) != Find(qj) {
				eq := true
				for x := 0; x < m; x++ {
					w1, w2 := pi[T[qi][x]], pi[T[qj][x]]
					if w1 != w2 {
						eq = false
						break
					}
				}
				if eq {
					Union(qi, qj)
					m_--
				}
			}
		}
	}
	for qi := 0; qi < n; qi++ {
		pi[qi] = Find(qi)
	}

	return m_, pi
}

func AufenkampHohn(T [][]int, O [][]string, n int, m int, q0 int) ([][]int, [][]string, int, int, int) {
	pi := make([]int, n)
	var m_, M_ int
	m_, pi = Split1(m_, pi, n, m, T, O)
	for {
		M_, pi = Split(M_, pi, n, m, T, O)
		if m_ == M_ {
			break
		}
		m_ = M_
	}

	a, b := make([]int, n), make([]int, n)
	for i, j := 0, 0; i < n; i++ {
		if pi[i] == i {
			a[j], b[i] = i, j
			j++
		}
	}

	res_n := m_
	res_m := m
	res_q0 := b[pi[q0]]
	matrixTransition := make([][]int, res_n)
	matrixOutput := make([][]string, res_n)
	for i := 0; i < res_n; i++ {
		matrixTransition[i] = make([]int, res_m)
		matrixOutput[i] = make([]string, res_m)
	}

	for i := 0; i < res_n; i++ {
		for j := 0; j < res_m; j++ {
			matrixTransition[i][j] = b[pi[T[a[i]][j]]]
			matrixOutput[i][j] = O[a[i]][j]
		}
	}
	return matrixTransition, matrixOutput, res_n, res_m, res_q0
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

	minT, minO, new_n, new_m, q := AufenkampHohn(matrixTransition, matrixOutput, n, m, q0)
	resT, resO := CanonicMatrix(minT, minO, q, new_n, new_m)

	fmt.Println("digraph {\n\trankdir = LR")
	for i := 0; i < new_n; i++ {
		for j := 0; j < new_m; j++ {
			fmt.Printf("\t%v -> %v [label = \"%c(%v)\"]\n", i, resT[i][j], 'a'+j, resO[i][j])
		}
	}
	fmt.Println("}")

}
