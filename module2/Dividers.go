package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
)

func main() {
	var n int
	bufstdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(bufstdin, &n)

	fmt.Println("graph {")
	var dividers []int

	for i := 1; i*i <= n; i++ {

		if n%i == 0 {
			dividers = append(dividers, i)
			fmt.Println(i)
			if i != n/i {
				dividers = append(dividers, n/i)
				fmt.Println(n / i)
			}
		}
	}
	sort.Ints(dividers)

	for _, v := range dividers {
		for _, v2 := range dividers {
			if v != v2 {
				if v2%v == 0 && big.NewInt(int64(v2/v)).ProbablyPrime(0) {
					fmt.Printf("%v -- %v\n", v2, v)
				}
			}
		}
	}
	fmt.Println("}")
}
