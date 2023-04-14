package main

import "fmt"

func qsort(n int, less func(i, j int) bool, swap func(i, l int)) {
	var qsort_rec func(left, right int)
	partition := func(left, right int) int {
		border := left
		for i := left + 1; i < right; i++ {
			if less(i, border) {
				border++
				if i != border {
					swap(border-1, border)
				}
				swap(i, border-1)
			}
		}
		return border
	}
	qsort_rec = func(left, right int) {
		if left < right {
			p := partition(left, right)
			qsort_rec(left, p)
			qsort_rec(p+1, right)
		}
	}

	qsort_rec(0, n)
}

func main() {
	var n int

	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	less := func(i, j int) bool { return array[i] < array[j] }
	swap := func(i, j int) { array[i], array[j] = array[j], array[i] }

	qsort(n, less, swap)

	for _, i := range array {
		fmt.Printf("%d ", i)
	}

}
