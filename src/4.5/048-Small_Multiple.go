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

	k := scan()

	dist := make([]int, k, k)
	for i := 0; i < k; i++ {
		dist[i] = 1 << 29
	}

	queue := make([]int, 0, k)
	dist[1] = 1
	queue = append(queue, 1)
	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]

		val10 := (val * 10) % k
		if dist[val10] > dist[val] {
			dist[val10] = dist[val]
			queue = append([]int{val10}, queue...)
		}

		val1 := (val + 1) % k
		if dist[val1] > dist[val]+1 {
			dist[val1] = dist[val] + 1
			queue = append(queue, val1)
		}
	}

	fmt.Println(dist[0])
}
