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

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scan()
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = scan()
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	ans := 0
	for i := 0; i < n; i++ {
		ans += abs(a[i] - b[i])
	}

	fmt.Println(ans)
}
