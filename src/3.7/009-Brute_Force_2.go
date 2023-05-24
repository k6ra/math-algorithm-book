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
	s := scan()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scan()
	}

	dp := make([][]bool, n+1)
	for i, _ := range dp {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if !dp[i-1][j] {
				continue
			}
			dp[i][j] = dp[i-1][j]
			if j+a[i] <= s {
				dp[i][j+a[i]] = true
			}
		}
	}

	ans := "No"
	if dp[n][s] {
		ans = "Yes"
	}
	fmt.Println(ans)
}
