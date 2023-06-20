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

	ans := 0
	for i := 1; i <= n; i++ {
		ans += a[i]*(i-1) + a[i]*(i-n)
	}

	fmt.Println(ans)
}
