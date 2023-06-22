package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, n)
	for i := range a {
		a[i] = scan()
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i]*(i+1-1) + a[i]*(i+1-n)
	}
	fmt.Println(ans)
}
