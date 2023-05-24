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
	in, _ := strconv.Atoi(sc.Text())
	return in
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scan()
	w := scan()
	ws := make([]int, n+1)
	vs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ws[i] = scan()
		vs[i] = scan()
	}

	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, w+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			if j < ws[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-ws[i]]+vs[i])
			}
		}
	}

	fmt.Println(dp[n][w])
}
