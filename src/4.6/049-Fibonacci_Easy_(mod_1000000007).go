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
	a[1] = 1
	a[2] = 1
	for i := 3; i <= n; i++ {
		a[i] = (a[i-2] + a[i-1]) % 1000000007
	}

	fmt.Println(a[n])
}
