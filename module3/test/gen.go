package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10
	rand.Seed(time.Now().UnixNano())
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Println(0+rand.Intn(n-0), " ", 0+rand.Intn(n-0))
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(3)
		if j == 1 {
			fmt.Print("x ")
		} else if j == 2 {

			fmt.Print("y ")
		} else {
			fmt.Print("- ")
		}
		j = rand.Intn(3)
		if j == 1 {
			fmt.Print("x ")
		} else if j == 2 {

			fmt.Print("y ")
		} else {
			fmt.Print("- ")
		}
		fmt.Println()
	}
	fmt.Println(0)
	fmt.Println(5)
}
