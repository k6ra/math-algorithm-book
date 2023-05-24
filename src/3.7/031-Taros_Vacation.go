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
	for i := 1; i <= n; i++ {
		a[i] = scan()
	}

	dp := make([]int, n+1)
	dp[1] = a[1]
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]
		if dp[i-1] < a[i]+dp[i-2] {
			dp[i] = a[i] + dp[i-2]
		}
	}

	fmt.Println(dp[n])
}
