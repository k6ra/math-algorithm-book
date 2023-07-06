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
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = scan()
	}

	a := make([]int, n)
	a[0] = b[0]
	for i := 1; i < n-1; i++ {
		if b[i-1] < b[i] {
			a[i] = b[i-1]
		} else {
			a[i] = b[i]
		}
	}
	a[n-1] = b[n-2]

	ans := 0
	for i := range a {
		ans += a[i]
	}

	fmt.Println(ans)
}
