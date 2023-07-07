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

var w = bufio.NewWriter(os.Stdout)

func write(text string) {
	defer w.Flush()

	fmt.Fprintln(w, text)
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	c1 := make([]int, n+1)
	c2 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		c := scan()
		p := scan()
		if c == 1 {
			c1[i] = c1[i-1] + p
			c2[i] = c2[i-1]
		} else {
			c1[i] = c1[i-1]
			c2[i] = c2[i-1] + p
		}
	}

	q := scan()
	ans := make([][2]int, q)
	for i := 0; i < q; i++ {
		l := scan()
		r := scan()

		ans[i][0] = c1[r] - c1[l-1]
		ans[i][1] = c2[r] - c2[l-1]
	}

	for i := range ans {
		write(fmt.Sprintf("%d %d", ans[i][0], ans[i][1]))
	}
}
