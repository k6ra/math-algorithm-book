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
	for i := 0; i < m; i++ {
		a := scan()
		b := scan()
		ab[a] = append(ab[a], b)
		ab[b] = append(ab[b], a)
	}

	count := make([]int, n+1)
	done := make([]bool, n+1)

	queue := make([]int, 0, m*2)
	queue = append(queue, 1)
	done[1] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, v := range ab[current] {
			if v < current {
				count[current]++
			}
			if !done[v] {
				done[v] = true
				queue = append(queue, v)
			}
		}
	}

	ans := 0
	for _, v := range count {
		if v == 1 {
			ans++
		}
	}

	fmt.Println(ans)
}
