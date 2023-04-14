package main

import (
	"container/heap"
	"fmt"
)

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].key < pq[j].key
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Vertex, key int) {
	item.key = key
	heap.Fix(pq, item.index)
}

type Vertex struct {
	index, key int
	value      []Edge
}

type Edge struct {
	u, length int
}

type PriorityQueue []*Vertex

func Prim(g []Vertex) int {
	res := 0
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	v := &(g[0])
	for {
		fmt.Println(res, q)
		v.index = -2
		for _, r := range v.value {
			u := &g[r.u]
			if u.index == -1 {
				u.key = r.length
				heap.Push(&q, u)
			} else if u.index != -2 && r.length < u.key {
				q.update(u, r.length)
			}

		}

		if q.Len() == 0 {
			break
		}
		v = heap.Pop(&q).(*Vertex)
		res += v.key
	}
	return res
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	graph := make([]Vertex, n)

	for i := 0; i < n; i++ {
		graph[i] = Vertex{}
		graph[i].index = -1
	}

	for i := 0; i < m; i++ {
		var a, b, l int
		fmt.Scan(&a, &b, &l)

		graph[a].value = append(graph[a].value, Edge{b, l})
		graph[b].value = append(graph[b].value, Edge{a, l})

	}

	fmt.Println(Prim(graph))
}
