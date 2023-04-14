package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var pos = 0

func pol_rec(end int, expr []rune) int {
	res := 0
	fmt.Print(pos, " ")
	for _, v := range expr {
		fmt.Print(string(v))
	}

	if pos == end-2 {
		digit, _ := strconv.Atoi(string(expr[0]))
		res = digit
		fmt.Printf("d %d", res)
		fmt.Println()
		return res
	}

	if unicode.IsDigit(expr[0]) {
		digit, _ := strconv.Atoi(string(expr[0]))
		pos++
		res = digit
		fmt.Printf("d %d, pos: %d", res, pos)
		fmt.Println()
	} else {
		pos++
		res += pol_rec(end, expr[pos:end])
		fmt.Printf("one pos:%d, %d", pos, res)
		fmt.Println()
		if expr[0] == '+' {
			res += pol_rec(end, expr[pos:end])
			fmt.Printf("two pos: %d, %d", pos, res)
			fmt.Println()
		} else if expr[0] == '*' {
			res *= pol_rec(end, expr[pos:end])
			fmt.Printf("two %v", res)
			fmt.Println()
		} else if expr[0] == '-' {
			res -= pol_rec(end, expr[pos:end])
			fmt.Printf("two %v", res)
			fmt.Println()
		}
	}
	return res
}

func polish(str string) int {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	runes := append([]rune(str), '.')
	len := len(runes)
	if len == 0 {
		return 0
	}
	return pol_rec(len, runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()
	fmt.Println(polish(expression))
} //wants := []int{0, 3, 12, 729, 35, 135, 1, -1, 0}
