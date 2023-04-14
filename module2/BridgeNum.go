package main

import "fmt"

var g [][]int
var res, tin, fup []int
var used []bool
var timer int

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dfs(v, p int) {
	used[v] = true
	timer++
	tin[v] = timer
	fup[v] = timer

	for i := 0; i < len(g[v]); i++ {
		to := g[v][i]
		if to == p {
			continue
		}
		if used[to] {
			fup[v] = min(fup[v], tin[to])
		} else {
			dfs(to, v)
			fup[v] = min(fup[v], fup[to])
			if fup[to] > tin[v] {
				res = append(res, 1)
			}
		}
	}
}

func find_bridges(n int) {
	timer = 0
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i, -1)
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	g = make([][]int, n)
	tin = make([]int, n)
	fup = make([]int, n)
	used = make([]bool, n)
	res = make([]int, 0)

	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	find_bridges(n)

	fmt.Println(len(res))

}
