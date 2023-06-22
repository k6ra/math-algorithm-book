package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = scan()
		y[i] = scan()
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	ans := 0
	for i := 0; i < n; i++ {
		ans += x[i]*(i+1-1) + x[i]*(i+1-n)
		ans += y[i]*(i+1-1) + y[i]*(i+1-n)
	}

	fmt.Println(ans)
}
