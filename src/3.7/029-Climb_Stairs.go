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

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		twoBefore := 0
		if i-2 >= 0 {
			twoBefore = dp[i-2]
		}
		dp[i] = dp[i-1] + twoBefore
	}

	fmt.Println(dp[n])
}
