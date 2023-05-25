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
	a := make([]int, n+1)
	diff := make([]int, n+1)
	for i := 1; i < n; i++ {
		a[i] = scan()
		diff[i+1] = diff[i] + a[i]
	}

	m := scan()
	b := make([]int, m+1)
	dist := 0

	b[1] = scan()
	for i := 2; i <= m; i++ {
		b[i] = scan()
		d := diff[b[i]] - diff[b[i-1]]
		if d < 0 {
			d *= -1
		}
		dist += d
	}

	fmt.Println(dist)
}
