package main

import "fmt"

type DSU struct {
	parent, depth []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{make([]int, n), make([]int, n)}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
	}
	return dsu
}

func (dsu *DSU) Find(x int) int {
	if x == dsu.parent[x] {
		return x
	}
	dsu.parent[x] = dsu.Find(dsu.parent[x])
	return dsu.parent[x]
}

func (dsu *DSU) Union(x, y int) {
	x, y = dsu.Find(x), dsu.Find(y)
	if x != y {
		if dsu.depth[x] < dsu.depth[y] {
			x, y = y, x
		}
		dsu.parent[y] = x
		if dsu.depth[x] == dsu.depth[y] {
			dsu.depth[x]++
		}
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
	dsu := NewDSU(n)
	for qi := 0; qi < n; qi++ {
		for qj := qi + 1; qj < n; qj++ {
			if dsu.Find(qi) != dsu.Find(qj) {
				eq := true
				for x := 0; x < m; x++ {
					if O[qi][x] != O[qj][x] {
						eq = false
						break
					}
				}
				if eq {
					dsu.Union(qi, qj)
					m_--
				}
			}
		}
	}
	for qi := 0; qi < n; qi++ {
		pi[qi] = dsu.Find(qi)
	}

	return m_, pi
}

func Split(m_ int, pi []int, n int, m int, T [][]int, O [][]string) (int, []int) {
	m_ = n
	dsu := NewDSU(n)
	for qi := 0; qi < n; qi++ {
		for qj := qi + 1; qj < n; qj++ {
			if pi[qi] == pi[qj] && dsu.Find(qi) != dsu.Find(qj) {
				eq := true
				for x := 0; x < m; x++ {
					w1, w2 := pi[T[qi][x]], pi[T[qj][x]]
					if w1 != w2 {
						eq = false
						break
					}
				}
				if eq {
					dsu.Union(qi, qj)
					m_--
				}
			}
		}
	}
	for qi := 0; qi < n; qi++ {
		pi[qi] = dsu.Find(qi)
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
	res_n := m_
	res_m := m
	matrixTransition := make([][]int, res_n)
	matrixOutput := make([][]string, res_n)
	for i := 0; i < res_n; i++ {
		matrixTransition[i] = make([]int, res_m)
		matrixOutput[i] = make([]string, res_m)
	}

	a, b := make([]int, n), make([]int, n)
	c := 0
	for i := 0; i < n; i++ {
		if pi[i] == i {
			a[c] = i
			b[i] = c
			c++
		}
	}
	res_q0 := b[pi[q0]]

	for i := 0; i < res_n; i++ {
		for j := 0; j < res_m; j++ {
			matrixTransition[i][j] = b[pi[T[a[i]][j]]]
			matrixOutput[i][j] = O[a[i]][j]
		}
	}
	return matrixTransition, matrixOutput, res_n, res_m, res_q0
}

func EqMealy(T1 [][]int, O1 [][]string, q1 int, n1 int, m1 int, T2 [][]int, O2 [][]string, q2 int, n2 int, m2 int) {
	if q1 != q2 || n1 != n2 || m1 != m2 {
		fmt.Println("NOT EQUAL")
		return
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < m1; j++ {
			if T1[i][j] != T2[i][j] || O1[i][j] != O2[i][j] {
				fmt.Println("NOT EQUAL")
				return
			}
		}
	}

	fmt.Println("EQUAL")

}

func main() {
	var m1, n1, q01, m2, n2, q02 int
	fmt.Scanf("%v\n%v\n%v", &n1, &m1, &q01)
	matrixTransition1 := make([][]int, n1)
	matrixOutput1 := make([][]string, n1)

	for i := 0; i < n1; i++ {
		matrixTransition1[i] = make([]int, m1)
		for j := 0; j < m1; j++ {
			fmt.Scan(&matrixTransition1[i][j])
		}
	}
	for i := 0; i < n1; i++ {
		matrixOutput1[i] = make([]string, m1)
		for j := 0; j < m1; j++ {
			fmt.Scan(&matrixOutput1[i][j])
		}
	}

	minT1, minO1, new_n1, new_m1, q1 := AufenkampHohn(matrixTransition1, matrixOutput1, n1, m1, q01)
	resT1, resO1 := CanonicMatrix(minT1, minO1, q1, new_n1, new_m1)

	fmt.Scanf("%v\n%v\n%v", &n2, &m2, &q02)
	matrixTransition2 := make([][]int, n2)
	matrixOutput2 := make([][]string, n2)

	for i := 0; i < n2; i++ {
		matrixTransition2[i] = make([]int, m2)
		for j := 0; j < m2; j++ {
			fmt.Scan(&matrixTransition2[i][j])
		}
	}
	for i := 0; i < n2; i++ {
		matrixOutput2[i] = make([]string, m2)
		for j := 0; j < m2; j++ {
			fmt.Scan(&matrixOutput2[i][j])
		}
	}

	minT2, minO2, new_n2, new_m2, q2 := AufenkampHohn(matrixTransition2, matrixOutput2, n2, m2, q02)
	resT2, resO2 := CanonicMatrix(minT2, minO2, q2, new_n2, new_m2)

	EqMealy(resT1, resO1, q1, new_n1, new_m1, resT2, resO2, q2, new_n2, new_m2)

}
