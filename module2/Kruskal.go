package main

import (
	"fmt"
	"math"
	"sort"
)

var Points []Point

type Point struct {
	x, y int
}

type Edge struct {
	begin int
	end   int
	dist  float64
}

var parent, depth []int

func Find(x int) int {
	if x == parent[x] {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

func Union(x, y int) {
	x, y = Find(x), Find(y)
	if x != y {
		if depth[x] < depth[y] {
			x, y = y, x
		}
		parent[y] = x
		if depth[x] == depth[y] {
			depth[x]++
		}
	}
}

func NewClass(n int) {
	parent, depth = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}

func Kruskal(m int, arr []Edge) float64 {
	var sum float64 = 0
	sort.Slice(arr, func(i, j int) bool { return arr[i].dist < arr[j].dist })

	NewClass(m)

	for i := 0; i < len(arr); i++ {
		if Find(arr[i].begin) != Find(arr[i].end) {
			sum += arr[i].dist
			//fmt.Println(arr[i].dist)
			Union(arr[i].begin, arr[i].end)
		}
	}

	return sum

}

func Len2(edge Edge) float64 {
	return (float64)(Points[edge.begin].x-Points[edge.end].x)*(float64)(Points[edge.begin].x-Points[edge.end].x) +
		(float64)(Points[edge.begin].y-Points[edge.end].y)*(float64)(Points[edge.begin].y-Points[edge.end].y)
}

func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	Points = make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &Points[i].x, &Points[i].y)
	}

	arr := make([]Edge, 0, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var a Edge
			a.begin, a.end = i, j
			a.dist = math.Sqrt(Len2(a))
			arr = append(arr, a)
		}
	}

	sum := Kruskal(n, arr)

	fmt.Printf("%.2f\n", sum)
}
