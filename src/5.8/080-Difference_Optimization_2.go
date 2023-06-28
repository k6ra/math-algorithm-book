package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type item struct {
	to   int
	cost int
}

type PriorityQueue []item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	m := scan()

	g := make([][]item, n+1)
	for i := 0; i < m; i++ {
		a := scan()
		b := scan()
		c := scan()

		g[a] = append(g[a], item{to: b, cost: c})
		g[b] = append(g[b], item{to: a, cost: c})
	}

	fmt.Println(dijkstra(n, &g))
}

func dijkstra(n int, g *[][]item) int {
	q := &PriorityQueue{}
	heap.Init(q)
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = 1 << 60
	}
	dist[1] = 0
	heap.Push(q, item{1, 0})
	for q.Len() > 0 {
		current := heap.Pop(q).(item)
		if dist[current.to] < current.cost {
			continue
		}

		for _, next := range (*g)[current.to] {
			cost := current.cost + next.cost
			if dist[next.to] > cost {
				dist[next.to] = cost
				heap.Push(q, item{to: next.to, cost: cost})
			}
		}
	}

	if dist[n] == 1<<60 {
		return -1
	}
	return dist[n]
}
