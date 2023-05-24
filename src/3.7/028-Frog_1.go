package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	in, _ := strconv.Atoi(sc.Text())
	return in
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scan()
	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		h[i] = scan()
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = abs(h[1] - h[2])
	for i := 3; i <= n; i++ {
		oneBefore := abs(h[i]-h[i-1]) + dp[i-1]
		twoBefore := abs(h[i]-h[i-2]) + dp[i-2]
		dp[i] = oneBefore
		if oneBefore > twoBefore {
			dp[i] = twoBefore
		}
	}

	fmt.Println(dp[n])
}
