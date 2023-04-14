package main

import "fmt"

func min(la, lb int) int {
	if la > lb {
		return lb
	}
	return la
}
func max(la, lb int) int {
	if la < lb {
		return lb
	} else if la > lb {
		return la
	}
	return 0
}

func inputLongInt() []int32 {
	var n int
	fmt.Scan(&n)
	num := make([]int32, n)
	for i := range num {
		fmt.Scan(&num[i])
	}
	return num
}

func add(a, b []int32, p int) []int32 {
	var c []int32
	var flag int32 = 0
	la := len(a)
	lb := len(b)
	if lb > la {
		a, b = b, a
	}
	lmin := min(la, lb)
	lmax := max(la, lb)
	for i := 0; i < lmin; i++ {
		c = append(c, (a[i]+b[i]+flag)%int32(p))
		flag = (a[i] + b[i] + flag) / int32(p)
	}
	for i := lmin; i < lmax; i++ {
		c = append(c, (a[i]+flag)%int32(p))
		flag = (a[i] + flag) / int32(p)
	}
	if flag > 0 {
		c = append(c, flag)
	}
	return c
}

func main() {
	//a := inputLongInt()
	//b := inputLongInt()
	a := []int32{1}    // 0xdcba
	b := []int32{1, 1} // 0xabcd
	p := 3
	fmt.Println(add(a, b, p))
}
