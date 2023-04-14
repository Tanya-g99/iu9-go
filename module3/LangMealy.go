package main

import (
	"fmt"
	"sort"
)

var M int
var hashTable map[string][]int
var O [][]string
var T [][]int

func f(q0 int, res string, n int) {
	if O[q0][n] != "-" {
		res = res + O[q0][n]
	}
	if len(res) > M {
		return
	}
	q := T[q0][n]
	var t bool
	A, ok := hashTable[res]
	for _, v := range A {
		if v == q {
			t = true
		}
	}
	if !ok {
		hashTable[res] = make([]int, 0)
	}
	if !t {
		hashTable[res] = append(A, q)

		f(q, res, 0)
		f(q, res, 1)

	}

}

func main() {
	var q, q0 int
	fmt.Scan(&q)
	T = make([][]int, q)
	O = make([][]string, q)
	for i := 0; i < q; i++ {
		T[i] = make([]int, 2)
		for j := range []string{"a", "b"} {
			fmt.Scan(&T[i][j])
		}
	}
	for i := 0; i < q; i++ {
		O[i] = make([]string, 2)
		for j := range []string{"a", "b"} {
			fmt.Scan(&O[i][j])
		}
	}

	fmt.Scan(&q0)
	fmt.Scan(&M)
	hashTable = make(map[string][]int)
	f(q0, "", 0)
	f(q0, "", 1)

	var keys []string
	//fmt.Println(hashTable)
	for i := range hashTable {
		if i != "" {
			keys = append(keys, i)
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%v ", k)
	}
	fmt.Println()
}
