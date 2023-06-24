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

	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		a := scan()
		b := scan()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	ans := make([]int, n+1)
	for i := range ans {
		ans[i] = -1
	}

	queue := make([]int, 0, n)
	queue = append(queue, 1)
	ans[1] = 0
	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]

		for _, v := range g[person] {
			if ans[v] == -1 {
				ans[v] = ans[person] + 1
				queue = append(queue, v)
			}
		}
	}

	for i := 1; i <= n; i++ {
		if 0 <= ans[i] && ans[i] <= 120 {
			fmt.Println(ans[i])
		} else {
			fmt.Println(120)
		}
	}
}
