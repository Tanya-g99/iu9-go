package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (num frac) Normalise() frac {
	a, b := num.a, num.b

	if a < 0 && b < 0 {
		a, b = -a, -b
	} else if a < 0 || b < 0 {
		a, b = -abs(a), abs(b)
	}

	a1, nod := abs(a), abs(b)
	for a1 != 0 {
		a1, nod = nod%a1, a1
	}

	return frac{a / nod, b / nod}
}

func (num frac) Flip() frac { return frac{num.b, num.a} }

func (num frac) Minus() frac { return frac{-num.a, num.b} }

func (num1 frac) Mul(num2 frac) frac { return frac{num1.a * num2.a, num1.b * num2.b}.Normalise() }

func (num1 frac) Add(num2 frac) frac {
	return frac{num1.a*num2.b + num1.b*num2.a, num1.b * num2.b}.Normalise()
}

type frac struct {
	a, b int
}

func swap(line1, line2 []frac, n int) {
	for i := 0; i < n; i++ {
		line1[i], line2[i] = line2[i], line1[i]
	}
}

func add_line(line, line_i []frac, k frac, n int) { // прибавление к строке строки k*a_i
	for i := 0; i < n; i++ {
		line[i] = line[i].Add(line_i[i].Mul(k))
	}
}

func mul_num(line []frac, k frac, n int) { //умножение строки на число k
	for j := 0; j < n; j++ {
		line[j] = line[j].Mul(k)
	}
}

func div_num(line []frac, num frac, n int) { //деление строки на число
	mul_num(line, num.Flip(), n)
}

func display(matrix [][]frac) {
	fmt.Printf("\n")
	for _, line := range matrix {
		for _, elem := range line {
			fmt.Printf("%d/%d ", elem.a, elem.b)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func triangularMatrix(matrix [][]frac, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && matrix[i][i].a == 0; j++ { //меняем строчки, чтобы на диагонали был ненулевой эл-т,
			swap(matrix[i], matrix[j], n) //                       если такой есть в столбце
		}
		if matrix[i][i].a == 0 {
			return
		}

		div_num(matrix[i], matrix[i][i], n+1) //сокращаем строку на аii => 1)вычитаем строки, умноженные на просто aji, а не aji/aii
		//                                                                 2)значения на диагонали будут единичными
		for j := i + 1; j < n; j++ {
			if matrix[j][i].a != 0 {
				add_line(matrix[j], matrix[i], matrix[j][i].Minus(), n+1) //зануляем под aii
			}
		}
	} //усё - треугольная
}

func gaussmethod(matrix [][]frac, n int) []frac {
	triangularMatrix(matrix, n)
	//display(matrix)
	solution_x := make([]frac, n)
	for i := n - 1; i >= 0; i-- {
		if matrix[i][i].a == 0 { //если нулевой, то бесконечное число решений
			return nil //                  в нашем случае нет решения
		}
		value_xi := matrix[i][n] //значение справа
		for j := i + 1; j < n; j++ {
			value_xi = value_xi.Add(matrix[i][j].Minus().Mul(solution_x[j])) //xi = value_r - ain*xn - ... - aij*xj
		}
		solution_x[i] = value_xi
	}
	if len(solution_x) == 0 {
		return nil
	}
	return solution_x
}

func main() {
	var n int
	fmt.Scan(&n)

	var matrix [][]frac
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]frac, n+1))
		for j := 0; j < n+1; j++ {
			fmt.Scan(&matrix[i][j].a)
			matrix[i][j].b = 1
		}
	}

	solution := gaussmethod(matrix, n)

	if solution != nil {
		for _, value := range solution {
			fmt.Printf("%d/%d\n", value.a, value.b)
		}
	} else {
		fmt.Println("No solution")
	}
}
