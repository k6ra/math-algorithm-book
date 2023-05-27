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

	color := make([]int, n+1)
	color[1] = 1

	queue := make([]int, 0, n)
	queue = append(queue, 1)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		nextColor := 1
		if color[current] == 1 {
			nextColor = 2
		}
		for _, v := range ab[current] {
			if color[v] == 0 {
				color[v] = nextColor
				queue = append(queue, v)
			} else if color[v] != nextColor {
				fmt.Println("No")
				return
			}
		}

		if len(queue) == 0 {
			for i := 1; i <= n; i++ {
				if color[i] == 0 {
					color[i] = 1
					queue = append(queue, i)
					break
				}
			}
		}
	}

	fmt.Println("Yes")
}
