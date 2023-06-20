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

	const mod = 1000000007
	n := scan()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scan()
	}

	power := make([]int, n+1)
	power[1] = 1
	for i := 2; i <= n; i++ {
		power[i] = power[i-1] * 2 % mod
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += power[i] * a[i] % mod
	}

	fmt.Println(ans % mod)
}
