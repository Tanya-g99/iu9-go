package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func index(expr []rune) int {
	i := 0
	if unicode.IsDigit(expr[0]) {
		return 1
	}
	for index, v := range expr {
		//fmt.Print(string(v))
		if v == '+' || v == '*' || v == '-' {
			i++
		} else {
			if i == 0 {
				return index + 1
			}
			i--
		}
	}
	return len(expr) + 1
}

func pol_rec(expr []rune) int {
	end := len(expr)
	/* for _, v := range expr {
		fmt.Print(string(v))
	}
	fmt.Println() */
	if unicode.IsDigit(expr[0]) {
		digit, _ := strconv.Atoi(string(expr[0]))
		return digit
	}

	index := index(expr[1:end])
	//fmt.Print(" index:  ", index, "\n")
	if index != 0 && index+1 != end {
		//fmt.Println("Sit 1")
		if expr[0] == '+' {
			return pol_rec(expr[1:index+1]) + pol_rec(expr[index+1:end])
		} else if expr[0] == '*' {
			return pol_rec(expr[1:index+1]) * pol_rec(expr[index+1:end])

		} else {
			return pol_rec(expr[1:index+1]) - pol_rec(expr[index+1:end])
		}
	} else if index != 0 {
		//fmt.Println("Sit 2")
		digit, _ := strconv.Atoi(string(expr[index]))
		if expr[0] == '+' {
			return pol_rec(expr[1:index+1]) + digit
		} else if expr[0] == '*' {
			return pol_rec(expr[1:index+1]) * digit

		} else {
			return pol_rec(expr[1:index+1]) - digit
		}
	} else if index != end {
		//fmt.Println("Sit 3")
		digit, _ := strconv.Atoi(string(expr[1]))
		if expr[0] == '+' {
			return digit + pol_rec(expr[index+1:end])
		} else if expr[0] == '*' {
			return digit * pol_rec(expr[index+1:end])

		} else {
			return digit - pol_rec(expr[index+1:end])
		}
	} else {
		//fmt.Println("Sit 4")
		digit1, _ := strconv.Atoi(string(expr[1]))
		digit2, _ := strconv.Atoi(string(expr[end-1]))
		if expr[0] == '+' {
			return digit1 + digit2
		} else if expr[0] == '*' {
			return digit1 * digit2

		} else {
			return digit1 - digit2

		}
	}
}

func polish(str string) int {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	runes := []rune(str)
	len := len(runes)
	if len == 0 {
		return 0
	}
	return pol_rec(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()
	fmt.Println(polish(expression))
} //wants := []int{0, 3, 12, 729, 35, 135, 1, -1, 0}
