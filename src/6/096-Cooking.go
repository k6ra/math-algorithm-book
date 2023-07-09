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

	sum := 0
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = scan()
		sum += t[i]
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
				continue
			}
			if j >= t[i-1] && dp[i-1][j-t[i-1]] {
				dp[i][j] = true
			}
		}
	}

	ans := 1 << 30
	for i := 0; i <= sum; i++ {
		if dp[n][i] {
			cookTime := i
			if i < sum-i {
				cookTime = sum - i
			}
			if ans > cookTime {
				ans = cookTime
			}
		}
	}
	fmt.Println(ans)
}
