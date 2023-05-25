package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	ab := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		a := scan()
		b := scan()
		ab[a] = append(ab[a], b)
		ab[b] = append(ab[b], a)
	}

	cost := make([]int, n+1)
	for i := 2; i <= n; i++ {
		cost[i] = -1
	}

	queue := make([]int, 0, n)
	queue = append(queue, 1)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		nextCost := cost[current] + 1
		for _, v := range ab[current] {
			if cost[v] == -1 || cost[v] > nextCost {
				cost[v] = nextCost
				queue = append(queue, v)
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Println(cost[i])
	}
}
